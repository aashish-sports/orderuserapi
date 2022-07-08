package user

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "root:Pass@123@tcp(127.0.0.1:3306)/orderuserapi?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected sucessfully")
}
