package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type UserSrvConfig struct {
	ServerName string
	Mysql      MysqlConfig `mapstructuer:"mysql" json:"mysql"`
}

var UserConfig *UserSrvConfig
