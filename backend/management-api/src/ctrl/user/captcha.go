/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	captcha "github.com/mojocn/base64Captcha"

	"management_backend/src/config"
	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	loggers "management_backend/src/logger"
	"management_backend/src/session"
)

const (
	memStoreCollectNum = 1024 * 256
	memStoreExpiration = 24 * time.Hour
	captchaSource      = "ACDEFGHJKLMNPRSTUVWXYZ234578"
)

var (
	once sync.Once
	cc   *captcha.Captcha
	log  = loggers.GetLogger(loggers.ModuleWeb)
)

// InitCaptcha initCaptcha
func InitCaptcha(conf *config.CaptchaConf) {
	once.Do(func() {
		// 初始化资源
		driver := captcha.NewDriverString(conf.Height, conf.Width, conf.NoiseCount,
			//captcha.OptionShowSineLine|captcha.OptionShowSlimeLine|captcha.OptionShowHollowLine,
			0,
			conf.Length, captchaSource, nil, nil)
		memStore := captcha.NewMemoryStore(memStoreCollectNum, memStoreExpiration)
		cc = captcha.NewCaptcha(driver, memStore)
	})
}

// CaptchaHandler captcha
type CaptchaHandler struct{}

// LoginVerify login verify
func (handler *CaptchaHandler) LoginVerify() bool {
	return false
}

// Handle deal
func (handler *CaptchaHandler) Handle(user *entity.User, ctx *gin.Context) {
	// 生成图片
	id, b64s, err := cc.Generate()
	if err != nil {
		// 无法生成图片，将异常返回
		common.ConvergeFailureResponse(ctx, common.ErrorGenerateCaptcha)
		return
	}

	// 将id写在session中
	err = session.CaptchaStoreSave(ctx, id)
	if err != nil {
		// 将session中的信息删除
		err = session.CaptchaStoreClear(ctx)
		if err != nil {
			log.Debugf("clean captcha failed")
		}
		common.ConvergeFailureResponse(ctx, common.ErrorGenerateCaptcha)
		return
	}
	// 将内容写入应答
	captchaView := NewCaptchaView(b64s)
	//log.Debugf("captcha string: %s", cc.Store.Get(id, false))
	common.ConvergeDataResponse(ctx, captchaView, nil)
}

// CheckCaptcha checkCaptcha
func CheckCaptcha(ctx *gin.Context, captcha string) bool {
	captchaStoreId, err := session.CaptchaStoreLoad(ctx)

	if err != nil {
		return false
	}
	// 判断是否相等
	match := cc.Verify(captchaStoreId, captcha, true)
	// 处理完成后将session中对应内容删除
	err = session.CaptchaStoreClear(ctx)
	if err != nil {
		log.Debugf("clean captcha failed")
	}
	return match
}
