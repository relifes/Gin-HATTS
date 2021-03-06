package controller

import (
	"Gin-HATTS/common"
	"Gin-HATTS/dto"
	"Gin-HATTS/model"
	"Gin-HATTS/response"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RequestUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

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
	var requestUser RequestUser
	// var requestUser model.User
	ctx.Bind(&requestUser)
	fmt.Println(requestUser)

	name := requestUser.Name
	email := requestUser.Email
	password := requestUser.Password
	fmt.Println(name, email, password)

	if len(email) == 0 || len(name) == 0 || len(password) == 0 {
		response.Response(ctx, 422, 422, nil, "User 为空")
		return
	}

	//判断账号是否存在
	if isEmailExist(db, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密失败")
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "token发放异常")
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}
