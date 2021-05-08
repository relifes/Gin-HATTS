package controller

import (
	"Gin-HATTS/common"
	"Gin-HATTS/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// 验证用户是否存在
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(ctx *gin.Context) {
	db := common.GetDB()
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
	newUser := model.User{
		Name: name,
		Email: email,
		Password: password,
	}
	db.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
