package common

import (
	"Gin-HATTS/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "hatts"
	username := "root"
	password := "456852Ss"
	charset := "utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to open database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
}

func GetDB() *gorm.DB{
	return DB
}
