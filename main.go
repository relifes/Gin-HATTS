package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);not null"`
	Email string `gorm:"varchar(40);not null;unique;"`
	Password string `gorm:"size 255;not null"`
}

func InitDB() *gorm.DB {
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
	log.Println("Database connect success")
	if err != nil {
		panic("failed to open database, err: " + err.Error())
	}
	err1 := db.AutoMigrate(&User{})
	if err1 != nil {
		return nil
	}
	log.Println("table create success")
	return db
}
// 验证用户是否存在
func isEmailExist(db *gorm.DB, email string) bool {
	var user User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		//判断账号是否存在
		if isEmailExist(db, email) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg": "用户已存在",
			})
			return
		}
		//创建用户
		newUser := User{
			Name: name,
			Email: email,
			Password: password,
		}
		db.Create(&newUser)

		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})

	r.Run(":8000")
}


