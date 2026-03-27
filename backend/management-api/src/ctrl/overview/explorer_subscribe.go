/*
Package overview comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package overview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"

	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	"management_backend/src/entity"
	"management_backend/src/global"
)

const (
	// EXIST_ERR exist err
	EXIST_ERR = "already exists"
	// CONNECT_ERR connect err
	CONNECT_ERR = "not connect"
)

// ExplorerSubscribeHandler explorerSubscribeHandler
type ExplorerSubscribeHandler struct {
}

// LoginVerify login verify
func (handler *ExplorerSubscribeHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *ExplorerSubscribeHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindExplorerSubscribeHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	chainInfo, err := chain.GetChainSubscribeByChainId(params.ChainId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return
	}

	subscribeChainParams := SubscribeChainParams{
		ChainId:     chainInfo.ChainId,
		OrgId:       chainInfo.OrgId,
		Addr:        chainInfo.NodeRpcAddress,
		TLSHostName: chainInfo.TlsHostName,
		NodeList:    make([]SubscribeNode, 0),
	}
	if chainInfo.TlsHostName == "" {
		subscribeChainParams.TLSHostName = ca.TLS_HOST
	}
	subscribeChainParams.Tls = chainInfo.Tls == 0

	if chainInfo.ChainMode == global.PUBLIC {
		userInfo, userErr := chain_participant.GetPemCert(chainInfo.AdminName)
		if userErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetUserAccount)
			return
		}
		subscribeChainParams.Tls = false
		subscribeChainParams.UserKey = userInfo.PrivateKey
		subscribeChainParams.AuthType = global.PUBLIC
		subscribeChainParams.HashType = userInfo.Algorithm
	} else {
		orgCa, orgErr := chain_participant.GetOrgCaCert(chainInfo.OrgId)
		if orgErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetOrgCaCert)
			return
		}

		userInfo, userErr := chain_participant.GetUserSignCert(chainInfo.UserName)
		if userErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetUserSignCert)
			return
		}
		subscribeChainParams.OrgCA = orgCa.Cert
		subscribeChainParams.UserCert = userInfo.Cert
		subscribeChainParams.UserKey = userInfo.PrivateKey
		subscribeChainParams.AuthType = "permissionedwithcert"
	}
	node := SubscribeNode{
		Addr:        subscribeChainParams.Addr,
		OrgCA:       subscribeChainParams.OrgCA,
		TLSHostName: subscribeChainParams.TLSHostName,
		Tls:         subscribeChainParams.Tls,
	}
	subscribeChainParams.NodeList = append(subscribeChainParams.NodeList, node)
	url := fmt.Sprintf("%v/chainmaker/?cmb=SubscribeChain", params.ExplorerUrl)
	jsonByte, err := json.Marshal(subscribeChainParams)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Errorf("%s request fail %s", url, err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorExplorerSubscribe)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorExplorerUrlSubscribe)
		return
	}
	defer func() {
		if resp.Body != nil {
			err = resp.Body.Close()
			if err != nil {
				return
			}
		}
	}()
	if resp.StatusCode != 200 {
		common.ConvergeFailureResponse(ctx, common.ErrorExplorerSubscribe)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	var respJson common.ExplorerResponse
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	if respJson.Response.Error.Code != "" {
		if strings.Contains(respJson.Response.Error.Message, CONNECT_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorExplorerConnect)
			return
		}
		if strings.Contains(respJson.Response.Error.Message, EXIST_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorExplorerExist)
			return
		}
		common.ConvergeHandleFailureResponse(ctx, &respJson.Response.Error)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// SubscribeChainParams sub
type SubscribeChainParams struct {
	ChainId     string
	OrgId       string
	UserCert    string
	UserKey     string
	AuthType    string
	HashType    int
	NodeList    []SubscribeNode
	Addr        string
	OrgCA       string
	Tls         bool
	TLSHostName string
}

// SubscribeNode node
type SubscribeNode struct {
	Addr        string
	OrgCA       string
	TLSHostName string
	Tls         bool
}
