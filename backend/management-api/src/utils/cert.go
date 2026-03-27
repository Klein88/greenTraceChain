/*
Package utils comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"time"

	"chainmaker.org/chainmaker/common/v3/crypto"

	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/crypto/hash"
	bcx509 "chainmaker.org/chainmaker/common/v3/crypto/x509"
	"github.com/tjfoc/gmsm/sm2"
)

const (
	defaultCountry            = "CN"
	defaultLocality           = "Beijing"
	defaultProvince           = "Beijing"
	defaultOrganizationalUnit = "ChainMaker"
	defaultOrganization       = "ChainMaker"
	defaultCommonName         = "chainmaker.org"
	defaultExpireYear         = 10

	//createFileFailedErrorTemplate       = "create file failed, %s"
	//parseCertificateFailedErrorTemplate = "ParseCertificateRequest failed, %s"
)

// CACertificateConfig CACertificate config
type CACertificateConfig struct {
	PrivKey            crypto.PrivateKey
	HashType           crypto.HashType
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
	ExpireYear         int32
	Sans               []string
}

// CreatePrivKey - create private key file
func CreatePrivKey(keyType crypto.KeyType) (crypto.PrivateKey, error) {
	algoName, ok := crypto.KeyType2NameMap[keyType]
	if !ok {
		return nil, fmt.Errorf("unknown key algo type [%d]", keyType)
	}

	privKey, err := asym.GenerateKeyPair(keyType)
	if err != nil {
		return nil, fmt.Errorf("generate key pair [%s] failed, %s", algoName, err.Error())
	}

	return privKey, nil
}

// GenerateCertTemplateConfig contains necessary parameters for creating private key.
type GenerateCertTemplateConfig struct {
	PrivKey            crypto.PrivateKey
	IsCA               bool
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
	ExpireYear         int32
	Sans               []string
}

// CreateCACertificate - create ca cert file
func CreateCACertificate(cfg *CACertificateConfig) (string, error) {
	template, err := GenerateCertTemplate(&GenerateCertTemplateConfig{
		PrivKey:            cfg.PrivKey,
		IsCA:               true,
		Country:            cfg.Country,
		Locality:           cfg.Locality,
		Province:           cfg.Province,
		OrganizationalUnit: cfg.OrganizationalUnit,
		Organization:       cfg.Organization,
		CommonName:         cfg.CommonName,
		ExpireYear:         cfg.ExpireYear,
		Sans:               cfg.Sans,
	})
	if err != nil {
		return "", fmt.Errorf("generateCertTemplate failed, %s", err.Error())
	}

	template.SubjectKeyId, err = ComputeSKI(cfg.HashType, cfg.PrivKey.PublicKey().ToStandardKey())
	if err != nil {
		return "", fmt.Errorf("create CA cert compute SKI failed, %s", err.Error())
	}

	certPemStr, err := createCertificate(cfg.PrivKey, template, template)
	if err != nil {
		return "", fmt.Errorf("createCertificate failed, %s", err.Error())
	}

	return certPemStr, nil
}

func createCertificate(privKey crypto.PrivateKey, template, parent *x509.Certificate) (string, error) {

	x509certEncode, err := bcx509.CreateCertificate(rand.Reader, template, parent,
		privKey.PublicKey().ToStandardKey(), privKey.ToStandardKey())
	if err != nil {
		return "", err
	}
	certPemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode})
	return string(certPemBytes), nil
}

// GenerateCertTemplate generate cert template
func GenerateCertTemplate(cfg *GenerateCertTemplateConfig) (*x509.Certificate, error) {
	sn, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return nil, err
	}
	notBefore := time.Now().Add(-10 * time.Minute).UTC()

	c := cfg.Country
	if c == "" {
		c = defaultCountry
	}

	l := cfg.Locality
	if l == "" {
		l = defaultLocality
	}

	p := cfg.Province
	if p == "" {
		p = defaultProvince
	}

	ou := cfg.OrganizationalUnit
	if ou == "" {
		ou = defaultOrganizationalUnit
	}

	o := cfg.Organization
	if o == "" {
		o = defaultOrganization
	}

	cn := cfg.CommonName
	if cn == "" {
		cn = defaultCommonName
	}

	basicConstraintsValid := false
	if cfg.IsCA {
		basicConstraintsValid = true
	}

	expireYear := cfg.ExpireYear
	if expireYear <= 0 {
		expireYear = defaultExpireYear
	}

	signatureAlgorithm := getSignatureAlgorithm(cfg.PrivKey)
	dnsName, ipAddrs := dealSANS(cfg.Sans)

	template := &x509.Certificate{
		SignatureAlgorithm:    signatureAlgorithm,
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Duration(expireYear) * 365 * 24 * time.Hour).UTC(),
		BasicConstraintsValid: basicConstraintsValid,
		IsCA:                  cfg.IsCA,
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IPAddresses: ipAddrs,
		DNSNames:    dnsName,
		Subject: pkix.Name{
			Country:            []string{c},
			Locality:           []string{l},
			Province:           []string{p},
			OrganizationalUnit: []string{ou},
			Organization:       []string{o},
			CommonName:         cn,
		},
	}

	return template, nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

// ComputeSKI compute SKI
func ComputeSKI(hashType crypto.HashType, pub interface{}) ([]byte, error) {
	encodedPub, err := bcx509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	var subPKI subjectPublicKeyInfo
	_, err = asn1.Unmarshal(encodedPub, &subPKI)
	if err != nil {
		return nil, err
	}

	pubHash, err := hash.Get(hashType, subPKI.SubjectPublicKey.Bytes)
	if err != nil {
		return nil, err
	}

	return pubHash[:], nil
}

func getSignatureAlgorithm(privKey crypto.PrivateKey) x509.SignatureAlgorithm {
	signatureAlgorithm := x509.ECDSAWithSHA256
	switch privKey.PublicKey().ToStandardKey().(type) {
	case *rsa.PublicKey:
		signatureAlgorithm = x509.SHA256WithRSA
	case *sm2.PublicKey:
		signatureAlgorithm = x509.SignatureAlgorithm(bcx509.SM3WithSM2)
	}

	return signatureAlgorithm
}

func dealSANS(sans []string) ([]string, []net.IP) {

	var dnsName []string
	var ipAddrs []net.IP

	for _, san := range sans {
		ip := net.ParseIP(san)
		if ip != nil {
			ipAddrs = append(ipAddrs, ip)
		} else {
			dnsName = append(dnsName, san)
		}
	}

	return dnsName, ipAddrs
}

//privKey crypto.PrivateKey, issuerCert *x509.Certificate, csr *bcx509.CertificateRequest, sn *big.Int

// IssueCertificateConfig contains necessary parameters for issuing cert.
type IssueCertificateConfig struct {
	HashType         crypto.HashType
	IsCA             bool
	IssuerPrivKeyPwd []byte
	ExpireYear       int32
	Sans             []string
	Uuid             string
	PrivKey          crypto.PrivateKey
	IssuerCert       *x509.Certificate
	Csr              *bcx509.CertificateRequest
}

// IssueCertificate - issue certification
func IssueCertificate(cfg *IssueCertificateConfig) (string, error) {

	csr := cfg.Csr
	issuerCert := cfg.IssuerCert
	privKey := cfg.PrivKey
	sn, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", fmt.Errorf("get sn failed, %s", err)
	}

	basicConstraintsValid := false
	if cfg.IsCA {
		basicConstraintsValid = true
	}

	expireYear := cfg.ExpireYear
	if expireYear <= 0 {
		expireYear = defaultExpireYear
	}

	dnsName, ipAddrs := dealSANS(cfg.Sans)

	var extraExtensions []pkix.Extension
	if cfg.Uuid != "" {
		extSubjectAltName := pkix.Extension{}
		extSubjectAltName.Id = bcx509.OidNodeId
		extSubjectAltName.Critical = false
		extSubjectAltName.Value = []byte(cfg.Uuid)

		extraExtensions = append(extraExtensions, extSubjectAltName)
	}

	notBefore := time.Now().Add(-10 * time.Minute).UTC()
	template := &x509.Certificate{
		Signature:             csr.Signature,
		SignatureAlgorithm:    x509.SignatureAlgorithm(csr.SignatureAlgorithm),
		PublicKey:             csr.PublicKey,
		PublicKeyAlgorithm:    x509.PublicKeyAlgorithm(csr.PublicKeyAlgorithm),
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Duration(expireYear) * 365 * 24 * time.Hour).UTC(),
		BasicConstraintsValid: basicConstraintsValid,
		IsCA:                  cfg.IsCA,
		Issuer:                issuerCert.Subject,
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IPAddresses:     ipAddrs,
		DNSNames:        dnsName,
		ExtraExtensions: extraExtensions,
		Subject:         csr.Subject,
	}

	if issuerCert.SubjectKeyId != nil {
		template.AuthorityKeyId = issuerCert.SubjectKeyId
	} else {
		template.AuthorityKeyId, err = ComputeSKI(cfg.HashType, issuerCert.PublicKey)
		if err != nil {
			return "", fmt.Errorf("issue cert compute issuer cert SKI failed, %s", err.Error())
		}
	}

	template.SubjectKeyId, err = ComputeSKI(cfg.HashType, csr.PublicKey.ToStandardKey())
	if err != nil {
		return "", fmt.Errorf("issue cert compute csr SKI failed, %s", err.Error())
	}

	x509certEncode, err := bcx509.CreateCertificate(rand.Reader, template, issuerCert,
		csr.PublicKey.ToStandardKey(), privKey.ToStandardKey())
	if err != nil {
		return "", fmt.Errorf("issue certificate failed, %s", err)
	}

	certPemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode})
	return string(certPemBytes), nil
}

// CSRConfig contains necessary parameters for creating csr.
type CSRConfig struct {
	PrivKey            crypto.PrivateKey
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
}

// CreateCSR create CSR
func CreateCSR(cfg *CSRConfig) (string, error) {

	templateX509 := GenerateCSRTemplate(cfg.PrivKey, cfg.Country, cfg.Locality,
		cfg.Province, cfg.OrganizationalUnit, cfg.Organization, cfg.CommonName)

	template, err := bcx509.X509CertCsrToChainMakerCertCsr(templateX509)
	if err != nil {
		return "", fmt.Errorf("generate csr failed, %s", err.Error())
	}

	data, err := bcx509.CreateCertificateRequest(rand.Reader, template, cfg.PrivKey.ToStandardKey())
	if err != nil {
		return "", fmt.Errorf("CreateCertificateRequest failed, %s", err.Error())
	}

	certPemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: data})
	return string(certPemBytes), nil
}

// GenerateCSRTemplate generate CSR template
func GenerateCSRTemplate(privKey crypto.PrivateKey,
	country, locality, province, organizationalUnit, organization, commonName string) *x509.CertificateRequest {
	c := country
	if c == "" {
		c = defaultCountry
	}

	l := locality
	if l == "" {
		l = defaultLocality
	}

	p := province
	if p == "" {
		p = defaultProvince
	}

	ou := organizationalUnit
	if ou == "" {
		ou = defaultOrganizationalUnit
	}

	o := organization
	if o == "" {
		o = defaultOrganization
	}

	cn := commonName
	if cn == "" {
		cn = defaultCommonName
	}

	signatureAlgorithm := getSignatureAlgorithm(privKey)

	template := &x509.CertificateRequest{
		SignatureAlgorithm: signatureAlgorithm,
		Subject: pkix.Name{
			Country:            []string{c},
			Locality:           []string{l},
			Province:           []string{p},
			OrganizationalUnit: []string{ou},
			Organization:       []string{o},
			CommonName:         cn,
		},
	}

	return template
}

// ParseCertificate parseCertificate
func ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	var (
		cert *bcx509.Certificate
		err  error
	)
	block, rest := pem.Decode(certBytes)
	if block == nil {
		cert, err = bcx509.ParseCertificate(rest)
	} else {
		cert, err = bcx509.ParseCertificate(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse x509 cert failed: %s", err.Error())
	}
	return bcx509.ChainMakerCertToX509Cert(cert)
}

// ParseCertificate1 parse certificate1 获取common 内部的 cert
func ParseCertificate1(certBytes []byte) (*bcx509.Certificate, error) {
	var (
		cert *bcx509.Certificate
		err  error
	)
	block, rest := pem.Decode(certBytes)
	if block == nil {
		cert, err = bcx509.ParseCertificate(rest)
	} else {
		cert, err = bcx509.ParseCertificate(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse x509 cert failed: %s", err.Error())
	}
	return cert, err
}

// ParsePrivateKey parse privateKey
func ParsePrivateKey(privateKeyBytes []byte) (crypto.PrivateKey, error) {
	var (
		privateKey crypto.PrivateKey
		err        error
	)
	block, rest := pem.Decode(privateKeyBytes)
	if block == nil {
		privateKey, err = asym.PrivateKeyFromDER(rest)
	} else {
		privateKey, err = asym.PrivateKeyFromDER(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse private key failed: %s", err.Error())
	}
	return privateKey, nil
}

//ParseCsr parse csr file to x.509 cert request
func ParseCsr(csrBytes []byte) (*bcx509.CertificateRequest, error) {
	var (
		csrBC *bcx509.CertificateRequest
		err   error
	)
	block, rest := pem.Decode(csrBytes)
	if block == nil {
		csrBC, err = bcx509.ParseCertificateRequest(rest)
	} else {
		csrBC, err = bcx509.ParseCertificateRequest(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse certificate request failed: %s", err.Error())
	}
	return csrBC, nil
}

// X509CertToChainMakerCert X509 cert to chainMaker cert
func X509CertToChainMakerCert(cert *x509.Certificate) (*bcx509.Certificate, error) {
	der, err := bcx509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("fail to parse re-encode (marshal) public key in certificate: %v", err)
	}
	pk, err := asym.PublicKeyFromDER(der)
	if err != nil {
		return nil, fmt.Errorf("fail to parse re-encode (unmarshal) public key in certificate: %v", err)
	}
	newCert := &bcx509.Certificate{
		Raw:                         cert.Raw,
		RawTBSCertificate:           cert.RawTBSCertificate,
		RawSubjectPublicKeyInfo:     cert.RawSubjectPublicKeyInfo,
		RawSubject:                  cert.RawSubject,
		RawIssuer:                   cert.RawIssuer,
		Signature:                   cert.Signature,
		SignatureAlgorithm:          bcx509.SignatureAlgorithm(cert.SignatureAlgorithm),
		PublicKeyAlgorithm:          bcx509.PublicKeyAlgorithm(cert.PublicKeyAlgorithm),
		PublicKey:                   pk,
		Version:                     cert.Version,
		SerialNumber:                cert.SerialNumber,
		Issuer:                      cert.Issuer,
		Subject:                     cert.Subject,
		NotBefore:                   cert.NotBefore,
		NotAfter:                    cert.NotAfter,
		KeyUsage:                    cert.KeyUsage,
		Extensions:                  cert.Extensions,
		ExtraExtensions:             cert.ExtraExtensions,
		UnhandledCriticalExtensions: cert.UnhandledCriticalExtensions,
		ExtKeyUsage:                 cert.ExtKeyUsage,
		UnknownExtKeyUsage:          cert.UnknownExtKeyUsage,
		BasicConstraintsValid:       cert.BasicConstraintsValid,
		IsCA:                        cert.IsCA,
		MaxPathLen:                  cert.MaxPathLen,
		MaxPathLenZero:              cert.MaxPathLenZero,
		SubjectKeyId:                cert.SubjectKeyId,
		AuthorityKeyId:              cert.AuthorityKeyId,
		OCSPServer:                  cert.OCSPServer,
		IssuingCertificateURL:       cert.IssuingCertificateURL,
		DNSNames:                    cert.DNSNames,
		EmailAddresses:              cert.EmailAddresses,
		IPAddresses:                 cert.IPAddresses,
		URIs:                        cert.URIs,
		PermittedDNSDomainsCritical: cert.PermittedDNSDomainsCritical,
		PermittedDNSDomains:         cert.PermittedDNSDomains,
		ExcludedDNSDomains:          cert.ExcludedDNSDomains,
		PermittedIPRanges:           cert.PermittedIPRanges,
		ExcludedIPRanges:            cert.ExcludedIPRanges,
		PermittedEmailAddresses:     cert.PermittedEmailAddresses,
		ExcludedEmailAddresses:      cert.ExcludedEmailAddresses,
		PermittedURIDomains:         cert.PermittedURIDomains,
		ExcludedURIDomains:          cert.ExcludedURIDomains,
		CRLDistributionPoints:       cert.CRLDistributionPoints,
		PolicyIdentifiers:           cert.PolicyIdentifiers,
	}
	return newCert, nil
}
