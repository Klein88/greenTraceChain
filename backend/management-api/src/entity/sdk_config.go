/*
Package entity comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

import "encoding/json"

// SdkConfig sdk config
type SdkConfig struct {
	ChainId         string
	OrgId           string
	UserName        string
	AdminName       string
	Tls             bool
	TlsHost         string
	Remote          string
	CaCert          []byte
	UserPrivKey     []byte
	UserCert        []byte
	UserSignPrivKey []byte
	UserSignCert    []byte
	AuthType        string // "permissionedwithcert" "permissionedwithkey" "public"
	UserPublicKey   []byte
	HashType        string
}

// ToJson to json
func (s *SdkConfig) ToJson() string {
	str, _ := json.Marshal(s)
	return string(str)
}
