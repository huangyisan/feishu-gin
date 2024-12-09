package main

import (
	_ "feishu-gin/feishu"
	"feishu-gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 中间件
	//r.Use()
	// 注册路由信息
	routers.RegisterRouters(r)
	r.Run("127.0.0.1:8087")
}
