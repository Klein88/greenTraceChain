package overview

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"management_backend/src/config"
	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/utils"
)

// DownloadSdkHandler downloadSdkHandler
type DownloadSdkHandler struct {
}

// LoginVerify login verify
func (handler *DownloadSdkHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *DownloadSdkHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindDownloadSdkConfigHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	confYml := global.GetConfYml()
	err := createSdkConfig(params, confYml)
	if err != nil {
		log.Error("createSdkConfig err : " + err.Error())
	}
	content, err := ioutil.ReadFile(params.ChainId + ".zip")
	if err != nil {
		log.Error("ReadFile err : " + err.Error())
	}
	defer func() {
		err = os.RemoveAll(params.ChainId + ".zip")
		if err != nil {
			log.Error("remove zip err :", err.Error())
		}
	}()
	fileName := params.ChainId + ".zip"

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Disposition", "attachment; filename="+utils.Base64Encode([]byte(fileName)))
	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition")
	_, err = ctx.Writer.Write(content)
	if err != nil {
		log.Error("ctx Write content err :", err.Error())
	}
}

// createSdkConfig
func createSdkConfig(param *DownloadSDKConfigParams, confYml string) error {
	chainId := param.ChainId
	err := os.RemoveAll("sdk_configs/")
	if err != nil {
		log.Error("Remove org path err : " + err.Error())
	}
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("GetChainByChainId err:" + err.Error())
		return err
	}
	var sdkConfigByte []byte
	if chainInfo.ChainMode == global.PUBLIC {
		sdkConfig := new(config.SdkPkConfig)
		sdkConfigFile, fileErr := ioutil.ReadFile(confYml + "/config_sdk_tpl/sdk_config.yml")
		if fileErr != nil {
			log.Error(fileErr.Error())
		}
		_ = yaml.Unmarshal(sdkConfigFile, sdkConfig)
		sdkConfig.ChainClient.ChainId = chainInfo.ChainId
		if !reflect.DeepEqual(param.MySqlInfo, MySqlInfo{}) {
			sdkConfig.ChainClient.Archive.Dest = fmt.Sprintf("%s:%s:%s:%s",
				param.MySqlInfo.Username, param.MySqlInfo.PassWord, param.MySqlInfo.HostName, param.MySqlInfo.Port)
		}
		sdkConfig.ChainClient.AuthType = chainInfo.ChainMode
		sdkConfig.ChainClient.Crypto.Hash = chainInfo.CryptoHash
		err = createPkAccount(chainId, sdkConfig)
		if err != nil {
			log.Error("Rename sdk_config.yml err : " + err.Error())
		}
		sdkConfigByte, _ = yaml.Marshal(sdkConfig)
	} else {
		sdkConfig := new(config.SdkConfig)
		sdkConfigFile, fileErr := ioutil.ReadFile(confYml + "/config_sdk_tpl/sdk_config.yml")
		if fileErr != nil {
			log.Error(fileErr.Error())
		}
		_ = yaml.Unmarshal(sdkConfigFile, sdkConfig)
		sdkConfig.ChainClient.ChainId = chainInfo.ChainId
		if !reflect.DeepEqual(param.MySqlInfo, MySqlInfo{}) {
			sdkConfig.ChainClient.Archive.Dest = fmt.Sprintf("%s:%s:%s:%s",
				param.MySqlInfo.Username, param.MySqlInfo.PassWord, param.MySqlInfo.HostName, param.MySqlInfo.Port)
		}
		err = createCert(chainId, chainInfo.TLS, sdkConfig)
		if err != nil {
			log.Error("Rename sdk_config.yml err : " + err.Error())
		}
		sdkConfigByte, _ = yaml.Marshal(sdkConfig)
	}
	sdkConfig1, err := os.Create("sdk_config.yml")
	defer func() {
		err = sdkConfig1.Close()
		if err != nil {
			log.Error("close file err :", err.Error())
		}
	}()

	_, err = sdkConfig1.Write(sdkConfigByte)
	if err != nil {
		log.Error(err.Error())
	}
	err = os.MkdirAll("sdk_configs", os.ModePerm)
	if err != nil {
		log.Error("Mkdir config path err : " + err.Error())
	}
	err = os.Rename("sdk_config.yml", "sdk_configs/sdk_config.yml")
	if err != nil {
		log.Error("Rename sdk_config.yml err : " + err.Error())
	}
	err = utils.Zip("sdk_configs", chainId+".zip")
	if err != nil {
		log.Error("zip file err :", err.Error())
	}
	return nil
}

// createPkAccount
func createPkAccount(chainId string, sdkConfig *config.SdkPkConfig) error {
	chainOrgNodes, err := relation.GetChainOrgByChainIdList(chainId)
	if err != nil {
		log.Error("GetChainOrgByChainIdList err : " + err.Error())
		return err
	}
	admins, err := relation.GetChainUserByChainId(chainId, "")
	if err != nil {
		log.Error("GetChainUserByChainId err : " + err.Error())
		return err
	}
	chainSubscribeConfig, err := chain.GetChainSubscribeByChainId(chainId)
	if err != nil {
		log.Error("getSubscribeConfig err : " + err.Error())
		return err
	}
	nodes := make([]config.PkSdkNodeConf, 0)
	for _, node := range chainOrgNodes {
		nodePath := "sdk_configs/crypto-config/" + node.NodeName
		err = os.MkdirAll(nodePath, os.ModePerm)
		if err != nil {
			log.Error("mkdirAll err : " + err.Error())
		}
		cert, err := chain_participant.GetPemCert(node.NodeName)
		if err != nil {
			log.Error("GetPemCert err : " + err.Error())
			return err
		}
		err = createPkNodeFile(node.NodeName, cert.PublicKey, cert.PrivateKey, node.NodeId, nodePath)
		if err != nil {
			log.Error(fmt.Sprintf("create node cert err : %v, orgId:%v", err, node.NodeName))
		}
		adminPath := "sdk_configs/crypto-config/" + node.NodeName + "/admin"
		err = os.MkdirAll(adminPath, os.ModePerm)
		if err != nil {
			log.Error("mkdirAll err : " + err.Error())
		}
		address := node.NodeIp + ":" + strconv.Itoa(node.NodeRpcPort)
		for _, admin := range admins {
			userPath := adminPath + "/" + admin.UserName
			err = os.MkdirAll(userPath, os.ModePerm)
			if err != nil {
				log.Error("mkdirAll err : " + err.Error())
			}
			adminInfo, err := chain_participant.GetPemCert(admin.UserName)
			if err != nil {
				log.Error("GetPemCert err : " + err.Error())
				return err
			}
			err = createPkAdminFile(admin.UserName, adminInfo.PublicKey, adminInfo.PrivateKey, adminInfo.Addr, userPath)
			if err != nil {
				log.Error(fmt.Sprintf("create node cert err : %v, orgId:%v", err, node.NodeName))
			}
			if address == chainSubscribeConfig.NodeRpcAddress && admin.UserName == chainSubscribeConfig.AdminName {
				sdkConfig.ChainClient.UserSignKeyFilePath = "./crypto-config/" +
					node.NodeName + "/admin/" + admin.UserName + "/" + admin.UserName + ".key"
			}
		}
		nodes = append(nodes, config.PkSdkNodeConf{
			NodeAddr: address,
			ConnCnt:  10,
		})
	}
	sdkConfig.ChainClient.Nodes = nodes
	return nil
}

// createCert
func createCert(chainId string, tls int, sdkConfig *config.SdkConfig) error {
	chainOrgs, err := relation.GetChainOrgList(chainId)
	if err != nil {
		log.Error("GetChainOrgByChainIdList err : " + err.Error())
		return err
	}
	chainSubscribeConfig, err := chain.GetChainSubscribeByChainId(chainId)
	if err != nil {
		log.Error("getSubscribeConfig err : " + err.Error())
		return err
	}
	for _, org := range chainOrgs {
		orgPath := "sdk_configs/crypto-config/" + org.OrgId + "/ca"
		err = os.MkdirAll(orgPath, os.ModePerm)
		if err != nil {
			log.Error("mkdirAll err : " + err.Error())
		}
		orgCert, err := chain_participant.GetOrgCaCert(org.OrgId)
		if err != nil {
			log.Error("GetOrgCaCert err : " + err.Error())
			return err
		}
		err = createCertFile("ca", orgCert.Cert, orgCert.PrivateKey, orgPath)
		if err != nil {
			log.Error(fmt.Sprintf("create org cert err : %v, orgId:%v", err, org.OrgId))
		}
		err = createNodeCert(org.OrgId, chainId)
		if err != nil {
			log.Error(fmt.Sprintf("create node cert err : %v, orgId:%v", err, org.OrgId))
		}
		err = createUserCert(org.OrgId, chainSubscribeConfig.UserName, sdkConfig)
		if err != nil {
			log.Error(fmt.Sprintf("create user cert err : %v, orgId:%v", err, org.OrgId))
		}
		if org.OrgId == chainSubscribeConfig.OrgId {
			sdkConfig.ChainClient.OrgId = org.OrgId
			sdkConfig.ChainClient.Nodes[0].NodeAddr = chainSubscribeConfig.NodeRpcAddress
			sdkConfig.ChainClient.Nodes[0].TrustRootPaths[0] = "./crypto-config/" + org.OrgId + "/ca"
			if tls != CHAINSTARTTLS {
				sdkConfig.ChainClient.Nodes[0].EnableTls = false
			} else {
				sdkConfig.ChainClient.Nodes[0].EnableTls = true
				certInfo, err := utils.ParseCertificate([]byte(orgCert.Cert))
				if err != nil {
					log.Error(fmt.Sprintf("parse org cert err : %v, orgId:%v", err, org.OrgId))
				}
				for _, name := range certInfo.DNSNames {
					if name != LOCALHOST {
						sdkConfig.ChainClient.Nodes[0].TlsHostName = name
						break
					}
				}
			}
		}
	}
	return nil
}

// createNodeCert
func createNodeCert(orgId, chainId string) error {
	chainOrgNodes, err := relation.GetChainOrgByChainIdList(chainId)
	if err != nil {
		log.Error("GetChainOrgNode err : " + err.Error())
	}
	err = os.MkdirAll("sdk_configs/crypto-config/"+orgId+"/node", os.ModePerm)
	if err != nil {
		log.Error("Mkdir org certs/node path err : " + err.Error())
	}
	for _, node := range chainOrgNodes {
		nodeName := node.NodeName
		path := "sdk_configs/crypto-config/" + orgId + "/node/" + nodeName
		err = os.MkdirAll(path, os.ModePerm)
		nodeCertList, nodeErr := chain_participant.GetNodeCert(nodeName)
		if nodeErr != nil {
			log.Error("Mkdir org certs/node path err : " + nodeErr.Error())
		}
		for _, nodeCert := range nodeCertList {
			if nodeCert.CertUse == global.SIGN {
				err = createCertFile(nodeName+".sign", nodeCert.Cert, nodeCert.PrivateKey, path)
				if err != nil {
					log.Error(fmt.Sprintf("create node cert err : %v, nodeId:%v, orgId:%v, certUse:%v",
						err, node.NodeId, orgId, nodeCert.CertUserName))
				}
			} else if nodeCert.CertUse == global.TLS {
				err = createCertFile(nodeName+".tls", nodeCert.Cert, nodeCert.PrivateKey, path)
				if err != nil {
					log.Error(fmt.Sprintf("create node cert err : %v, nodeId:%v, "+
						"orgId:%v, certUse:%v", err, node.NodeId, orgId, nodeCert.CertUserName))
				}
			}
		}
	}
	return err
}

// createUserCert
func createUserCert(orgId string, subscribeUserName string, sdkConfig *config.SdkConfig) error {
	userCertList, _, err := chain_participant.GetUserCertList(orgId)
	if err != nil {
		log.Error("GetChainOrgNode err : " + err.Error())
	}
	err = os.MkdirAll("sdk_configs/crypto-config/"+orgId+"/user", os.ModePerm)
	if err != nil {
		log.Error("Mkdir org certs/node path err : " + err.Error())
	}
	for _, user := range userCertList {
		userName := user.CertUserName
		path := "sdk_configs/crypto-config/" + orgId + "/user/" + userName
		err = os.MkdirAll(path, os.ModePerm)
		if user.CertUse == global.SIGN {
			if userName == subscribeUserName {
				sdkConfig.ChainClient.UserSignKeyFilePath = "./crypto-config/" + orgId +
					"/user/" + userName + "/" + userName + ".sign.key"
				sdkConfig.ChainClient.UserSignCrtFilePath = "./crypto-config/" + orgId +
					"/user/" + userName + "/" + userName + ".sign.crt"
			}
			err = createCertFile(userName+".sign", user.Cert, user.PrivateKey, path)
			if err != nil {
				log.Error(fmt.Sprintf("create user cert err : %v, username:%v, orgId:%v, "+
					"certUse:%v", err, userName, orgId, user.CertUserName))
			}

		} else if user.CertUse == global.TLS {
			if userName == subscribeUserName {
				sdkConfig.ChainClient.UserKeyFilePath = "./crypto-config/" + orgId + "/user/" +
					userName + "/" + userName + ".tls.key"
				sdkConfig.ChainClient.UserCrtFilePath = "./crypto-config/" + orgId + "/user/" +
					userName + "/" + userName + ".tls.crt"
			}
			err = createCertFile(userName+".tls", user.Cert, user.PrivateKey, path)
			if err != nil {
				log.Error(fmt.Sprintf("create user cert err : %v, username:%v, "+
					"orgId:%v, certUse:%v", err, userName, orgId, user.CertUserName))
			}
		}
	}
	return err
}
