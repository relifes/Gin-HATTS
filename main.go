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
// 协议头 - 加密协议 /
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcsImV4cCI6MTYyMTA2MDcyMSwiaWF0IjoxNjIwNDU1OTIxLCJpc3MiOiJIQVRUUyIsInN1YiI6InVzZXIgdG9rZW4ifQ.fCexYsIAtXM3_VKIsYDQDZBgdtmwhYfKwq_aXLlTnUo
