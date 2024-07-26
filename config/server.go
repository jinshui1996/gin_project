package config

type ServerConfig struct {
  Port string `mapstructure:"port" json:"port" yaml:"port"`
}