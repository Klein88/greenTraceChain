/*
Package utils comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"bytes"
	"crypto/rand"
	random "math/rand"

	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	// RandomRange random
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

// RandomString random string
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

// CurrentMillSeconds current mill seconds
func CurrentMillSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

// CurrentSeconds current seconds
func CurrentSeconds() int64 {
	return time.Now().UnixNano() / 1e9
}

// PathExists path exists
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

// ConvertToPercent convert to percent
func ConvertToPercent(percent float64) string {
	doubleDecimal := fmt.Sprintf("%.2f", 100*percent)
	// Strip trailing zeroes
	for doubleDecimal[len(doubleDecimal)-1] == '0' {
		doubleDecimal = doubleDecimal[:len(doubleDecimal)-1]
	}
	// Strip the decimal point if it's trailing.
	if doubleDecimal[len(doubleDecimal)-1] == '.' {
		doubleDecimal = doubleDecimal[:len(doubleDecimal)-1]
	}
	return fmt.Sprintf("%s%%", doubleDecimal)
}

// GetHostFromAddress get host from address
func GetHostFromAddress(addr string) string {
	s := strings.Split(addr, ":")
	return s[0]
}

var (
	globalChainMaxHeightMp sync.Map //chainid - blockheight(string-> int64)
)

// GetMaxHeight 获取链最大高度
func GetMaxHeight(chainId string) int64 {
	height, _ := globalChainMaxHeightMp.Load(chainId)
	return height.(int64)
}

// SetMaxHeight 对链最大高度赋值
func SetMaxHeight(chainId string, height int64) {
	globalChainMaxHeightMp.Store(chainId, height)
}

// nolint
func RandomSleepTime() int {
	return random.Intn(20) + 50

}
