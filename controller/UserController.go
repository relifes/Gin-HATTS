package controller

import (
	"Gin-HATTS/common"
	"Gin-HATTS/model"
	"Gin-HATTS/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	db.Create(&newUser)

	//返回结果
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	log.Println(email, password)
	// 判断账号是否存在
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) ; err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil,"密码错误")
		return
	}
	// 发放token
	token := "11"
	//返回结果
	response.Success(ctx, gin.H{"token":token}, "登录成功")
}