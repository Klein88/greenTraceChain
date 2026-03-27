package chain_management

import "management_backend/src/global"

// ResourceName resource map
var ResourceName = map[int]string{
	0: "CHAIN_CONFIG-NODE_ID_UPDATE",
	1: "CHAIN_CONFIG-TRUST_ROOT_ADD",
	2: "CHAIN_CONFIG-CERTS_FREEZE",
	3: "CHAIN_CONFIG-BLOCK_UPDATE",
	4: "CONTRACT_MANAGE-INIT_CONTRACT",
	5: "CONTRACT_MANAGE-UPGRADE_CONTRACT",
	6: "CONTRACT_MANAGE-FREEZE_CONTRACT",
	7: "CONTRACT_MANAGE-UNFREEZE_CONTRACT",
	8: "CONTRACT_MANAGE-REVOKE_CONTRACT",
}

// ResourceNameType resource type map
var ResourceNameType = map[string]int{
	"CHAIN_CONFIG-NODE_ID_UPDATE":       0,
	"CHAIN_CONFIG-TRUST_ROOT_ADD":       1,
	"CHAIN_CONFIG-CERTS_FREEZE":         2,
	"CHAIN_CONFIG-BLOCK_UPDATE":         3,
	"CONTRACT_MANAGE-INIT_CONTRACT":     4,
	"CONTRACT_MANAGE-UPGRADE_CONTRACT":  5,
	"CONTRACT_MANAGE-FREEZE_CONTRACT":   6,
	"CONTRACT_MANAGE-UNFREEZE_CONTRACT": 7,
	"CONTRACT_MANAGE-REVOKE_CONTRACT":   8,
}

// ChainModes chain mode map
var ChainModes = map[string]string{
	global.PUBLIC:               "公钥模式",
	global.PERMISSIONEDWITHCERT: "证书模式",
}
