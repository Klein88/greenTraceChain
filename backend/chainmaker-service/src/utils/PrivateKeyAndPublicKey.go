package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func GetPrivateKeyAndPublicKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, []byte, []byte) {

	// 生成随机的私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// 从私钥中获取公钥
	publicKey := &privateKey.PublicKey

	// 打印私钥（这里只是演示，实际不应这样做，因为私钥需要保密）
	//fmt.Printf("Private Key: %+v\n", *privateKey)

	// 打印公钥
	fmt.Printf("Public Key: %+v\n", *publicKey)

	// 将私钥编码为PEM格式
	privateKeyPEM, err := EncodePrivateKeyToPEM(privateKey)
	if err != nil {
		log.Fatalf("Failed to encode private key to PEM: %v", err)
	}
	//fmt.Printf("Private Key (PEM):\n%s", privateKeyPEM)
	// 将公钥编码为PEM格式
	publicKeyPEM, err := EncodePublicKeyToPEM(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to encode public key to PEM: %v", err)
	}
	fmt.Printf("Public Key (PEM):\n%s", publicKeyPEM)

	return privateKey, publicKey, privateKeyPEM, publicKeyPEM
}

// EncodePrivateKeyToPEM 编码私钥为PEM格式
func EncodePrivateKeyToPEM(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	privBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	privPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: privBytes,
		},
	)
	return privPEM, nil
}

// EncodePublicKeyToPEM 编码公钥为PEM格式
func EncodePublicKeyToPEM(publicKey *ecdsa.PublicKey) ([]byte, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubBytes,
		},
	)
	return pubPEM, nil
}

// 读取PEM字节格式的私钥
func ReadPrivateKeyPEMFromBytes(pemBytes []byte) (*ecdsa.PrivateKey, error) {
	// 解码PEM字符串
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	// 解析私钥
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	// 将*ecdsa.PrivateKey转换为*ecdsa.PrivateKey
	return privateKey, nil
}

// 读取PEM字节格式的公钥
func ReadPublicKeyPEM(pemBytes []byte) (*ecdsa.PublicKey, error) {

	// 解码PEM字符串
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the public key")
	}

	// 解析公钥
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// 转换为*ecdsa.PublicKey
	switch pk := parsedKey.(type) {
	case *ecdsa.PublicKey:
		return pk, nil
	default:
		return nil, fmt.Errorf("unsupported key type: %T", parsedKey)
	}
}
