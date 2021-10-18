package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"mxshop/usersrv/config"
)

func InitConfig() error {
	v := viper.New()
	v.SetConfigFile("usersrv/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&config.UserConfig)
	if err != nil {
		return err
	}
	fmt.Println(config.UserConfig)
	return nil
}
