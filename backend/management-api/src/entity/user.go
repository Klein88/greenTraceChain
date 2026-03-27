/*
Package entity comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

// User user
type User struct {
	Id   int64
	Name string
}

// NewUser new user
func NewUser(id int64, name string) *User {
	return &User{
		Id:   id,
		Name: name,
	}
}

// GetName getName
func (user *User) GetName() string {
	return user.Name
}

// GetId getId
func (user *User) GetId() int64 {
	return user.Id
}
