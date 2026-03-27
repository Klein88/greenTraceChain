/*
Package contract_invoke comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_invoke

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	pbcommon "chainmaker.org/chainmaker/pb-go/v3/common"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"management_backend/src/ctrl/ca"
	"management_backend/src/db"
	"management_backend/src/global"
	"management_backend/src/utils"
)

// Param param
type Param map[string]interface{}

// GetEvmKv get evm kv
func GetEvmKv(abiKey, methodName, chainMode, hashType string,
	parameterParams []*ParameterParams, crtBytes []byte) ([]*pbcommon.KeyValuePair, string, error) {
	id, userId, hash, err := ca.ResolveUploadKey(abiKey)
	if err != nil {
		return nil, "", err
	}
	upload, err := db.GetUploadByIdAndUserIdAndHash(id, userId, hash)
	if err != nil {
		return nil, "", err
	}

	myAbi, err := abi.JSON(strings.NewReader(string(upload.Content)))
	if err != nil {
		return nil, "", err
	}
	method := myAbi.Methods[methodName]

	var paramMapList []Param

	for _, input := range method.Inputs {
		for _, parameterParam := range parameterParams {
			if parameterParam.Key == input.Name {
				paramMap := make(Param)
				paramMap[input.Type.String()] = parameterParam.Value
				paramMapList = append(paramMapList, paramMap)
			}
		}
	}

	if len(method.Inputs) > 0 && len(paramMapList) < 1 {
		var addr string
		if chainMode == global.PUBLIC {
			publicKey, publicKeyErr := asym.PublicKeyFromPEM([]byte(crtBytes))
			if publicKeyErr != nil {
				return nil, "", publicKeyErr
			}
			addr, err = commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM, crypto.HashAlgoMap[hashType])
		} else {
			_, addr, _, err =
				utils.MakeAddrAndSkiFromCrtBytes(crtBytes)
		}

		if err != nil {
			return nil, "", err
		}
		paramMap := make(Param)
		paramMap["address"] = addr
		paramMapList = append(paramMapList, paramMap)
	}

	paramBytes, err := GetPaddedParam(&method, paramMapList)
	if err != nil {
		return nil, "", err
	}
	inputData := append(method.ID, paramBytes...)

	inputDataHexStr := hex.EncodeToString(inputData)

	kvs := []*pbcommon.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(inputDataHexStr),
		},
	}

	return kvs, inputDataHexStr[0:8], nil
}

// GetPaddedParam get padded param
func GetPaddedParam(method *abi.Method, param []Param) ([]byte, error) {
	values := make([]interface{}, 0)

	for _, p := range param {
		if len(p) != 1 {
			return nil, fmt.Errorf("invalid param %+v", p)
		}
		for k, v := range p {
			if k == "uint" {
				k = "uint256"
			} else if strings.HasPrefix(k, "uint[") {
				k = strings.Replace(k, "uint[", "uint256[", 1)
			}
			ty, err := abi.NewType(k, "", nil)
			if err != nil {
				return nil, fmt.Errorf("invalid param %+v: %+v", p, err)
			}

			if ty.T == abi.SliceTy || ty.T == abi.ArrayTy {
				if ty.Elem.T == abi.AddressTy {
					v, err = convetTyToAddress(v)
					if err != nil {
						return nil, err
					}
				}

				if (ty.Elem.T == abi.IntTy || ty.Elem.T == abi.UintTy) && reflect.TypeOf(v).Elem().Kind() == reflect.Interface {
					v, err = getTmpData(ty, v)
					if err != nil {
						return nil, err
					}
				}
			}

			if ty.T == abi.AddressTy {
				if v, err = convetToAddress(v); err != nil {
					return nil, err
				}
			}

			if (ty.T == abi.IntTy || ty.T == abi.UintTy) && reflect.TypeOf(v).Kind() == reflect.String {
				v = convertToInt(ty, v)
			}

			values = append(values, v)
		}
	}

	// convert params to bytes
	return method.Inputs.PackValues(values)
}

func convetTyToAddress(v interface{}) (interface{}, error) {
	tmp, ok := v.([]interface{})
	if !ok {
		log.Error("v not is []interface")
		return nil, errors.New("v not is []interface")
	}
	v = make([]common.Address, 0)
	for i := range tmp {
		addr, err := convetToAddress(tmp[i])
		if err != nil {
			return nil, err
		}
		v = append(v.([]common.Address), addr)
	}
	return v, nil
}

// 进行简单拆分
func getTmpData(ty abi.Type, v interface{}) (res interface{}, err error) {
	if ty.Elem.Size > 64 {
		tmp := make([]*big.Int, 0)
		for _, i := range v.([]interface{}) {
			if s, ok := i.(string); ok {
				value, _ := new(big.Int).SetString(s, 10)
				tmp = append(tmp, value)
			} else {
				return nil, fmt.Errorf("abi: cannot use %T as type string as argument", i)
			}
		}
		res = tmp
	} else {
		tmpI := make([]interface{}, 0)
		for _, i := range v.([]interface{}) {
			if s, ok := i.(string); ok {
				value, valueErr := strconv.ParseUint(s, 10, ty.Elem.Size)
				if valueErr != nil {
					return nil, valueErr
				}
				tmpI = append(tmpI, value)
			} else {
				return nil, fmt.Errorf("abi: cannot use %T as type string as argument", i)
			}
		}
		switch ty.Elem.Size {
		case 8:
			tmp := make([]uint8, len(tmpI))
			for i, sv := range tmpI {
				tmp[i] = uint8(sv.(uint64))
			}
			res = tmp
		case 16:
			tmp := make([]uint16, len(tmpI))
			for i, sv := range tmpI {
				tmp[i] = uint16(sv.(uint64))
			}
			res = tmp
		case 32:
			tmp := make([]uint32, len(tmpI))
			for i, sv := range tmpI {
				tmp[i] = uint32(sv.(uint64))
			}
			res = tmp
		case 64:
			tmp := make([]uint64, len(tmpI))
			ok := false
			for i, sv := range tmpI {
				tmp[i], ok = sv.(uint64)
				if !ok {
					log.Error("v not is uint64")
				}
			}
			res = tmp
		}
	}
	return res, nil
}

// convetToAddress
func convetToAddress(v interface{}) (common.Address, error) {
	switch data := v.(type) {
	case string:
		if !common.IsHexAddress(data) {
			return common.Address{}, fmt.Errorf("invalid address %s", data)
		}
		return common.HexToAddress(data), nil
	}
	return common.Address{}, fmt.Errorf("invalid address %v", v)
}

// convertToInt
func convertToInt(ty abi.Type, v interface{}) interface{} {
	if ty.T == abi.IntTy && ty.Size <= 64 {
		tmp, _ := strconv.ParseInt(v.(string), 10, ty.Size)
		switch ty.Size {
		case 8:
			v = int8(tmp)
		case 16:
			v = int16(tmp)
		case 32:
			v = int32(tmp)
		case 64:
			v = int64(tmp)
		}
	} else if ty.T == abi.UintTy && ty.Size <= 64 {
		tmp, _ := strconv.ParseUint(v.(string), 10, ty.Size)
		switch ty.Size {
		case 8:
			v = uint8(tmp)
		case 16:
			v = uint16(tmp)
		case 32:
			v = uint32(tmp)
		case 64:
			v = uint64(tmp)
		}
	} else {
		v, _ = new(big.Int).SetString(v.(string), 10)
	}
	return v
}
