package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`               // Environmental value
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`            // Port value
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`    // Database type:mysql(default)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"` // Oss type
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // Multi-point login interception
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // Use redis
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"` // Use mongo
}
