/*
Package config comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package config

import (
	"fmt"
	"strconv"
	"strings"
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

// PUBLIC public模式
const PUBLIC = "public"

const (
	// LISTENING 启动
	LISTENING = 0
	// STOPPED 停止
	STOPPED = 1
)

const (
	// STOP 停止
	STOP = 1
	// END 结束
	END = 2
)

const (
	// START 正常
	START = 0
	// NO_START 异常
	NO_START = 1
	// NO_WORK 不工作
	NO_WORK = 2
)

const (
	// SM2 国密
	SM2 = 0
	// ECDSA 非国密
	ECDSA = 1
)

// ContractStatusMap 合约状态
var ContractStatusMap = map[int]string{
	0:  "正常",
	1:  "正常",
	2:  "升级失败",
	3:  "发布成功",
	4:  "升级失败",
	5:  "升级成功",
	6:  "冻结失败",
	7:  "冻结成功",
	8:  "解冻失败",
	9:  "解冻成功",
	10: "注销失败",
	11: "注销成功",
}

// WebConf Http配置
type WebConf struct {
	Address     string `mapstructure:"address"`
	Port        int    `mapstructure:"port"`
	CrossDomain bool   `mapstructure:"cross_domain"`
}

// RwSet 读写集
type RwSet struct {
	Index int    `json:"index"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NodeConf 节点配置
type NodeConf struct {
	UpdateTime int       `mapstructure:"update_time"`
	SyncTime   int       `mapstructure:"sync_time"`
	ChainId    string    `mapstructure:"chain_id"`
	OrgId      string    `mapstructure:"org_id"`
	Tls        bool      `mapstructure:"tls"`
	TlsHost    string    `mapstructure:"tls_host"`
	CaPaths    string    `mapstructure:"ca_paths"`
	Remotes    string    `mapstructure:"remotes"`
	UserConf   *CertConf `mapstructure:"user"`
}

// GetUpdateTime get
func (nodeConfig *NodeConf) GetUpdateTime() int {
	return nodeConfig.UpdateTime
}

// GetSyncTime get
func (nodeConfig *NodeConf) GetSyncTime() int {
	return nodeConfig.SyncTime
}

// GetChainId get
func (nodeConfig *NodeConf) GetChainId() string {
	return nodeConfig.ChainId
}

// GetOrgId get
func (nodeConfig *NodeConf) GetOrgId() string {
	return nodeConfig.OrgId
}

// GetTls get
func (nodeConfig *NodeConf) GetTls() bool {
	return nodeConfig.Tls
}

// GetTlsHost get
func (nodeConfig *NodeConf) GetTlsHost() string {
	return nodeConfig.TlsHost
}

// GetCaPaths get
func (nodeConfig *NodeConf) GetCaPaths() []string {
	return strings.Split(nodeConfig.CaPaths, ",")
}

// GetRemotes get
func (nodeConfig *NodeConf) GetRemotes() []string {
	return strings.Split(nodeConfig.Remotes, ",")
}

// GetUserPrivKeyFile get
func (nodeConfig *NodeConf) GetUserPrivKeyFile() string {
	if nodeConfig.UserConf == nil {
		return ""
	}
	return nodeConfig.UserConf.PrivKeyFile
}

// GetUserCertFile get
func (nodeConfig *NodeConf) GetUserCertFile() string {
	if nodeConfig.UserConf == nil {
		return ""
	}
	return nodeConfig.UserConf.CertFile
}

// CertConf 认证配置
type CertConf struct {
	PrivKeyFile string `mapstructure:"priv_key_file"`
	CertFile    string `mapstructure:"cert_file"`
}

// DBConf 数据库配置
type DBConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Passwd   string `mapstructure:"passwd"`
}

// RedisConf 数据库配置
type RedisConf struct {
	RedisUrl    string `mapstructure:"redis_url"`
	RedisPasswd string `mapstructure:"redis_passwd"`
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

// ChainConf 链基础配置
type ChainConf struct {
	ShowConfig bool `mapstructure:"show_config"`
}

// Config 整体配置
type Config struct {
	WebConf   *WebConf   `mapstructure:"web"`
	NodeConf  *NodeConf  `mapstructure:"node"`
	DBConf    *DBConf    `mapstructure:"db"`
	LogConf   *LogConf   `mapstructure:"log"`
	ChainConf *ChainConf `mapstructure:"chain"`
}

// ToUrl to
func (dbConfig *DBConf) ToUrl() string {
	url := fmt.Sprintf("tcp(%s:%s)/%s", dbConfig.Host, dbConfig.Port, dbConfig.Database)
	return dbConfig.User + ":" + dbConfig.Passwd + "@" + url + DbDefaultConf
}

// ToUrl to
func (webConfig *WebConf) ToUrl() string {
	return webConfig.Address + ":" + strconv.Itoa(webConfig.Port)
}
