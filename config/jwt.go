package config

type JwtConfig struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expires  uint    `mapstructure:"expires" json:"expires" yaml:"expires"`
}