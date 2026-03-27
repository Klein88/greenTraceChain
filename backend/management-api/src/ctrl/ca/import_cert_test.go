/*
Copyright (C) BABEC. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjfoc/gmsm/sm2"

	"management_backend/src/utils"

	"chainmaker.org/chainmaker/common/v3/crypto"
	bcx509 "chainmaker.org/chainmaker/common/v3/crypto/x509"
)

func Test_SM2CertParsel(t *testing.T) {
	caCertStr := "-----BEGIN CERTIFICATE-----\nMIIChjCCAiugAwIBAgIDCtqLMAoGCCqBHM9VAYN1MIGAMQswCQYDVQQGEwJDTjEQ\nMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEaMBgGA1UEChMRdGYt\nb3JnMS50YWlmdS5vcmcxEjAQBgNVBAsTCXJvb3QtY2VydDEdMBsGA1UEAxMUY2Eu\ndGYtb3JnMS50YWlmdS5vcmcwHhcNMjIwNTI2MDc1OTQ1WhcNMzIwNTIzMDc1OTQ1\nWjCBgDELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0Jl\naWppbmcxGjAYBgNVBAoTEXRmLW9yZzEudGFpZnUub3JnMRIwEAYDVQQLEwlyb290\nLWNlcnQxHTAbBgNVBAMTFGNhLnRmLW9yZzEudGFpZnUub3JnMFkwEwYHKoZIzj0C\nAQYIKoEcz1UBgi0DQgAEWEaEWlAm73nE5QvxYJ8sRwyE5MHyL00HlGWaHKahyRuZ\nB0YUXPRxGN14H6CCBswwsT/lmQ0FiGON6AqHdyYSgaOBkTCBjjAOBgNVHQ8BAf8E\nBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zApBgNVHQ4EIgQgldYN7qz3/nPWqYyY4Blq\nW80iqR414MOS0n/+EOwmQG8wQAYDVR0RBDkwN4IOY2hhaW5tYWtlci5vcmeCCWxv\nY2FsaG9zdIIUY2EudGYtb3JnMS50YWlmdS5vcmeHBH8AAAEwCgYIKoEcz1UBg3UD\nSQAwRgIhAJp4XkM4EgNaupsABwxn4OwwpvEoZnoyZHYwRygfbTVpAiEA9vzRZ2lf\n3DyFmoCTBoDggkEYlhkS2mczW+5zCQA1MgU=\n-----END CERTIFICATE-----\n"
	caKeyStr := "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQggS47Yl8cxiC6NO+U\ncbTNCZnhWzqy3T/k9gZeEwHl5xWgCgYIKoEcz1UBgi2hRANCAARYRoRaUCbvecTl\nC/FgnyxHDITkwfIvTQeUZZocpqHJG5kHRhRc9HEY3XgfoIIGzDCxP+WZDQWIY43o\nCod3JhKB\n-----END PRIVATE KEY-----\n"

	data := []byte("abc")

	privateKey, err := utils.ParsePrivateKey([]byte(caKeyStr))
	assert.Nil(t, err)

	// SignWithOpts
	var opts crypto.SignOpts
	opts.Hash = crypto.HASH_TYPE_SM3
	opts.UID = crypto.CRYPTO_DEFAULT_UID
	signture, err := privateKey.SignWithOpts([]byte("abc"), &opts)
	assert.Nil(t, err)
	fmt.Printf("SignWithOpts: %x\n", signture)

	cert, err := utils.ParseCertificate([]byte(caCertStr))
	assert.Nil(t, err)
	pk := cert.PublicKey.(*sm2.PublicKey)

	r := pk.Verify(data, signture)
	assert.True(t, r)

	// sign
	signture, err = privateKey.Sign([]byte("abc"))
	assert.Nil(t, err)
	fmt.Printf("Sign: %x\n", signture)

	cert2, err := parseCert(caCertStr)
	assert.Nil(t, err)
	pk2 := cert2.PublicKey

	r, err = pk2.Verify(data, signture)
	assert.Nil(t, err)
	assert.True(t, r)
}

func Test_SM21CertParsel(t *testing.T) {
	var (
		msg     = []byte("cfca test")
		sm2Opts = &crypto.SignOpts{
			Hash: crypto.HASH_TYPE_SM3,
			UID:  crypto.CRYPTO_DEFAULT_UID,
		}
	)
	certStr := "-----BEGIN CERTIFICATE-----\n" +
		"MIICfjCCAiKgAwIBAgIIMwAAAAlYEVcwDAYIKoEcz1UBg3UFADBhMQswCQYDVQQG\n" +
		"EwJDTjEwMC4GA1UECgwnQ2hpbmEgRmluYW5jaWFsIENlcnRpZmljYXRpb24gQXV0\n" +
		"aG9yaXR5MSAwHgYDVQQDDBdDRkNBIEFDUyBURVNUIFNNMiBPQ0EzMzAeFw0yMjA4\n" +
		"MjYwNzE1MTFaFw0yMzA4MjYwNzE1MTFaMIGEMQswCQYDVQQGEwJDTjELMAkGA1UE\n" +
		"CAwCQkoxCzAJBgNVBAcMAkJKMR0wGwYDVQQKDBRqdC1vcmcuY2huZW5lcmd5Lm9y\n" +
		"ZzESMBAGA1UECwwJY29uc2Vuc3VzMSgwJgYDVQQDDB9jb25zZW5zdXMxLmp0LW9y\n" +
		"Zy5jaG5lbmVyZ3kub3JnMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEWsluIWKi\n" +
		"yR0f0rL1ZOpR2MRl2bNRTbMb0m6gKo2oq+9MpAOykQxj9MDmy/ZwTwbWFyyFBAzh\n" +
		"FRX/Lxflmht3dKOBnTCBmjAfBgNVHSMEGDAWgBQObS0zBI7wCm48/LZ8tTP1uhF7\n" +
		"+DAJBgNVHRMEAjAAMD0GA1UdHwQ2MDQwMqAwoC6GLGh0dHA6Ly91Y3JsLmNmY2Eu\n" +
		"Y29tLmNuL09DQTMzL1NNMi9jcmw0MzEuY3JsMA4GA1UdDwEB/wQEAwIHgDAdBgNV\n" +
		"HQ4EFgQULkfhEnLpUXnx0ym6l2KD+Fw3/TowDAYIKoEcz1UBg3UFAANIADBFAiAF\n" +
		"mHRK4N2Fy1ooPcsM6fXdTtZmvA0kNavPBneRkNTr0gIhAOjTx+Xknc+tow+xeGFX\n" +
		"Kuw5a/lEfHlNUQmN+vBuxG1G\n" +
		"-----END CERTIFICATE-----"
	certRaw := []byte(certStr)

	block, _ := pem.Decode(certRaw)
	cert, err := bcx509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)

	keyRaw := []byte("-----BEGIN EC PRIVATE KEY-----\n" +
		"MHcCAQEEIGTw58Huh6IYtNq49mpzxO2V/HTToAxw9GAEy116HyaCoAoGCCqBHM9V\n" +
		"AYItoUQDQgAEWsluIWKiyR0f0rL1ZOpR2MRl2bNRTbMb0m6gKo2oq+9MpAOykQxj\n" +
		"9MDmy/ZwTwbWFyyFBAzhFRX/Lxflmht3dA==\n" +
		"-----END EC PRIVATE KEY-----")
	assert.NoError(t, err)
	block, _ = pem.Decode(keyRaw)

	fmt.Printf(base64.StdEncoding.EncodeToString(block.Bytes))
	key, err := bcx509.ParseECPrivateKey(block.Bytes)
	assert.NoError(t, err)

	sm2Key := &sm2.PrivateKey{
		D: key.D,
		PublicKey: sm2.PublicKey{
			X:     key.X,
			Y:     key.Y,
			Curve: sm2.P256Sm2(),
		},
	}
	sig, err := sm2Key.Sign(rand.Reader, msg, nil)
	assert.NoError(t, err)

	//sm2Key, err := asym.PrivateKeyFromPEM(keyRaw, nil)
	//assert.NoError(t, err)
	//
	//sig, err := sm2Key.SignWithOpts(msg, sm2Opts)
	//assert.NoError(t, err)

	ok, err := cert.PublicKey.VerifyWithOpts(msg, sig, sm2Opts)
	assert.NoError(t, err)

	assert.True(t, ok)
}

func parseCert(root string) (cert *bcx509.Certificate, err error) {
	block, _ := pem.Decode([]byte(root))
	cert, err = bcx509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ParseCertificate cert failed, %s", err)
	}
	return
}
