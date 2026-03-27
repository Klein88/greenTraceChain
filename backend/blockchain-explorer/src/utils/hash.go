/*
Package utils comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/gogo/protobuf/proto"

	"chainmaker.org/chainmaker/common/v3/crypto/hash"
	"chainmaker.org/chainmaker/pb-go/v3/common"
)

const (
	// SHA256 sha
	SHA256 = "SHA256"
)

// CalcUnsignedCompleteTxBytes calc
func CalcUnsignedCompleteTxBytes(t *common.Transaction) ([]byte, error) {
	if t == nil {
		return nil, errors.New("calc unsigned complete tx bytes error, tx == nil")
	}
	var senderBytes []byte
	var err error
	if t.Sender != nil {
		senderBytes, err = proto.Marshal(t.Sender)
		if err != nil {
			return nil, err
		}
	}
	var resultBytes []byte
	if t.Result != nil {
		resultBytes, err = proto.Marshal(t.Result)
		if err != nil {
			return nil, err
		}
	}
	var payloadBytes []byte
	if t.Payload != nil {
		payloadBytes, err = proto.Marshal(t.Payload)
		if err != nil {
			return nil, err
		}
	}

	completeTxBytes := bytes.Join([][]byte{senderBytes, payloadBytes, resultBytes}, []byte{})
	return completeTxBytes, nil
}

// CalcTxHash calculate transaction hash, incloud tx.Header, tx.Payload, tx.Result
func CalcTxHash(hashType string, t *common.Transaction) ([]byte, error) {
	txBytes, err := CalcUnsignedCompleteTxBytes(t)
	if err != nil {
		return nil, err
	}

	hashedTx, err := hash.GetByStrType(hashType, txBytes)
	if err != nil {
		return nil, err
	}
	return hashedTx, nil
}

// CalcHash calc
func CalcHash(hashType string, content []byte) ([]byte, error) {
	return hash.GetByStrType(hashType, content)
}

// Sha256 sha256
func Sha256(content []byte) ([]byte, error) {
	return hash.GetByStrType(SHA256, content)
}

// Sha256HexString sha256HexString
func Sha256HexString(content []byte) string {
	hash, err := Sha256(content)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hash)
}
