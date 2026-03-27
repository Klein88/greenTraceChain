/*
Package config 配置
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package config

// Chainmaker 整体配置
type Chainmaker struct {
	AuthType       string            `yaml:"auth_type"`
	ChainLogConf   *ChainLogConf     `yaml:"log"`
	CryptoEngine   string            `yaml:"crypto_engine"`
	BlockchainConf []*BlockchainConf `yaml:"blockchain"`
	NodeConf       *NodeConf         `yaml:"node"`
	NetConf        *NetConf          `yaml:"net"`
	TxpoolConf     *TxpoolConf       `yaml:"txpool"`
	RpcConf        *RpcConf          `yaml:"rpc"`
	TxFilter       *TxFilterConfig   `yaml:"tx_filter"`
	MonitorConf    *MonitorConf      `yaml:"monitor"`
	PprofConf      *PprofConf        `yaml:"pprof"`
	Consensus      *ConsensusConf    `yaml:"consensus"`
	StorageConf    *StorageConf      `yaml:"storage"`
	SchedulerConf  *SchedulerConf    `yaml:"scheduler"`
	VmConf         *VmConf           `yaml:"vm"`
	//CoreConf       *CoreConf         `yaml:"core"`
}

// PKChainmaker pk 整体配置
type PKChainmaker struct {
	AuthType       string            `yaml:"auth_type"`
	ChainLogConf   *ChainLogConf     `yaml:"log"`
	CryptoEngine   string            `yaml:"crypto_engine"`
	BlockchainConf []*BlockchainConf `yaml:"blockchain"`
	NodeConf       *NodePkConf       `yaml:"node"`
	NetConf        *NetConf          `yaml:"net"`
	TxpoolConf     *TxpoolConf       `yaml:"txpool"`
	RpcConf        *RpcConf          `yaml:"rpc"`
	TxFilter       *TxFilterConfig   `yaml:"tx_filter"`
	MonitorConf    *MonitorConf      `yaml:"monitor"`
	PprofConf      *PprofConf        `yaml:"pprof"`
	Consensus      *ConsensusConf    `yaml:"consensus"`
	StorageConf    *StorageConf      `yaml:"storage"`
	SchedulerConf  *SchedulerConf    `yaml:"scheduler"`
	VmConf         *VmConf           `yaml:"vm"`
	//CoreConf       *CoreConf         `yaml:"core"`
}

// ChainLogConf log config
type ChainLogConf struct {
	ConfigFile string `yaml:"config_file"`
}

// BlockchainConf block config
type BlockchainConf struct {
	ChainId string `yaml:"chainId"`
	Genesis string `yaml:"genesis"`
}

// NodeConf node config
type NodeConf struct {
	OrgId             string        `yaml:"org_id"`
	PrivKeyFile       string        `yaml:"priv_key_file"`
	CertFile          string        `yaml:"cert_file"`
	CertCacheSize     int           `yaml:"cert_cache_size"`
	CertKeyUsageCheck bool          `yaml:"cert_key_usage_check"`
	Pkcs11            *Pkcs11Conf   `yaml:"pkcs11"`
	FastSync          *FastSyncConf `yaml:"fast_sync"`
}

// NodePkConf node pk config
type NodePkConf struct {
	PrivKeyFile   string        `yaml:"priv_key_file"`
	CertCacheSize int           `yaml:"cert_cache_size"`
	Pkcs11        *Pkcs11Conf   `yaml:"pkcs11"`
	FastSync      *FastSyncConf `yaml:"fast_sync"`
}

// NetConf net config
type NetConf struct {
	Provider   string   `yaml:"provider"`
	ListenAddr string   `yaml:"listen_addr"`
	Seeds      []string `yaml:"seeds"`
	Tls        *TlsConf `yaml:"tls"`
}

// TlsConf tls config
type TlsConf struct {
	Enabled     bool   `yaml:"enabled"`
	PrivKeyFile string `yaml:"priv_key_file"`
	CertFile    string `yaml:"cert_file"`
}

// TxpoolConf tx pool config
type TxpoolConf struct {
	PoolType            string `yaml:"pool_type"`
	MaxTxpoolSize       int    `yaml:"max_txpool_size"`
	MaxConfigTxpoolSize int    `yaml:"max_config_txpool_size"`
	IsDumpTxsInQueue    bool   `yaml:"is_dump_txs_in_queue"`
	CommonQueueNum      int    `yaml:"common_queue_num"`
	BatchMaxSize        int    `yaml:"batch_max_size"`
	BatchCreateTimeout  int    `yaml:"batch_create_timeout"`
}

// RpcConf rpc config
type RpcConf struct {
	Provider                               string          `yaml:"provider"`
	Port                                   int             `yaml:"port"`
	CheckChainConfTrustRootsChangeInterval int             `yaml:"check_chain_conf_trust_roots_change_interval"`
	Ratelimit                              *RatelimitConf  `yaml:"ratelimit"`
	Subscriber                             *SubscriberConf `yaml:"subscriber"`
	Tls                                    *RpcTlsConf     `yaml:"tls"`
	MaxSendMsgSize                         int             `yaml:"max_send_msg_size"`
	MaxRecvMsgSize                         int             `yaml:"max_recv_msg_size"`
	//BlackList                            *BlackListConf  `yaml:"blacklist"`
	GatewayConf *GatewayConf `yaml:"gateway"`
}

// RpcTlsConf rpc tls config
type RpcTlsConf struct {
	Mode        string `yaml:"mode"`
	PrivKeyFile string `yaml:"priv_key_file"`
	CertFile    string `yaml:"cert_file"`
}

// GatewayConf gateway  config
type GatewayConf struct {
	Enabled         bool `yaml:"enabled"`
	MaxRespBodySize int  `yaml:"max_resp_body_size"`
}

// BlackListConf black list config
type BlackListConf struct {
	Addresses []string `yaml:"addresses"`
}

// SubscriberConf subscriber config
type SubscriberConf struct {
	Ratelimit *RatelimitConf `yaml:"ratelimit"`
}

// RatelimitConf rate limit config
type RatelimitConf struct {
	Enabled         bool `yaml:"enabled"`
	Type            int  `yaml:"type"`
	TokenPerSecond  int  `yaml:"token_per_second"`
	TokenBucketSize int  `yaml:"token_bucket_size"`
}

// MonitorConf monitor config
type MonitorConf struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

// PprofConf pprof config
type PprofConf struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

// ConsensusConf  consensus config
type ConsensusConf struct {
	Raft *RaftConf `yaml:"raft"`
}

// RaftConf raft config
type RaftConf struct {
	SnapCount    int  `yaml:"snap_count"`
	AsyncWalSave bool `yaml:"async_wal_save"`
	Ticker       int  `yaml:"ticker"`
}

// StorageConf storage config
type StorageConf struct {
	StorePath                  string                `yaml:"store_path"`
	UnarchiveBlockHeight       int64                 `yaml:"unarchive_block_height"`
	DisableBlockFileDb         bool                  `yaml:"disable_block_file_db"`
	LogdbSegmentAsync          bool                  `yaml:"logdb_segment_async"`
	LogdbSegmentSize           int                   `yaml:"logdb_segment_size"`
	EnableBigfilter            bool                  `yaml:"enable_bigfilter"`
	BigfilterConfig            *BigfilterConfig      `yaml:"bigfilter_config"`
	RollingWindowCacheCapacity int                   `yaml:"rolling_window_cache_capacity"`
	WriteBlockType             int                   `yaml:"write_block_type"`
	StateCacheConfig           *StorageCacheConf     `yaml:"state_cache_config"`
	BlockdbConfig              *StorageDbConf        `yaml:"blockdb_config"`
	StatedbConfig              *StorageDbConf        `yaml:"statedb_config"`
	HistorydbConfig            *StorageHistoryDbConf `yaml:"historydb_config"`
	ResultdbConfig             *StorageDbConf        `yaml:"resultdb_config"`
	DisableContractEventdb     bool                  `yaml:"disable_contract_eventdb"`
	ContractEventdbConfig      *StorageEventDbConf   `yaml:"contract_eventdb_config"`
}

// BigfilterConfig big filter config
type BigfilterConfig struct {
	RedisHostsPort string  `yaml:"redis_hosts_port"`
	RedisPassword  string  `yaml:"redis_password"`
	TxCapacity     int64   `yaml:"tx_capacity"`
	FpRate         float64 `yaml:"fp_rate"`
}

// StorageEventDbConf storage event db config
type StorageEventDbConf struct {
	Provider    string            `yaml:"provider"`
	SqldbConfig *StorageSqlDbConf `yaml:"sqldb_config"`
}

// StorageSqlDbConf storage sql db config
type StorageSqlDbConf struct {
	SqldbType string `yaml:"sqldb_type"`
	Dsn       string `yaml:"dsn"`
}

// StorageCacheConf storage cache config
type StorageCacheConf struct {
	LifeWindow       int64 `yaml:"life_window"`
	CleanWindow      int64 `yaml:"clean_window"`
	MaxEntrySize     int   `yaml:"max_entry_size"`
	HardMaxCacheSize int   `yaml:"hard_max_cache_size"`
}

// StorageHistoryDbConf storage history db config
type StorageHistoryDbConf struct {
	Provider               string         `yaml:"provider"`
	DisableKeyHistory      bool           `yaml:"disable_key_history"`
	DisableContractHistory bool           `yaml:"disable_contract_history"`
	DisableAccountHistory  bool           `yaml:"disable_account_history"`
	LeveldbConfig          *LevelDbDbConf `yaml:"leveldb_config"`
}

// StorageDbConf storage db config
type StorageDbConf struct {
	Provider      string         `yaml:"provider"`
	LeveldbConfig *LevelDbDbConf `yaml:"leveldb_config"`
}

// LevelDbDbConf level db config
type LevelDbDbConf struct {
	StorePath string `yaml:"store_path"`
}

// Pkcs11Conf pkcs11 config
type Pkcs11Conf struct {
	Enabled          bool   `yaml:"enabled"`
	Type             string `yaml:"type"`
	Library          string `yaml:"library"`
	Label            string `yaml:"label"`
	Password         string `yaml:"password"`
	SessionCacheSize int    `yaml:"session_cache_size"`
	Hash             string `yaml:"hash"`
}

// FastSyncConf fast sync config
type FastSyncConf struct {
	Enabled       bool `yaml:"enabled"`
	MinFullBlocks int  `yaml:"min_full_blocks"`
}

// CoreConf core config
type CoreConf struct {
	Evidence bool `yaml:"evidence"`
}

// SchedulerConf scheduler config
type SchedulerConf struct {
	RwsetLog bool `yaml:"rwset_log"`
}

// VmConf vm config
type VmConf struct {
	Common     *Common     `yaml:"common"`
	DockerGo   *DockerGo   `yaml:"go"`
	DockerJava *DockerJava `yaml:"java"`
}

// Common 链配置
type Common struct {
	MaxSendMsgSize int                `yaml:"max_send_msg_size"`
	MaxRecvMsgSize int                `yaml:"max_recv_msg_size"`
	DialTimeout    int                `yaml:"dial_timeout"`
	RuntimeServer  *RuntimeServerConf `yaml:"runtime_server"`
}

// DockerGo docker go
type DockerGo struct {
	Enable bool `yaml:"enable"`
	//DockervmContainerName string `yaml:"dockervm_container_name"`
	DataMountPath  string              `yaml:"data_mount_path"`
	LogMountPath   string              `yaml:"log_mount_path"`
	Protocol       string              `yaml:"protocol"`
	LogInConsole   bool                `yaml:"log_in_console"`
	LogLevel       string              `yaml:"log_level"`
	MaxConcurrency int                 `yaml:"max_concurrency"`
	ContractEngine *ContractEngineConf `yaml:"contract_engine"`
}

// DockerJava docker java
type DockerJava struct {
	Enable         bool                `yaml:"enable"`
	DataMountPath  string              `yaml:"data_mount_path"`
	LogMountPath   string              `yaml:"log_mount_path"`
	LogInConsole   bool                `yaml:"log_in_console"`
	LogLevel       string              `yaml:"log_level"`
	MaxConcurrency int                 `yaml:"max_concurrency"`
	ContractEngine *ContractEngineConf `yaml:"contract_engine"`
}

// PkVmConf pk vm config
type PkVmConf struct {
	EnableDockervm bool `yaml:"enable_dockervm"`
	//DockervmContainerName string `yaml:"dockervm_container_name"`
	DockervmMountPath string               `yaml:"dockervm_mount_path"`
	DockervmLogPath   string               `yaml:"dockervm_log_path"`
	StartNow          bool                 `yaml:"start_now"`
	LogInConsole      bool                 `yaml:"log_in_console"`
	LogLevel          string               `yaml:"log_level"`
	UdsOpen           bool                 `yaml:"uds_open"`
	RuntimeServer     *PKRuntimeServerConf `yaml:"runtime_server"`
	ContractEngine    *ContractEngineConf  `yaml:"contract_engine"`
}

// RuntimeServerConf runtime server config
type RuntimeServerConf struct {
	Port int `yaml:"port"`
}

// PKRuntimeServerConf  pk runtime server config
type PKRuntimeServerConf struct {
	Port           int `yaml:"port"`
	DialTimeout    int `yaml:"dial_timeout"`
	MaxSendMsgSize int `yaml:"max_send_msg_size"`
	MaxRecvMsgSize int `yaml:"max_recv_msg_size"`
}

// ContractEngineConf contract engine config
type ContractEngineConf struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	MaxConnection int    `yaml:"max_connection"`
}

// TxFilterConfig tx filter config
type TxFilterConfig struct {
	Type      int              `yaml:"type"`
	BirdsNest *BirdsNestConfig `yaml:"birds_nest"`
	Sharding  *Sharding        `yaml:"sharding"`
}

// ShardingBirdsNestConfig sharding birds nest config
type ShardingBirdsNestConfig struct {
	Length uint32        `yaml:"length"`
	Rules  *RulesConfig  `yaml:"rules"`
	Cuckoo *CuckooConfig `yaml:"cuckoo"`
}

// BirdsNestConfig birds nest config
type BirdsNestConfig struct {
	Length   uint32                    `yaml:"length"`
	Rules    *RulesConfig              `yaml:"rules"`
	Cuckoo   *CuckooConfig             `yaml:"cuckoo"`
	Snapshot *SnapshotSerializerConfig `yaml:"snapshot"`
}

// RulesConfig rules config
type RulesConfig struct {
	AbsoluteExpireTime int64 `yaml:"absolute_expire_time"`
}

// CuckooConfig cuckoo config
type CuckooConfig struct {
	KeyType       int    `yaml:"key_type"`
	TagsPerBucket uint32 `yaml:"tags_per_bucket"`
	BitsPerItem   uint32 `yaml:"bits_per_item"`
	MaxNumKeys    uint32 `yaml:"max_num_keys"`
	TableType     uint32 `yaml:"table_type"`
}

// SnapshotSerializerConfig snapshot serializer config
type SnapshotSerializerConfig struct {
	Type              int                      `yaml:"type"`
	Timed             *SerializeIntervalConfig `yaml:"timed"`
	BlockHeight       *SerializeIntervalConfig `yaml:"block_height"`
	SerializeInterval int                      `yaml:"serialize_interval"`
	Path              string                   `yaml:"path"`
}

// SerializeIntervalConfig serialize interval config
type SerializeIntervalConfig struct {
	Interval int64 `yaml:"interval"`
}

// Sharding sharding
type Sharding struct {
	Length    int                       `yaml:"length"`
	Timeout   int                       `yaml:"timeout"`
	Snapshot  *SnapshotSerializerConfig `yaml:"snapshot"`
	BirdsNest *ShardingBirdsNestConfig  `yaml:"birds_nest"`
}
