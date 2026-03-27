package config

// SdkConfig SdkConfig
type SdkConfig struct {
	ChainClient ChainClientConf `yaml:"chain_client"`
}

// SdkPkConfig SdkPkConfig
type SdkPkConfig struct {
	ChainClient PkChainClientConf `yaml:"chain_client"`
}

// PkChainClientConf pk chain client config
type PkChainClientConf struct {
	ChainId             string          `yaml:"chain_id"`
	UserSignKeyFilePath string          `yaml:"user_sign_key_file_path"`
	Crypto              CryptoHashConf  `yaml:"crypto"`
	AuthType            string          `yaml:"auth_type"`
	EnableNormalKey     bool            `yaml:"enable_normal_key"`
	Nodes               []PkSdkNodeConf `yaml:"nodes"`
	Archive             ArchiveConf     `yaml:"archive"`
	RpcClient           RpcClientConf   `yaml:"rpc_client"`
}

// ChainClientConf ChainClientConf
type ChainClientConf struct {
	ChainId             string        `yaml:"chain_id"`
	OrgId               string        `yaml:"org_id"`
	UserKeyFilePath     string        `yaml:"user_key_file_path"`
	UserCrtFilePath     string        `yaml:"user_crt_file_path"`
	UserSignKeyFilePath string        `yaml:"user_sign_key_file_path"`
	UserSignCrtFilePath string        `yaml:"user_sign_crt_file_path"`
	RetryLimit          int           `yaml:"retry_limit"`
	RetryInterval       int           `yaml:"retry_interval"`
	Nodes               []SdkNodeConf `yaml:"nodes"`
	Archive             ArchiveConf   `yaml:"archive"`
	RpcClient           RpcClientConf `yaml:"rpc_client"`
	Pkcs11              Pkcs11Conf    `yaml:"pkcs11"` // 目前和chainmaker里pkcs_11结构相同
}

// CryptoHashConf crypto hash conf
type CryptoHashConf struct {
	Hash string `yaml:"hash"`
}

// PkSdkNodeConf pk sdk node conf
type PkSdkNodeConf struct {
	NodeAddr string `yaml:"node_addr"`
	ConnCnt  int    `yaml:"conn_cnt"`
}

// SdkNodeConf sdk node conf
type SdkNodeConf struct {
	NodeAddr       string   `yaml:"node_addr"`
	ConnCnt        int      `yaml:"conn_cnt"`
	EnableTls      bool     `yaml:"enable_tls"`
	TrustRootPaths []string `yaml:"trust_root_paths"`
	TlsHostName    string   `yaml:"tls_host_name"`
}

// ArchiveConf archive conf
type ArchiveConf struct {
	Type      string `yaml:"type"`
	Dest      string `yaml:"dest"`
	SecretKey string `yaml:"secret_key"`
}

// RpcClientConf rpc client conf
type RpcClientConf struct {
	MaxReceiveMessageSize int `yaml:"max_receive_message_size"`
	MaxSendMessageSize    int `yaml:"max_send_message_size"`
}
