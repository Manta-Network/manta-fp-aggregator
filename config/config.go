package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Node                   NodeConfig      `yaml:"node"`
	Manager                ManagerConfig   `yaml:"manager"`
	Contracts              ContractsConfig `yaml:"contracts"`
	CelestiaConfig         CelestiaConfig  `yaml:"celestia_config"`
	Metrics                MetricsConfig   `yaml:"metrics"`
	EthStartingHeight      int64           `yaml:"eth_starting_height"`
	BabylonStartingHeight  int64           `yaml:"babylon_starting_height"`
	CelestiaStartingHeight int64           `yaml:"celestia_starting_height"`
	EthBlockStep           uint64          `yaml:"eth_block_step"`
	BabylonBlockStep       uint64          `yaml:"babylon_block_step"`
	CelestiaBlockStep      uint64          `yaml:"celestia_block_step"`
	BabylonRpc             string          `yaml:"babylon_rpc"`
	EthRpc                 string          `yaml:"eth_rpc"`
	EthChainID             uint64          `yaml:"eth_chain_id"`
	SymbioticStakeUrl      string          `yaml:"symbiotic_stake_url"`
	MinMantaStakeAmount    string          `yaml:"min_manta_stake_amount"`
	EnableKms              bool            `yaml:"enable_kms"`
}

type NodeConfig struct {
	LevelDbFolder    string        `yaml:"level_db_folder"`
	KeyPath          string        `yaml:"key_path"`
	WsAddr           string        `yaml:"ws_addr"`
	HttpAddr         string        `yaml:"http_addr"`
	SignTimeout      time.Duration `yaml:"sign_timeout"`
	WaitScanInterval time.Duration `yaml:"wait_scan_interval"`
}

type ManagerConfig struct {
	LevelDbFolder         string        `yaml:"level_db_folder"`
	WsAddr                string        `yaml:"ws_addr"`
	HttpAddr              string        `yaml:"http_addr"`
	SignTimeout           time.Duration `yaml:"sign_timeout"`
	FPTimeout             time.Duration `yaml:"fp_timeout"`
	NodeMembers           string        `yaml:"node_members"`
	MaxBabylonOperatorNum uint64        `yaml:"max_babylon_operator_num"`
	OutputSubmitInterval  time.Duration `yaml:"output_submit_interval"`
}

type ContractsConfig struct {
	FrmContractAddress  string `yaml:"frm_contract_address"`
	BarContactAddress   string `yaml:"bar_contact_address"`
	OpFinalityGadgat    string `yaml:"op_finality_gadgat"`
	MsmContractAddress  string `yaml:"msm_contract_address"`
	L2ooContractAddress string `yaml:"l2oo_contract_address"`
}

type CelestiaConfig struct {
	Namespace string        `yaml:"namespace"`
	DaRpc     string        `yaml:"da_rpc"`
	Timeout   time.Duration `yaml:"timeout"`
}

type MetricsConfig struct {
	ListenAddr string `yaml:"listen_addr"`
	ListenPort int    `yaml:"listen_port"`
}

func DefaultConfiguration() *Config {
	return &Config{
		Node: NodeConfig{
			LevelDbFolder:    "node_storage",
			KeyPath:          "key.store",
			WsAddr:           "127.0.0.1:8081",
			HttpAddr:         "127.0.0.1:8082",
			SignTimeout:      10,
			WaitScanInterval: 3,
		},
		Manager: ManagerConfig{
			LevelDbFolder: "manager_storage",
			WsAddr:        "127.0.0.1:8081",
			SignTimeout:   20,
			FPTimeout:     30,
		},
		Metrics: MetricsConfig{
			ListenAddr: "127.0.0.1",
			ListenPort: 8082,
		},
		EthStartingHeight:      1,
		BabylonStartingHeight:  1,
		CelestiaStartingHeight: 1,
		EthBlockStep:           100,
		BabylonBlockStep:       10,
		CelestiaBlockStep:      100,
		EthChainID:             1,
	}
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return DefaultConfiguration(), nil
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
