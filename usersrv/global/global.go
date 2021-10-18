package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mxshop/usersrv/config"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	dsn := config.UserConfig.Mysql.User + ":" + config.UserConfig.Mysql.Password + "@tcp(" + config.UserConfig.Mysql.Host + ")/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("gorm error:" + err.Error())
	}
}
