package app

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf *Config

type EthereumConfig struct {
	SepoliaRpcUrl     string `mapstructure:"SEPOLIA_RPC_URL"`
	SepoliaPrivateKey string `mapstructure:"SEPOLIA_PRIVATE_KEY"`
	PrivateKey        string `mapstructure:"PRIVATE_KEY"`
}
type Config struct {
	Ethereum EthereumConfig `mapstructure:"ethereum"`
}

func LoadConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}
	Conf = &Config{}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}
	return nil
}
