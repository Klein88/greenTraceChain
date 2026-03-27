/*
Package utils comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"

	bcx509 "chainmaker.org/chainmaker/common/v3/crypto/x509"
	"chainmaker.org/chainmaker/common/v3/evmutils"
	pbcommon "chainmaker.org/chainmaker/pb-go/v3/common"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"management_backend/src/global"
)

// GetEvmMethodsByAbi get evm methods by abi
func GetEvmMethodsByAbi(content []byte) (string, int, error) {
	myAbi, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		return "", -1, err
	}
	var methods = make([]*global.Method, 0)

	// 0：正常方法 1：构造函数
	functionType := global.FUNCTION
	if len(myAbi.Constructor.Inputs) > 0 {
		functionType = global.CONSTRUCTOR
	}
	for methodName, methodVale := range myAbi.Methods {
		method := &global.Method{
			MethodFunc: "invoke",
		}
		method.MethodName = methodName

		var methodKeyStr string
		inputs := methodVale.Inputs
		for _, input := range inputs {
			methodKeyStr = methodKeyStr + input.Name + ","
		}

		methodKeyStr = strings.TrimRight(methodKeyStr, ",")
		method.MethodKey = methodKeyStr
		methods = append(methods, method)
	}

	methodJson, err := json.Marshal(methods)
	if err != nil {
		return "", -1, err
	}

	methodStr := string(methodJson)
	if methodStr == global.NULL {
		methodStr = ""
	}
	return methodStr, functionType, nil
}

// GetConstructorKeyValuePair get constructor key value pair
func GetConstructorKeyValuePair(chainMode, hashType string, crtBytes []byte,
	content []byte, params string) ([]*pbcommon.KeyValuePair, error) {
	//var addrInt *evmutils.Int
	var err error
	//if chainMode == global.PUBLIC {
	//	publicKey, publicKeyErr := asym.PublicKeyFromPEM([]byte(crtBytes))
	//	if publicKeyErr != nil {
	//		return nil, publicKeyErr
	//	}
	//	addrInt, err = commonutils.PkToAddrInt(publicKey, pbconfig.AddrType_ETHEREUM, crypto.HashAlgoMap[hashType])
	//	if err != nil {
	//		return nil, err
	//	}
	//} else {
	//	_, _, client1AddrSki, skiErr :=
	//		MakeAddrAndSkiFromCrtBytes(crtBytes)
	//	if skiErr != nil {
	//		return nil, skiErr
	//	}
	//	addrInt, err = evmutils.MakeAddressFromHex(client1AddrSki)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//addr := evmutils.BigToAddress(addrInt)
	myAbi, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		return nil, err
	}
	paramsTmp := make([]interface{}, 0)

	if params != "" {
		var paramsMap []Param
		err = json.Unmarshal([]byte(params), &paramsMap)
		if err != nil {
			return nil, err
		}

		for _, p := range paramsMap {
			switch strings.ToLower(p.Key) {
			case "uint256":
				uint256Val, ok := new(big.Int).SetString(p.Value, 10)
				if !ok {
					return nil, errors.New("uint256 not match")
				}
				paramsTmp = append(paramsTmp, uint256Val)
			case "string":
				paramsTmp = append(paramsTmp, p.Value)
			case "address":
				addressVal, addrErr := evmutils.StringToAddress(p.Value)
				if addrErr != nil {
					return nil, addrErr
				}
				paramsTmp = append(paramsTmp, addressVal)
			}
		}
	}
	dataByte, err := myAbi.Pack("", paramsTmp...)
	if err != nil {
		return nil, err
	}

	data := hex.EncodeToString(dataByte)
	pairs := []*pbcommon.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}

	return pairs, nil
}

// Param param
type Param struct {
	Key   string
	Value string
}

// MakeAddrAndSkiFromCrtBytes MakeAddrAndSkiFromCrtBytes
func MakeAddrAndSkiFromCrtBytes(crtBytes []byte) (string, string, string, error) {
	blockCrt, _ := pem.Decode(crtBytes)
	crt, err := bcx509.ParseCertificate(blockCrt.Bytes)
	if err != nil {
		return "", "", "", err
	}

	ski := hex.EncodeToString(crt.SubjectKeyId)
	addrInt, err := evmutils.MakeAddressFromHex(ski)
	if err != nil {
		return "", "", "", err
	}

	log.Info(fmt.Sprintf("0x%s", addrInt.AsStringKey()))

	return addrInt.String(), fmt.Sprintf("0x%x", addrInt.AsStringKey()), ski, nil
}
