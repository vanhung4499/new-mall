package config

import (
	"github.com/spf13/viper"
	"os"
)

type Server struct {
	System        *System        `yaml:"system"`
	MySql         *MySql         `yaml:"mysql"`
	Email         *Email         `yaml:"email"`
	Redis         *Redis         `yaml:"redis"`
	Token         *Token         `yaml:"token"`
	EncryptSecret *EncryptSecret `yaml:"encryptSecret"`
	Cache         *Cache         `yaml:"cache"`
	Local         *Local         `yaml:"local"`
	AwsS3         *AwsS3         `yaml:"awsS3"`
}

type System struct {
	AppEnv      string `yaml:"appEnv"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"host"`
	UploadModel string `yaml:"uploadModel"`
}

type Token struct {
	AccessTokenExpiry  int `yaml:"accessTokenExpiry"`
	RefreshTokenExpiry int `yaml:"refreshTokenExpiry"`
}

type AwsS3 struct {
	Bucket           string `yaml:"bucket"`
	Region           string `yaml:"region"`
	Endpoint         string `yaml:"endpoint"`
	SecretID         string `yaml:"secretId"`
	SecretKey        string `yaml:"secretKey"`
	BaseURL          string `yaml:"baseUrl"`
	PathPrefix       string `yaml:"pathPrefix"`
	S3ForcePathStyle bool   `yaml:"s3ForcePathStyle"`
	DisableSSL       bool   `yaml:"disableSsl"`
}

type Local struct {
	Path      string `yaml:"path"`      // local file access path
	StorePath string `yaml:"storePath"` // Local file storage path
}

type MySql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Email struct {
	ValidEmail string `yaml:"validEmail"`
	SmtpHost   string `yaml:"smtpHost"`
	SmtpEmail  string `yaml:"smtpEmail"`
	SmtpPass   string `yaml:"smtpPass"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
	MoneySecret string `yaml:"moneySecret"`
}

type Cache struct {
	CacheType    string `yaml:"cacheType"`
	CacheExpires int64  `yaml:"cacheExpires"`
	CacheWarmUp  bool   `yaml:"cacheWarmUp"`
	CacheServer  string `yaml:"cacheServer"`
}

func LoadConfig() *Server {
	var Config Server
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
	return &Config
}
