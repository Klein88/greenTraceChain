/*
Package session comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package session

import (
	"encoding/gob"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

const (
	userSessionKey       = "userSessionKey"
	captchaSessionKey    = "captchaSessionKey"
	sessionCryptoKeyPair = "chainmaker-manager"
	// SessionID session id
	SessionID = "gsessionid"
)

func init() {
	// 注册session可存储的数据结构
	gob.Register(UserStore{})
}

// UserStore user store
type UserStore struct {
	ID   int64
	Name string
}

func newUserStore(id int64, userName string) *UserStore {
	return &UserStore{
		ID:   id,
		Name: userName,
	}
}

func newCaptchaStore(id string) string {
	return id
}

// NewSessionStore new session store
func NewSessionStore(sessionAge int) sessions.Store {
	// 处理session
	store := memstore.NewStore([]byte(sessionCryptoKeyPair))
	var options = sessions.Options{
		Path:     "/",
		MaxAge:   sessionAge,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	store.Options(options)
	return store
}

// UserStoreLoad user store load
func UserStoreLoad(ctx *gin.Context) (*UserStore, error) {
	session := sessions.Default(ctx)
	userSession := session.Get(userSessionKey)
	if userSession == nil {
		return nil, errors.New("session is nil")
	}
	if userStore, ok := userSession.(UserStore); ok {
		return &userStore, nil
	}
	return nil, errors.New("session store is error")
}

// UserStoreSave user store save
func UserStoreSave(ctx *gin.Context, id int64, userName string) error {
	userStore := newUserStore(id, userName)
	session := sessions.Default(ctx)
	session.Set(userSessionKey, *userStore)
	return session.Save()
}

// UserStoreClear user store clear
func UserStoreClear(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Delete(userSessionKey)
	return session.Save()
}

// CaptchaStoreLoad captcha store load
func CaptchaStoreLoad(ctx *gin.Context) (string, error) {
	session := sessions.Default(ctx)
	captchaSession := session.Get(captchaSessionKey)
	if captchaSession == nil {
		return "", errors.New("session is nil")
	}
	if captchaId, ok := captchaSession.(string); ok {
		return captchaId, nil
	}
	return "", errors.New("session store is error")
}

// CaptchaStoreSave captcha store save
func CaptchaStoreSave(ctx *gin.Context, id string) error {
	captchaStore := newCaptchaStore(id)
	session := sessions.Default(ctx)
	session.Set(captchaSessionKey, captchaStore)
	return session.Save()
}

// CaptchaStoreClear captcha store clear
func CaptchaStoreClear(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Delete(captchaSessionKey)
	return session.Save()
}
