/*
Package sync comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"management_backend/src/db/relation"
	"net/http"
	"strings"

	"management_backend/src/config"
	"management_backend/src/db/chain"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/utils"
)

// ErrorLogResponse error log response
type ErrorLogResponse struct {
	Code int64
	Msg  string
	Data []ErrorLog
}

// ErrorLog error log
type ErrorLog struct {
	LogId      string `json:"log_id"`
	LogType    string `json:"log_type"`
	LogContent string `json:"log_content"`
	NodeId     string `json:"node_id"`
	ChainId    string `json:"chain_id"`
	LogTime    int64  `json:"log_time"`
}

// GetChainIp get chain ip
func GetChainIp(chainId string) (string, error) {
	pool := GetSdkClientPool()
	sdkClient, ok := pool.SdkClients[chainId]
	if !ok {
		return "", errors.New("not subscribed")
	}
	host := utils.GetHostFromAddress(sdkClient.SdkConfig.Remote)
	return host, nil
}

// PullChainErrorLog pull chain error log
func PullChainErrorLog(ip string) error {
	var (
		errorLogResponse ErrorLogResponse
		err              error
	)

	// 拼接 日志agent 的url
	url := fmt.Sprintf("http://%s:%d/v1/logs", ip, config.GlobalConfig.WebConf.AgentPort)
	// nolint
	response, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return err
	}
	// 读取获取的消息
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	// http 错误处理
	if response.StatusCode != http.StatusOK {
		log.Errorf("can not fetch log, got error: %d, %s", response.StatusCode, body)
		return errors.New("fetch log failed")
	}
	// 序列化处理
	err = json.Unmarshal(body, &errorLogResponse)
	if err != nil {
		log.Error(err)
		return err
	}
	// 响应错误处理
	if errorLogResponse.Code != 0 {
		log.Error("fetch chain error log failed")
		return errors.New("fetch log failed")
	}

	if len(errorLogResponse.Data) == 0 {
		return nil
	}

	var logRecords []*dbcommon.ChainErrorLog

	for _, errorLog := range errorLogResponse.Data {
		errLog := errorLog
		// 补充 errorLog 中的nodeId 信息，添加ip地址（log agent 无法主动获取到该ip）
		errorLog.NodeId = fmt.Sprintf("%s:%s", ip, errorLog.NodeId)
		logRecord, logErr := GenerateChainErrorLogRecord(&errLog)
		logRecords = append(logRecords, logRecord)
		if logErr != nil {
			log.Error("store error_log failed")
		}
	}

	chainInfo, err := chain.GetChainByChainId(logRecords[0].ChainId)
	if err != nil {
		return err
	}
	if chainInfo.AutoReport == 1 {
		err = ReportLogs(logRecords)
		if err != nil {
			log.Error("auto report error log failed")
		}
	}

	return nil
}

// GenerateChainErrorLogRecord 转换成数据库里的格式并存储
func GenerateChainErrorLogRecord(errorLog *ErrorLog) (*dbcommon.ChainErrorLog, error) {
	logContent, err := json.Marshal(errorLog)
	if err != nil {
		log.Error("Marshal ErrorLog failed")
		return nil, err
	}
	record := &dbcommon.ChainErrorLog{
		ChainId: errorLog.ChainId,
		NodeId:  errorLog.NodeId,
		LogId:   errorLog.LogId,
		LogTime: errorLog.LogTime,
		Type:    errorLog.LogType,
		Log:     string(logContent),
	}
	err = chain.CreateErrorLogRecord(record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// ReportInfo reportInfo
type ReportInfo struct {
	ManagementId string `json:"management_id"`
	Log          string `json:"log"`
	ChainConfig  string `json:"chainconfig"`
}

// ReportLogs report logs
func ReportLogs(logRecords []*dbcommon.ChainErrorLog) error {
	var (
		pushData []*ReportInfo
		err      error
	)

	for _, logRecord := range logRecords {
		chainConfig, configErr := chain.GetLastChainConfigRecord(logRecord.ChainId, logRecord.LogTime)
		if configErr != nil {
			log.Error("GetLastChainConfigRecord of %s failed", logRecord.LogId)
			continue
		}

		info := &ReportInfo{
			ManagementId: config.CMMID,
			Log:          logRecord.Log,
			ChainConfig:  chainConfig.Config,
		}
		pushData = append(pushData, info)
	}

	jsonData, err := json.Marshal(pushData)
	if err != nil {
		log.Error(err)
		return err
	}

	url := config.GlobalConfig.WebConf.ReportUrl
	// nolint
	response, err := http.Post(url, "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(err)
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Errorf("can not fetch log, got error: %d, %s", response.StatusCode, body)
		return errors.New("fetch log failed")
	}

	return nil
}

// ReportChainInfo report chain info
type ReportChainInfo struct {
	TxNum        int64  `json:"tx_num"`
	BlockHeight  int64  `json:"block_height"`
	NodeNum      int    `json:"node_num"`
	OrgNum       int    `json:"org_num"`
	ChainId      string `json:"chain_id"`
	ManagementId string `json:"management_id"`
}

// ReportChainData report chain data
func ReportChainData(chainId string) error {

	txNum, err := chain.GetTxNumByChainId(chainId)
	if err != nil {
		log.Error(err)
	}
	blockHeight := chain.GetMaxBlockHeight(chainId)
	nodeNum, err := relation.GetNodeCountByChainId(chainId)
	if err != nil {
		log.Error(err)
	}
	orgNum, err := relation.GetOrgCountByChainId(chainId)
	if err != nil {
		log.Error(err)
	}

	dataView := ReportChainInfo{
		TxNum:        txNum,
		BlockHeight:  blockHeight,
		NodeNum:      nodeNum,
		OrgNum:       orgNum,
		ChainId:      chainId,
		ManagementId: config.CMMID,
	}

	jsonData, err := json.Marshal(dataView)
	if err != nil {
		log.Error(err)
		return err
	}

	url := config.GlobalConfig.WebConf.ReportUrl
	url = strings.Replace(url, "reportLogs", "reportChainInfo", -1)
	// nolint
	response, err := http.Post(url, "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(err)
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Errorf("can not fetch log, got error: %d, %s", response.StatusCode, body)
		return errors.New("fetch log failed")
	}

	return nil
}
