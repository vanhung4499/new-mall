package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // To email
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // From email
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // SMTP Server like smtp.gmail.com
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // Secret: The key used for login
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // Nickname of the sender
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // SMTP Port: Google use port 465
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // Whether SSL Whether to enable SSL
}
