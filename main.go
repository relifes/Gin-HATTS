package main

import (
	"Gin-HATTS/common"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":8000"))
}


