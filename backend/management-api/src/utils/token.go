/*
Package utils comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// nolint
const (
	Secret         = "chainmaker"
	ExpiresAtHours = 24
)

// NewSignedToken new signed token
func NewSignedToken(userId int64, user string) (string, error) {
	claims := &JwtClaims{
		UserId:   userId,
		UserName: user,
	}
	claims.IssuedAt = time.Now().Unix()
	// 不会过期
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(ExpiresAtHours)).Unix()
	return ToSignedToken(claims)
}

// CheckTokenTime check token time
func CheckTokenTime(claims *JwtClaims) error {
	currentTime := time.Now().Unix()
	if claims.ExpiresAt <= currentTime {
		// 已经过期
		return errors.New("token has been expired")
	}
	return nil
}

// ToSignedToken to signed token
func ToSignedToken(claims *JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// LoadJwtClaims load jwt claims
func LoadJwtClaims(tokenText string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenText, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	// 强制类型转换，类似于Java中的instance of
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("can not load token")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}

// JwtClaims jwt claims
type JwtClaims struct {
	jwt.StandardClaims
	UserId   int64
	UserName string
}
