package config

type MySQLConfig struct {
  Host     string	`mapstructure:"host" json:"host" yaml:"host"`
  Port     string	`mapstructure:"port" json:"port" yaml:"port"`
  Username string	`mapstructure:"user" json:"user" yaml:"user"`
  Password string	`mapstructure:"password" json:"password" yaml:"password"`
  Database string	`mapstructure:"database" json:"database" yaml:"database"`
}