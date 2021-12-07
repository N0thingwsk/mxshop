package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type UserSrvConfig struct {
	UserServiceName string       `mapstructure:"name" json:"name"`
	Mysql           MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Consul          ConsulConfig `mapstructure:"consul" json:"consul"`
}

var UserConfig *UserSrvConfig
