/*
Package utils comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	bcx509 "chainmaker.org/chainmaker/common/v3/crypto/x509"
)

const (
	// RandomRange 所有字符
	RandomRange = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// Base64Encode base64Encode
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode base64Decode
func Base64Decode(data string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return decodeBytes
}

// RandomString rs
func RandomString(len int) string {
	var container string
	b := bytes.NewBufferString(RandomRange)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(RandomRange[randomInt.Int64()])
	}
	return container
}

// CurrentMillSeconds cms
func CurrentMillSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

// CurrentSeconds cs
func CurrentSeconds() int64 {
	return time.Now().UnixNano() / 1e9
}

// PathExists isExist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ParseCertificate parse cert
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
		return nil, fmt.Errorf("[Parse cert] parseCertificate cert failed, %s", err)
	}

	return bcx509.ChainMakerCertToX509Cert(cert)
}

// X509CertToChainMakerCert x509 to cert
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
