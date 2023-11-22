package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var Config *Server

type Server struct {
	System        *System           `yaml:"system"`
	MySql         map[string]*MySql `yaml:"mysql"`
	Email         *Email            `yaml:"email"`
	Redis         *Redis            `yaml:"redis"`
	EncryptSecret *EncryptSecret    `yaml:"encryptSecret"`
	Cache         *Cache            `yaml:"cache"`
	Local         *Local            `yaml:"local"`
	AwsS3         *AwsS3            `yaml:"awsS3"`
}

type System struct {
	AppEnv      string `yaml:"appEnv"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"host"`
	UploadModel string `yaml:"uploadModel"`
}

type AwsS3 struct {
	Bucket           string `yaml:"bucket"`
	Region           string `yaml:"region"`
	Endpoint         string `yaml:"endpoint"`
	SecretID         string `yaml:"secret-id"`
	SecretKey        string `yaml:"secret-key"`
	BaseURL          string `yaml:"base-url"`
	PathPrefix       string `yaml:"path-prefix"`
	S3ForcePathStyle bool   `yaml:"s3-force-path-style"`
	DisableSSL       bool   `yaml:"disable-ssl"`
}

type Local struct {
	Path      string `yaml:"path"`       // local file access path
	StorePath string `yaml:"store-path"` // Local file storage path
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

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/locales")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}

func GetExpiresTime() int64 {
	if Config.Cache.CacheExpires == 0 {
		return int64(30 * time.Minute) // 30min
	}

	if Config.Cache.CacheExpires == -1 {
		return -1 // Redis.KeepTTL = -1
	}

	return int64(time.Duration(Config.Cache.CacheExpires) * time.Minute)
}
