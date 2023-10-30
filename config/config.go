package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo  Mongo  `json:"mongo" yaml:"mongo" mapstructure:"mongo"`
	Email  Email  `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	AwsS3 AwsS3 `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	// Cross-domain-origin configuration
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
