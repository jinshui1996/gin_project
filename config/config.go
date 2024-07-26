package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql  MySQLConfig  `mapstructure:"database" json:"database" yaml:"database"`
	Server ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
	Jwt    JwtConfig	`mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

var EnvConfig *Config

func LoadConfig() *Config {

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config") // 设置config.yaml的路径

	if err := viper.ReadInConfig(); err != nil { // 读取config.yaml中的内容
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}

func init() {
	EnvConfig = LoadConfig()
}
