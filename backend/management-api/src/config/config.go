/*
Package config 	comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package config

import (
	"fmt"
	"strconv"
)

const (
	// MySql mysql
	MySql = "mysql"
	// DbDefaultConf db default config
	DbDefaultConf = "?charset=utf8mb4&parseTime=True&loc=Local"
	// DbMaxIdleConns db max idle config
	DbMaxIdleConns = 50
	// DbMaxOpenConns db max open config
	DbMaxOpenConns = 50
)

// WebConf Http配置
type WebConf struct {
	Address           string       `mapstructure:"address"`
	Port              int          `mapstructure:"port"`
	CrossDomain       bool         `mapstructure:"cross_domain"`
	SessionAge        int          `mapstructure:"session_age"`
	CaptchaConf       *CaptchaConf `mapstructure:"captcha"`
	ErrmsgLang        int          `mapstructure:"errmsg_lang"`
	LoadPeriodSeconds int          `mapstructure:"load_period_seconds"`
	Password          string       `mapstructure:"password"`
	AgentPort         int          `mapstructure:"agent_port"`
	ReportUrl         string       `mapstructure:"report_url"`
}

// DBConf 数据库配置
type DBConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Passwd   string `mapstructure:"passwd"`
}

// LogConf 日志配置
type LogConf struct {
	LogLevelDefault string            `mapstructure:"log_level_default"`
	LogLevels       map[string]string `mapstructure:"log_levels"`
	FilePath        string            `mapstructure:"file_path"`
	MaxAge          int               `mapstructure:"max_age"`
	RotationTime    int               `mapstructure:"rotation_time"`
	LogInConsole    bool              `mapstructure:"log_in_console"`
	ShowColor       bool              `mapstructure:"show_color"`
}

// CaptchaConf captcha config
type CaptchaConf struct {
	Height     int `mapstructure:"height"`
	Width      int `mapstructure:"width"`
	NoiseCount int `mapstructure:"noise_count"`
	Length     int `mapstructure:"length"`
}

// Config 整体配置
type Config struct {
	WebConf *WebConf `mapstructure:"web"`
	DBConf  *DBConf  `mapstructure:"db"`
	LogConf *LogConf `mapstructure:"log"`
}

// ToUrl to url
func (dbConfig *DBConf) ToUrl() string {
	url := fmt.Sprintf("tcp(%s:%s)/%s", dbConfig.Host, dbConfig.Port, dbConfig.Database)
	return dbConfig.User + ":" + dbConfig.Passwd + "@" + url + DbDefaultConf
}

// ToUrl ToUrl
func (webConfig *WebConf) ToUrl() string {
	return webConfig.Address + ":" + strconv.Itoa(webConfig.Port)
}
