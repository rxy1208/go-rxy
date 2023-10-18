package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rxy/bootstrap"
)

func main() {
	// 初始化gin实例
	g := gin.New()

	// 初始化路由和全局中间件
	bootstrap.SetupRoute(g)

	// 运行
	err := g.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
