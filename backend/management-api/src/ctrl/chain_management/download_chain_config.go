/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"management_backend/src/global"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"chainmaker.org/chainmaker/pb-go/v3/consensus"

	"management_backend/src/config"
	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	loggers "management_backend/src/logger"
	"management_backend/src/utils"

	"gopkg.in/yaml.v2"
)

const (
	// SIGN_USE sign
	SIGN_USE = "sign"
	// TLS_USE tls
	TLS_USE = "tls"
)

// MONITOR_START monitor start
const MONITOR_START = 1

// NO_TLS no tls
const NO_TLS = 1

// TLS_MODE_ONEWAY oneway
const TLS_MODE_ONEWAY = "oneway"

// NO_DOCKER_VM no docker vm
const NO_DOCKER_VM = 0

// DOCKER_VM docker vm
const DOCKER_VM = 1

// DOCKER_JAVA_VM docker java vm
const DOCKER_JAVA_VM = 1

// NO_BLOCK_TX_TIMESTAMP_VERIFE  no block tx
const NO_BLOCK_TX_TIMESTAMP_VERIFE = 0

// NO_FAST_SYNC no fast sync
const NO_FAST_SYNC = 0

// DEFAULT_RPC_TLS_MODE defauly rcp tls mode
const DEFAULT_RPC_TLS_MODE = "disable"

var log = loggers.GetLogger(loggers.ModuleWeb)

// DownloadChainConfigHandler download chain config
type DownloadChainConfigHandler struct{}

// LoginVerify login verify
func (downloadChainConfigHandler *DownloadChainConfigHandler) LoginVerify() bool {
	return true
}

// Handle deal
//
//	@Description:
//	@receiver downloadChainConfigHandler
//	@param user
//	@param ctx
func (downloadChainConfigHandler *DownloadChainConfigHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindDownloadChainConfigHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	confYml := global.GetConfYml()
	chainId := params.ChainId
	var chainName string
	var err error
	_, err = os.Stat("chain_config")
	if os.IsNotExist(err) {
		err = os.MkdirAll("chain_config", os.ModePerm)
		if err != nil {
			log.Error("make zippath err :", err.Error())
			return
		}
	}
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("Get chainInfo by chainId err : " + err.Error())
		return
	}
	if chainInfo.ChainMode == global.PUBLIC {
		chainName, err = CreatePkConfig(chainId, confYml)
		if err != nil {
			log.Error("CreatePkConfig err : " + err.Error())
		}
	} else {
		chainName, err = createConfig(chainId, confYml)
		if err != nil {
			log.Error("createConfig err : " + err.Error())
		}
	}

	common.ConvergeDataResponse(ctx, url.QueryEscape(chainName), nil)

}

// createConfig
//
//	@Description:
//	@param chainId
//	@param confYml
//	@return chainName
//	@return err
func createConfig(chainId, confYml string) (chainName string, err error) {

	//创建bc
	nodeIdMap, err := createBc(chainId, confYml)
	if err != nil {
		log.Error("CreateBc err : " + err.Error())
	}
	chainOrgNodes, err := relation.GetChainOrgByChainIdList(chainId)
	if err != nil {
		log.Error("GetChainOrgNode err : " + err.Error())
	}

	var tls int
	var dockerVm int
	var dockerJavaVm int
	monitorStart := false
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("Get chainInfo by chainId err : " + err.Error())
		chainName = chainId
		tls = 0
		dockerVm = 0
		dockerJavaVm = 0
	} else {
		chainName = chainInfo.ChainName
		tls = chainInfo.TLS
		dockerVm = chainInfo.DockerVm
		dockerJavaVm = chainInfo.DockerJavaVm
		if chainInfo.Monitor == MONITOR_START {
			err = createLogAgent(confYml)
			if err != nil {
				log.Error("createLogAgent err : " + err.Error())
			}
			monitorStart = true
		}
	}

	var nodePaths string

	for _, chainOrgNode := range chainOrgNodes {
		//创建chainmaker
		err = createChainmaker(chainId, chainOrgNode.OrgId, chainOrgNode.NodeName,
			confYml, nodeIdMap, tls, dockerVm, dockerJavaVm, chainInfo)
		if err != nil {
			log.Error("createChainmaker err : " + err.Error())
		}
		//创建bin lib
		err = createBinAndLib(chainOrgNode.OrgId, chainOrgNode.NodeName, confYml, chainInfo.DockerVm, chainInfo.DockerJavaVm)
		if err != nil {
			log.Error("createBinAndLib err : " + err.Error())
		}
		//创建cert
		err = createCert(chainId, chainOrgNode.OrgId, chainOrgNode.NodeName)
		if err != nil {
			log.Error("createCert err : " + err.Error())
		}

		nodePaths = nodePaths + "./" + chainOrgNode.OrgId + "-" + chainOrgNode.NodeName + ","
	}

	nodePaths = strings.TrimRight(nodePaths, ",")

	err = createChainmakerAndScript(chainName, confYml, nodePaths, monitorStart)
	if err != nil {
		log.Error("createChainmakerAndScript err : " + err.Error())
	}
	return
}

// createLogAgent
//
//	@Description:
//	@param confYml
//	@return error
func createLogAgent(confYml string) error {
	logAgentFile, err := os.Create("release/cmlogagentd")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = logAgentFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = logAgentFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}
	_, err = utils.CopyFile("release/cmlogagentd", confYml+"/bin/cmlogagentd")
	if err != nil {
		log.Error("CopyFile bin/cmlogagentd err : " + err.Error())
	}

	return nil
}

// createChainmakerAndScript
//
//	@Description:
//	@param chainName
//	@param confYml
//	@param nodePaths
//	@param monitorStart
//	@return error
func createChainmakerAndScript(chainName, confYml, nodePaths string, monitorStart bool) error {
	chainmakerFile, err := os.Create("release/chainmaker")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = chainmakerFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = chainmakerFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}
	_, err = utils.CopyFile("release/chainmaker", confYml+"/bin/chainmaker")
	if err != nil {
		log.Error("CopyFile bin/chainmaker err : " + err.Error())
	}

	//创建启动脚本
	startFile, err := os.Create("release/start.sh")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = startFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = startFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}
	if monitorStart {
		_, err = utils.CopyFile("release/start", confYml+"/bin/logagentd_start.sh")
		if err != nil {
			log.Error("CopyFile bin/start.sh err : " + err.Error())
		}

		err = utils.RePlace("release/start", "{node_paths}", nodePaths)
		if err != nil {
			log.Error("rePlace release/start.sh err : " + err.Error())
		}
	} else {
		_, err = utils.CopyFile("release/start.sh", confYml+"/bin/start.sh")
		if err != nil {
			log.Error("CopyFile bin/start.sh err : " + err.Error())
		}
	}

	quickStopFile, err := os.Create("release/quick_stop.sh")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = quickStopFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = quickStopFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}
	_, err = utils.CopyFile("release/quick_stop.sh", confYml+"/bin/quick_stop.sh")
	if err != nil {
		log.Error("CopyFile bin/quick_stop.sh err : " + err.Error())
	}

	err = utils.Zip("release", "./chain_config/"+chainName+".zip")
	if err != nil {
		log.Error("zip file err :", err.Error())
	}

	return nil
}

// createCert
//
//	@Description:
//	@param chainId
//	@param orgId
//	@param nodeName
//	@return error
func createCert(chainId, orgId, nodeName string) error {

	err := createNodeCert(orgId, nodeName)
	if err != nil {
		log.Error("createNodeCert err : " + err.Error())
		return err
	}

	err = createUserCert(orgId, nodeName)
	if err != nil {
		log.Error("createUserCert err : " + err.Error())
		return err
	}

	err = createOrgCert(chainId, orgId, nodeName)
	if err != nil {
		log.Error("createOrgCert err : " + err.Error())
		return err
	}

	return nil
}

// createNodeCert
//
//	@Description:
//	@param orgId
//	@param nodeName
//	@return error
func createNodeCert(orgId, nodeName string) error {
	nodeCertList, err := chain_participant.GetNodeCert(nodeName)
	if err != nil {
		log.Error("GetNodeCert erBlockr : " + err.Error())
		return err
	}

	nodeId, err := os.Create(nodeName + ".nodeid")
	defer func() {
		err = nodeId.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	if err != nil {
		log.Error(err.Error())
	}
	nodeInfo, err := chain_participant.GetNodeByNodeName(nodeName)
	if err != nil {
		log.Error("GetNodeByNodeName err : " + err.Error())
		return err
	}
	_, err = nodeId.Write([]byte(nodeInfo.NodeId))
	if err != nil {
		log.Error("nodeId Write err : " + err.Error())
	}

	err = os.MkdirAll("release/"+orgId+"-"+nodeName+"/config/"+orgId+"/certs/node/"+nodeName, os.ModePerm)
	if err != nil {
		log.Error("Mkdir org certs/node path err : " + err.Error())
	}

	err = os.Rename(nodeName+".nodeid", "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/certs/node/"+nodeName+"/"+nodeName+".nodeid")
	if err != nil {
		log.Error("Rename nodeid err : " + err.Error())
	}

	for _, nodeCert := range nodeCertList {
		if nodeCert.CertUse == global.SIGN {
			err := mkdirNodeCert(orgId, nodeName, SIGN_USE, nodeCert)
			if err != nil {
				log.Error("Mkdir certs/node path err : " + err.Error())
			}

		} else if nodeCert.CertUse == global.TLS {
			err := mkdirNodeCert(orgId, nodeName, TLS_USE, nodeCert)
			if err != nil {
				log.Error("Mkdir certs/node path err : " + err.Error())
			}
		}
	}
	return nil
}

// createUserCert
//
//	@Description:
//	@param orgId
//	@param nodeName
//	@return error
func createUserCert(orgId, nodeName string) error {
	userCertList, _, err := chain_participant.GetUserCertList(orgId)
	if err != nil {
		log.Error("GetUserCertList err : " + err.Error())
		return err
	}

	for _, userCert := range userCertList {
		userName := userCert.CertUserName
		err = os.MkdirAll("release/"+orgId+"-"+nodeName+"/config/"+orgId+"/certs/user/"+userName, os.ModePerm)
		if err != nil {
			log.Error("Mkdir org certs/user path err : " + err.Error())
		}
		if userCert.CertUse == global.SIGN {
			err = mkdirUserCert(userName, orgId, nodeName, SIGN_USE, userCert)
			if err != nil {
				log.Error("mkdirUserCerterr : " + err.Error())
			}
		} else {
			err = mkdirUserCert(userName, orgId, nodeName, TLS_USE, userCert)
			if err != nil {
				log.Error("mkdirUserCerterr : " + err.Error())
			}
		}
	}

	return nil
}

// mkdirNodeCert
//
//	@Description:
//	@param orgId
//	@param nodeName
//	@param certUse
//	@param nodeCert
//	@return error
func mkdirNodeCert(orgId, nodeName, certUse string, nodeCert *dbcommon.Cert) error {
	nodeCertName := nodeName + "." + certUse + ".crt"
	nodeKeyName := nodeName + "." + certUse + ".key"

	nodeTlsCrt, err := os.Create(nodeCertName)
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = nodeTlsCrt.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()

	nodeTlsKey, err := os.Create(nodeKeyName)
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = nodeTlsKey.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()

	_, err = nodeTlsCrt.Write([]byte(nodeCert.Cert))
	if err != nil {
		log.Error("file Write err :", err.Error())
	}
	_, err = nodeTlsKey.Write([]byte(nodeCert.PrivateKey))
	if err != nil {
		log.Error("file Write err :", err.Error())
	}

	err = os.Rename(nodeCertName, "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/certs/node/"+nodeName+"/"+nodeCertName)
	if err != nil {
		log.Error("Rename node sign.crt err : " + err.Error())
	}

	err = os.Rename(nodeKeyName, "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/certs/node/"+nodeName+"/"+nodeKeyName)
	if err != nil {
		log.Error("Rename node sign.key err : " + err.Error())
	}
	return nil
}

// mkdirUserCert
//
//	@Description:
//	@param userName
//	@param orgId
//	@param nodeName
//	@param certUse
//	@param userCert
//	@return error
func mkdirUserCert(userName, orgId, nodeName, certUse string, userCert *dbcommon.Cert) error {
	userCertName := userName + "." + certUse + ".crt"
	userKeyName := userName + "." + certUse + ".key"

	userSignCrt, err := os.Create(userCertName)
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = userSignCrt.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	userSignKey, err := os.Create(userKeyName)
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = userSignKey.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	_, err = userSignCrt.Write([]byte(userCert.Cert))
	if err != nil {
		log.Error("write file err :", err.Error())
	}
	_, err = userSignKey.Write([]byte(userCert.PrivateKey))
	if err != nil {
		log.Error("write file err :", err.Error())
	}

	err = os.Rename(userCertName, "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/certs/user/"+userName+"/"+userCertName)
	if err != nil {
		log.Error("Rename user tls.crt err : " + err.Error())
	}

	err = os.Rename(userKeyName, "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/certs/user/"+userName+"/"+userKeyName)
	if err != nil {
		log.Error("Rename user tls.key err : " + err.Error())
	}
	return nil
}

// createOrgCert
//
//	@Description:
//	@param chainId
//	@param orgId
//	@param nodeName
//	@return error
func createOrgCert(chainId, orgId, nodeName string) error {
	chainOrgs, err := relation.GetChainOrgList(chainId)
	if err != nil {
		log.Error("GetChainOrgByChainIdList err : " + err.Error())
		return err
	}

	for _, orgInfo := range chainOrgs {
		err = os.MkdirAll("release/"+orgId+"-"+nodeName+"/config/"+orgId+"/certs/ca/"+orgInfo.OrgId, os.ModePerm)
		if err != nil {
			log.Error("Mkdir org certs/ca path err : " + err.Error())
		}
		orgCert, err := chain_participant.GetOrgCaCert(orgInfo.OrgId)
		if err != nil {
			log.Error("GetOrgCaCert err : " + err.Error())
			return err
		}
		f, err := os.Create("ca.crt")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		_, err = f.Write([]byte(orgCert.Cert))
		if err != nil {
			log.Error("write file err :", err.Error())
		}

		err = os.Rename("ca.crt", "release/"+
			orgId+"-"+nodeName+"/config/"+orgId+"/certs/ca/"+orgInfo.OrgId+"/ca.crt")
		if err != nil {
			log.Error("Rename chainmaker.yml err : " + err.Error())
		}
		defer func() {
			err = f.Close()
			if err != nil {
				log.Error("close file err :", err.Error())
			}
		}()
	}

	return nil
}

// createBinAndLib
//
//	@Description:
//	@param orgId
//	@param nodeName
//	@param confYml
//	@param dockerVm
//	@return error
func createBinAndLib(orgId, nodeName, confYml string, dockerVm int, dockerJavaVm int) error {
	err := os.MkdirAll("release/"+orgId+"-"+nodeName+"/bin", os.ModePerm)
	if err != nil {
		log.Error("Mkdir bin path err : " + err.Error())
	}
	dockerEnable := dockerVm == DOCKER_VM
	dockerJavaEnable := dockerJavaVm == DOCKER_JAVA_VM
	replace := map[string]string{
		"{org_id}":             orgId,
		"{docker_enable}":      strconv.FormatBool(dockerEnable),
		"{docker_java_enable}": strconv.FormatBool(dockerJavaEnable),
		"{node_addr}":          fmt.Sprintf("%s-%s", orgId, nodeName),
	}

	restartFile, err := os.Create("release/" + orgId + "-" + nodeName + "/bin/restart.sh")
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = restartFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = restartFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}
	stopFile, err := os.Create("release/" + orgId + "-" + nodeName + "/bin/stop.sh")
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = stopFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = stopFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}

	_, err = utils.CopyFile("release/"+
		orgId+"-"+nodeName+"/bin/restart", confYml+"/bin/restart.sh")
	if err != nil {
		log.Error("CopyFile bin/restart.sh err : " + err.Error())
	}

	err = utils.RePlaceMore("release/"+orgId+"-"+nodeName+"/bin/restart", replace)
	if err != nil {
		log.Error("rePlace bin/restart.sh err : " + err.Error())
	}

	_, err = utils.CopyFile("release/"+orgId+"-"+nodeName+"/bin/stop", confYml+"/bin/stop.sh")
	if err != nil {
		log.Error("CopyFile bin/stop.sh err : " + err.Error())
	}
	err = utils.RePlaceMore("release/"+orgId+"-"+nodeName+"/bin/stop", replace)
	if err != nil {
		log.Error("rePlace bin/stop.sh err : " + err.Error())
	}
	if dockerEnable {
		err = utils.CreateAndCopy("release/"+orgId+"-"+nodeName+"/bin/docker_start", confYml+"/bin/docker_start.sh", 0777)
		if err != nil {
			log.Error("create and copy docker_start.sh file err :", err.Error())
			return err
		}
		dockerReplace := map[string]string{
			"{org_id}":    orgId,
			"{node_addr}": fmt.Sprintf("%v-%v", orgId, nodeName),
		}
		err = utils.RePlaceMore("release/"+orgId+"-"+nodeName+"/bin/docker_start", dockerReplace)
		if err != nil {
			log.Error("rePlace /bin/docker_start.sh err : " + err.Error())
		}
	}
	if dockerJavaEnable {
		err = utils.CreateAndCopy("release/"+orgId+"-"+nodeName+"/bin/docker_java_start",
			confYml+"/bin/docker_java_start.sh", 0777)
		if err != nil {
			log.Error("create and copy docker_java_start.sh file err :", err.Error())
			return err
		}
		dockerJavaReplace := map[string]string{
			"{org_id}":    orgId,
			"{node_addr}": fmt.Sprintf("%v-%v", orgId, nodeName),
		}
		err = utils.RePlaceMore("release/"+orgId+"-"+nodeName+"/bin/docker_java_start", dockerJavaReplace)
		if err != nil {
			log.Error("rePlace bin/docker_java_start.sh err : " + err.Error())
		}
	}
	err = createLib(orgId, nodeName, confYml)
	if err != nil {
		log.Error("createLib err : " + err.Error())
	}

	return nil
}

// createLib
//
//	@Description:
//	@param orgId
//	@param nodeName
//	@param confYml
//	@return error
func createLib(orgId, nodeName, confYml string) error {
	err := os.MkdirAll("release/"+orgId+"-"+nodeName+"/lib", os.ModePerm)
	if err != nil {
		log.Error("Mkdir lib/chainmaker err : " + err.Error())
	}

	dylibFile, err := os.Create("release/" + orgId + "-" + nodeName + "/lib/libwasmer.so")
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = dylibFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = dylibFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}

	wxdecFile, err := os.Create("release/" + orgId + "-" + nodeName + "/lib/wxdec")
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err = wxdecFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	err = wxdecFile.Chmod(0777)
	if err != nil {
		log.Error(err.Error())
	}

	_, err = utils.CopyFile("release/"+
		orgId+"-"+nodeName+"/lib/libwasmer.so", confYml+"/lib/libwasmer.so")
	if err != nil {
		log.Error("CopyFile lib/chainmaker err : " + err.Error())
	}

	_, err = utils.CopyFile("release/"+orgId+"-"+nodeName+"/lib/wxdec", confYml+"/lib/wxdec")
	if err != nil {
		log.Error("CopyFile lib/chainmaker.service err : " + err.Error())
	}
	return nil
}

// createChainmaker
//
//	@Description:
//	@param chainId
//	@param orgId
//	@param nodeName
//	@param confYml
//	@param nodeIdMap
//	@param tls
//	@param dockerVm
//	@param chainInfo
//	@return error
func createChainmaker(chainId, orgId, nodeName, confYml string,
	nodeIdMap map[string]*dbcommon.ChainOrgNode, tls int, dockerVm int,
	dockerJavaVm int, chainInfo *dbcommon.Chain) error {
	nodeInfo, err := relation.GetChainOrgByNodeNameAndChainId(nodeName, chainId)
	if err != nil {
		log.Error("GetNodeByNodeName err : " + err.Error())
		return err
	}
	conf := new(config.Chainmaker)
	yamlFile, _ := ioutil.ReadFile(confYml + "/config_tpl/chainmaker.yml")
	_ = yaml.Unmarshal(yamlFile, conf)

	err = setChainmaker(chainId, orgId, nodeIdMap, conf, nodeInfo, tls, dockerVm, dockerJavaVm, chainInfo)
	if err != nil {
		log.Error("setChainmaker err : " + err.Error())
	}

	chainmakerBytes, _ := yaml.Marshal(conf)
	chainmaker, err := os.Create("chainmaker.yml")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = chainmaker.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()

	_, err = chainmaker.Write(chainmakerBytes)
	if err != nil {
		log.Error(err.Error())
	}
	err = os.MkdirAll("release/"+orgId+"-"+nodeName+"/config/"+orgId, os.ModePerm)
	if err != nil {
		log.Error("Mkdir org chainmaker path err : " + err.Error())
	}
	err = os.Rename("chainmaker.yml", "release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/chainmaker.yml")
	if err != nil {
		log.Error("Rename chainmaker.yml err : " + err.Error())
		return err
	}

	logFile, err := os.Create("release/" + orgId + "-" + nodeName + "/config/" + orgId + "/log.yml")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()
	_, err = utils.CopyFile("release/"+
		orgId+"-"+nodeName+"/config/"+orgId+"/log.yml", confYml+"/config_tpl/log.yml")
	if err != nil {
		log.Error("copy log.yml err : " + err.Error())
		return err
	}

	return nil
}

// setChainmaker
//
//	@Description:
//	@param chainId
//	@param orgId
//	@param nodeIdMap
//	@param conf
//	@param nodeInfo
//	@param tls
//	@param dockerVm
//	@param chainInfo
//	@return error
func setChainmaker(chainId, orgId string, nodeIdMap map[string]*dbcommon.ChainOrgNode, conf *config.Chainmaker,
	nodeInfo *dbcommon.ChainOrgNode, tls int, dockerVm int, dockerJavaVm int, chainInfo *dbcommon.Chain) error {
	confingFile := conf.ChainLogConf.ConfigFile
	conf.ChainLogConf.ConfigFile = strings.Replace(confingFile, "{org_path}", orgId, -1)

	blockchainConf := config.BlockchainConf{}
	blockchainConf.ChainId = chainId
	blockchainConf.Genesis = "../config/" + orgId + "/chainconfig/bc1.yml"
	chainList := []*config.BlockchainConf{}
	chainList = append(chainList, &blockchainConf)
	conf.BlockchainConf = chainList

	conf.NodeConf.OrgId = orgId
	certFile := conf.NodeConf.CertFile
	certFile = strings.Replace(certFile, "{org_path}", orgId, -1)
	certFile = strings.Replace(certFile, "{node_cert_path}",
		"node/"+nodeInfo.NodeName+"/"+nodeInfo.NodeName+".sign", -1)
	conf.NodeConf.CertFile = certFile

	if chainInfo.EnableHttp == 1 {
		conf.RpcConf.GatewayConf.Enabled = true
	}

	privKeyFile := conf.NodeConf.PrivKeyFile
	privKeyFile = strings.Replace(privKeyFile, "{org_path}", orgId, -1)
	privKeyFile = strings.Replace(privKeyFile, "{node_cert_path}",
		"node/"+nodeInfo.NodeName+"/"+nodeInfo.NodeName+".sign", -1)
	conf.NodeConf.PrivKeyFile = privKeyFile

	seedList := []string{}
	for nodeId, nodeInfo := range nodeIdMap {
		var nodeIp string
		if chainInfo.Single == SINGLE {
			nodeIp = LOCAL_IP
		} else {
			nodeIp = nodeInfo.NodeIp
		}
		seedList = append(seedList, "/ip4/"+nodeIp+"/tcp/"+strconv.Itoa(nodeInfo.NodeP2pPort)+"/p2p/"+nodeId)
	}
	conf.NetConf.Seeds = seedList

	tlsCertFile := conf.NetConf.Tls.CertFile
	tlsCertFile = strings.Replace(tlsCertFile, "{org_path}", orgId, -1)
	tlsCertFile = strings.Replace(tlsCertFile, "{net_cert_path}",
		"node/"+nodeInfo.NodeName+"/"+nodeInfo.NodeName+".tls", -1)
	conf.NetConf.Tls.CertFile = tlsCertFile

	tlsKeyFile := conf.NetConf.Tls.PrivKeyFile
	tlsKeyFile = strings.Replace(tlsKeyFile, "{org_path}", orgId, -1)
	tlsKeyFile = strings.Replace(tlsKeyFile, "{net_cert_path}",
		"node/"+nodeInfo.NodeName+"/"+nodeInfo.NodeName+".tls", -1)
	conf.NetConf.Tls.PrivKeyFile = tlsKeyFile

	conf.RpcConf.Port = nodeInfo.NodeRpcPort
	conf.RpcConf.Tls.CertFile = tlsCertFile
	conf.RpcConf.Tls.PrivKeyFile = tlsKeyFile
	if tls == NO_TLS {
		conf.RpcConf.Tls.Mode = DEFAULT_RPC_TLS_MODE
	} else {
		conf.RpcConf.Tls.Mode = chainInfo.RpcTlsMode
	}
	conf.NetConf.ListenAddr = "/ip4/0.0.0.0/tcp/" + strconv.Itoa(nodeInfo.NodeP2pPort)

	storePath := conf.StorageConf.StorePath
	conf.StorageConf.StorePath = strings.Replace(storePath, "{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	blockPath := conf.StorageConf.BlockdbConfig.LeveldbConfig.StorePath
	conf.StorageConf.BlockdbConfig.LeveldbConfig.StorePath = strings.Replace(blockPath,
		"{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	statePath := conf.StorageConf.StatedbConfig.LeveldbConfig.StorePath
	conf.StorageConf.StatedbConfig.LeveldbConfig.StorePath = strings.Replace(statePath,
		"{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	historyPath := conf.StorageConf.HistorydbConfig.LeveldbConfig.StorePath
	conf.StorageConf.HistorydbConfig.LeveldbConfig.StorePath = strings.Replace(historyPath,
		"{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	resultPath := conf.StorageConf.ResultdbConfig.LeveldbConfig.StorePath
	conf.StorageConf.ResultdbConfig.LeveldbConfig.StorePath = strings.Replace(resultPath,
		"{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	if dockerVm == NO_DOCKER_VM {
		conf.VmConf.DockerGo.Enable = false
	} else {
		conf.VmConf.DockerGo.Enable = true
		conf.VmConf.Common.RuntimeServer.Port = nodeInfo.NodeRpcPort + 20050
		conf.VmConf.DockerGo.ContractEngine.Port = nodeInfo.NodeRpcPort + 10050
	}

	if dockerJavaVm == DOCKER_JAVA_VM {
		conf.VmConf.DockerJava.Enable = true
		conf.VmConf.Common.RuntimeServer.Port = nodeInfo.NodeRpcPort + 20050
		conf.VmConf.DockerJava.ContractEngine.Port = nodeInfo.NodeRpcPort + 11050
	} else {
		conf.VmConf.DockerJava.Enable = false
	}

	javaLogPath := conf.VmConf.DockerJava.LogMountPath
	conf.VmConf.DockerJava.LogMountPath = strings.Replace(javaLogPath, "{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	javaMountPath := conf.VmConf.DockerJava.DataMountPath
	conf.VmConf.DockerJava.DataMountPath = strings.Replace(javaMountPath, "{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	logPath := conf.VmConf.DockerGo.LogMountPath
	conf.VmConf.DockerGo.LogMountPath = strings.Replace(logPath, "{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	mountPath := conf.VmConf.DockerGo.DataMountPath
	conf.VmConf.DockerGo.DataMountPath = strings.Replace(mountPath, "{org_id}", orgId+"-"+nodeInfo.NodeName, -1)

	//containerName := conf.VmConf.DockervmContainerName
	//conf.VmConf.DockervmContainerName = strings.Replace(containerName,
	//"dockervm_container_name", "chainmaker-vm-docker-go-container"+orgId+"-"+nodeInfo.NodeName, -1)

	if chainInfo.NodeFastSyncEnabled != NO_FAST_SYNC {
		conf.NodeConf.FastSync = &config.FastSyncConf{
			Enabled: true,
		}
	}
	conf.TxpoolConf.MaxTxpoolSize = chainInfo.TxPoolMaxSize
	conf.RpcConf.MaxSendMsgSize = chainInfo.RpcMaxSendMsgSize
	conf.RpcConf.MaxRecvMsgSize = chainInfo.RpcMaxRecvMsgSize

	return nil
}

// createBc
//
//	@Description:
//	@param chainId
//	@param confYml
//	@return map[string]*dbcommon.ChainOrgNode
//	@return error
func createBc(chainId, confYml string) (map[string]*dbcommon.ChainOrgNode, error) {
	err := os.RemoveAll("release/")
	if err != nil {
		log.Error("Remove org path err : " + err.Error())
	}
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("GetChainByChainId err : " + err.Error())
		return nil, err
	}

	chainOrgs, err := relation.GetChainOrgList(chainId)
	if err != nil {
		log.Error("GetChainOrgList err : " + err.Error())
		return nil, err
	}

	trustList := []*config.TrustRootsConf{}
	nodeList := []*config.NodesConf{}
	nodeIdMap := map[string]*dbcommon.ChainOrgNode{}

	for _, chainOrg := range chainOrgs {
		var chainOrgNodes []*dbcommon.ChainOrgNode
		chainOrgNodes, err = relation.GetChainOrg(chainOrg.OrgId, chainId)
		if err != nil {
			log.Error("GetChainOrg err : " + err.Error())
			return nil, err
		}
		nodes := config.NodesConf{}
		nodes.OrgId = chainOrg.OrgId

		for _, orgNode := range chainOrgNodes {
			nodeInfo, nodeInfoErr := chain_participant.GetConsensusNodeByNodeName(orgNode.NodeName)
			if nodeInfoErr != nil {
				log.Error("GetConsensusNodeByNodeName err : " + nodeInfoErr.Error())
				continue
			}
			if nodeInfo.Type == chain_participant.NODE_CONSENSUS {
				nodes.NodeId = append(nodes.NodeId, nodeInfo.NodeId)
				nodeIdMap[nodeInfo.NodeId] = orgNode
			}
		}
		nodeList = append(nodeList, &nodes)
	}

	bcConf := new(config.Bc)
	bcFile, err := ioutil.ReadFile(confYml + "/config_tpl/chainconfig/bc1.yml")
	if err != nil {
		log.Error(err.Error())
	}
	_ = yaml.Unmarshal(bcFile, bcConf)
	bcConf.Crypto.Hash = chainInfo.CryptoHash
	bcConf.ChainId = chainId
	bcConf.Block.TxTimeout = chainInfo.TxTimeout
	bcConf.Block.BlockTxCapacity = chainInfo.BlockTxCapacity
	bcConf.Block.BlockInterval = int(chainInfo.BlockInterval)
	bcConf.Block.BlockTimestampVerify = chainInfo.BlockTxTimestampVerify != 0
	bcConf.Consensus.Nodes = nodeList
	bcConf.TrustRoots = trustList
	bcConf.Consensus.Type = consensus.ConsensusType_value[chainInfo.Consensus]
	if chainInfo.CryptoHash != "" {
		bcConf.Crypto.Hash = chainInfo.CryptoHash
	}
	bcConf.Block.TxTimestampVerify = chainInfo.BlockTxTimestampVerify == NO_BLOCK_TX_TIMESTAMP_VERIFE
	bcConf.Core.TxSchedulerTimeout = chainInfo.CoreTxSchedulerTimeout

	var resourcePolicyConf []*config.ResourcePolicyConf
	_ = json.Unmarshal([]byte(chainInfo.ResourcePolicies), &resourcePolicyConf)
	if len(resourcePolicyConf) > 0 {
		bcConf.ResourcePolicies = resourcePolicyConf
	}

	chainOrgNodes, err := relation.GetChainOrgByChainIdList(chainId)
	if err != nil {
		log.Error("GetChainOrgNode err : " + err.Error())
		return nil, err
	}
	err = removeBc(bcConf, chainOrgNodes, trustList, chainOrgs)
	if err != nil {
		log.Error("removeBc err : " + err.Error())
		return nil, err
	}
	return nodeIdMap, nil
}

// removeBc
//
//	@Description:
//	@param bcConf
//	@param chainOrgNodes
//	@param trustList
//	@param chainOrgs
//	@return error
func removeBc(bcConf *config.Bc, chainOrgNodes []*dbcommon.ChainOrgNode, trustList []*config.TrustRootsConf,
	chainOrgs []*dbcommon.ChainOrg) error {
	for _, chainOrgNode := range chainOrgNodes {
		trustList = trustList[0:0]
		for _, orgInfo := range chainOrgs {
			trust := config.TrustRootsConf{}
			trust.OrgId = orgInfo.OrgId
			trust.Root = []string{"../config/" + chainOrgNode.OrgId + "/certs/ca/" + orgInfo.OrgId + "/ca.crt"}
			trustList = append(trustList, &trust)
		}
		bcConf.TrustRoots = trustList
		bcBytes, _ := yaml.Marshal(bcConf)
		bc1, err := os.Create("bc1.yml")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		defer func() {
			err = bc1.Close()
			if err != nil {
				log.Error("close file err :", err.Error())
			}
		}()

		_, err = bc1.Write(bcBytes)
		if err != nil {
			log.Error(err.Error())
		}
		err = os.MkdirAll("release/"+chainOrgNode.OrgId+"-"+chainOrgNode.NodeName+
			"/config/"+chainOrgNode.OrgId+"/chainconfig", os.ModePerm)
		if err != nil {
			log.Error("Mkdir bc1 path err : " + err.Error())
		}
		err = os.Rename("bc1.yml", "release/"+chainOrgNode.OrgId+"-"+chainOrgNode.NodeName+
			"/config/"+chainOrgNode.OrgId+"/chainconfig"+"/bc1.yml")
		if err != nil {
			log.Error("Rename bc1.yml err : " + err.Error())
		}
	}
	return nil
}
