package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mxshop/usersrv/model"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("gorm error:" + err.Error())
	}
	//_ = db.AutoMigrate(&model.User{})
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode("12345", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	for i := 0; i < 10; i++ {
		user := model.User{
			BaseModel: model.BaseModel{},
			Mobile:    fmt.Sprintf("1810765231%d", i),
			Password:  newPassword,
			NickName:  fmt.Sprintf("hello%d", i),
			Birthday:  nil,
			Gender:    "",
			Role:      0,
		}
		db.Save(&user)
	}
}
