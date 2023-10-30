package config

import (
	"fmt"
	"strings"
)

type Mongo struct {
	Coll             string       `json:"coll" yaml:"coll" mapstructure:"coll"`                                           // collection name
	Options          string       `json:"options" yaml:"options" mapstructure:"options"`                                  // mongodb options
	Database         string       `json:"database" yaml:"database" mapstructure:"database"`                               // database name
	Username         string       `json:"username" yaml:"username" mapstructure:"username"`                               // username
	Password         string       `json:"password" yaml:"password" mapstructure:"password"`                               // password
	MinPoolSize      uint64       `json:"min-pool-size" yaml:"min-pool-size" mapstructure:"min-pool-size"`                // Minimum connection pool
	MaxPoolSize      uint64       `json:"max-pool-size" yaml:"max-pool-size" mapstructure:"max-pool-size"`                // Maximum connection pool
	SocketTimeoutMs  int64        `json:"socket-timeout-ms" yaml:"socket-timeout-ms" mapstructure:"socket-timeout-ms"`    // socket timeout
	ConnectTimeoutMs int64        `json:"connect-timeout-ms" yaml:"connect-timeout-ms" mapstructure:"connect-timeout-ms"` // Connection timeout
	IsZap            bool         `json:"is-zap" yaml:"is-zap" mapstructure:"is-zap"`                                     // Whether to enable zap log
	Hosts            []*MongoHost `json:"hosts" yaml:"hosts" mapstructure:"hosts"`                                        // Host list
}

type MongoHost struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"` // IP address
	Port string `json:"port" yaml:"port" mapstructure:"port"` // port
}

// Uri .
func (x *Mongo) Uri() string {
	length := len(x.Hosts)
	hosts := make([]string, 0, length)
	for i := 0; i < length; i++ {
		if x.Hosts[i].Host != "" && x.Hosts[i].Port != "" {
			hosts = append(hosts, x.Hosts[i].Host+":"+x.Hosts[i].Port)
		}
	}
	if x.Options != "" {
		return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ","), x.Database, x.Options)
	}
	return fmt.Sprintf("mongodb://%s/%s", strings.Join(hosts, ","), x.Database)
}
