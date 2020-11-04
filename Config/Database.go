package Config

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"../Models"
)

var dsn = "root@tcp(127.0.0.1:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func ConnectDB(){
	dbase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to connect to database!")
	}
	DB = dbase
	//DB.AutoMigrate(&Models.User{}, &Models.Role{})
}