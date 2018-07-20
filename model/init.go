package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"

	"fmt"
	"web/conf"
)

var DB *gorm.DB

func init() {
	var err error
	user := conf.Mysql.Key("user").String()
	password := conf.Mysql.Key("password").String()
	dbname := conf.Mysql.Key("dbname").String()
	host := conf.Mysql.Key("host").String()

	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",user,password,host,dbname))
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{})
}
