package chain_management

import (
	"strings"
	"testing"

	"gotest.tools/assert"

	"management_backend/src/config"
	"management_backend/src/db/common"
)

//
// TestCreatePkConfig
//  @Description:
//  @param t
//
func TestCreatePkConfig(t *testing.T) {
	chainId := "chain1"
	var err error
	confYml := "../../../dependence"
	pkConfig := ChainPkConfig{
		Consensus: &config.ConsensusPkBcConf{},
		DposNode:  map[string]string{},
	}
	// 获取chain
	pkConfig.Chain = &common.Chain{
		ChainId:         chainId,
		ChainName:       chainId,
		Consensus:       "TBFT",
		TxTimeout:       600,
		BlockTxCapacity: 100,
		BlockInterval:   10,
		Policy:          "",
		Status:          1,
		Version:         "3.0.0",
		Sequence:        "2",
		CryptoHash:      "SHA256",
		ChainMode:       "public",
	}
	// 获取节点
	pkConfig.Nodes = append(pkConfig.Nodes, &common.ChainOrgNode{
		ChainId:     chainId,
		NodeId:      "testnode",
		NodeName:    "node1",
		NodeIp:      "127.0.0.1",
		NodeRpcPort: 12301,
		NodeP2pPort: 11301,
		Type:        0,
	})
	// 获取用户
	pkConfig.Admins = append(pkConfig.Admins, &common.ChainUser{
		ChainId:  chainId,
		UserName: "admin1",
		Addr:     "addr1",
		Type:     0,
	})
	// 配置基本信息
	pkConfig.AdminCerts = append(pkConfig.AdminCerts, &common.Cert{
		CertType:     2,
		CertUse:      1,
		PrivateKey:   "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIDDxWJUtcN/xlstm/lnsJcM1748AOHk7/e6IF2vLLT4DoAoGCCqGSM49\nAwEHoUQDQgAEGKd2SyrFWAM5KMkAGaKJCPqbvkr0WreEAHusIqMzMm8SkMpzqQZG\nb5+zP/tztffr920bVD7vnbGfJQ3eguKqIA==\n-----END EC PRIVATE KEY-----\n",
		PublicKey:    "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEm1h7rNcDBuRMTTGQvy6OZhSDrMp2\n/+zvMtc7r+fB2QD8gbo0YIPIA563vq8nktbwDKNXuXTDaHCBYCea321Yfg==\n-----END PUBLIC KEY-----\n",
		OrgId:        "",
		OrgName:      "",
		CertUserName: "admin1",
		NodeName:     "node1",
		Algorithm:    0,
		Addr:         "e343",
		RemarkName:   "",
		ChainMode:    "public",
	})

	pkConfig.Seeds = append(pkConfig.Seeds, "/ip4/")
	err = pkConfig.DealBaseInfo()
	assert.Equal(t, err, nil)

	var nodePaths string
	// 创建node对应的文件
	for _, node := range pkConfig.Nodes {
		err = pkConfig.CreateBc(node.NodeName, confYml)
		assert.Equal(t, err, nil)
		err = pkConfig.CreateChainMaker(node, confYml)
		assert.Equal(t, err, nil)
		err = pkConfig.createBinAndLib(node.NodeName, "MFkwEwYHKoZIzj0CAQYIKoZI", confYml)
		assert.Equal(t, err, nil)
		nodePaths = nodePaths + "./chainmaker-v3.0.0-" + node.NodeName + ","
	}
	nodePaths = strings.TrimRight(nodePaths, ",")
	// 创建chainmaker
	_ = pkConfig.createChainmakerAndScript(chainId, confYml, nodePaths)
}
