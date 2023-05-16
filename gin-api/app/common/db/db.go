package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Model *gorm.DB
var err error

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/api?charset=utf8&parseTime=True&loc=Local"
	Model, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}
