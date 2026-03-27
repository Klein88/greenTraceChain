/*
Package utils comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"encoding/hex"

	"chainmaker.org/chainmaker/common/v3/crypto/hash"
)

const (
	// SHA256 sha256
	SHA256 = "SHA256"
)

// Sha256 sha256
func Sha256(content []byte) ([]byte, error) {
	return hash.GetByStrType(SHA256, content)
}

// Sha256HexString sha256 hex string
func Sha256HexString(content []byte) string {
	hash, err := Sha256(content)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hash)
}
