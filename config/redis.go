package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // Server address:port
	Password string `mapstructure:"password" json:"password" yaml:"password"` // password
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // database of redis
}
