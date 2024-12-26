package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Node                  NodeConfig      `yaml:"node"`
	Manager               ManagerConfig   `yaml:"manager"`
	Contracts             ContractsConfig `yaml:"contracts"`
	EthStartingHeight     int64           `yaml:"eth_starting_height"`
	BabylonStartingHeight int64           `yaml:"babylon_starting_height"`
	EthBlockStep          uint64          `yaml:"eth_block_step"`
	BabylonBlockStep      uint64          `yaml:"babylon_block_step"`
	BabylonRpc            string          `yaml:"babylon_rpc"`
	EthRpc                string          `yaml:"eth_rpc"`
	EthChainID            uint64          `yaml:"eth_chain_id"`
}

type NodeConfig struct {
	LevelDbFolder    string        `yaml:"level_db_folder"`
	KeyPath          string        `yaml:"key_path"`
	WsAddr           string        `yaml:"ws_addr"`
	SignTimeout      time.Duration `yaml:"sign_timeout"`
	WaitScanInterval time.Duration `yaml:"wait_scan_interval"`
}

type ManagerConfig struct {
	LevelDbFolder string        `yaml:"level_db_folder"`
	SdkRpc        string        `yaml:"sdk_rpc"`
	WsAddr        string        `yaml:"ws_addr"`
	SignTimeout   time.Duration `yaml:"sign_timeout"`
	NodeMembers   string        `yaml:"node_members"`
}

type ContractsConfig struct {
	FrmContractAddress string `yaml:"frm_contract_address"`
	BarContactAddress  string `yaml:"bar_contact_address"`
}

func DefaultConfiguration() *Config {
	return &Config{
		Node: NodeConfig{
			LevelDbFolder:    "node_storage",
			KeyPath:          "key.store",
			WsAddr:           "127.0.0.1:8081",
			SignTimeout:      10,
			WaitScanInterval: 3,
		},
		Manager: ManagerConfig{
			LevelDbFolder: "manager_storage",
			SdkRpc:        "127.0.0.1:9000",
			WsAddr:        "127.0.0.1:8081",
			SignTimeout:   20,
		},
		EthStartingHeight:     1,
		BabylonStartingHeight: 1,
		EthBlockStep:          100,
		BabylonBlockStep:      10,
		EthChainID:            1,
	}
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
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
