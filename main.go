package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 初始化gin实例
	g := gin.New()

	// 注册中间件
	g.Use(gin.Logger(), gin.Recovery())

	// 注册路由
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world!",
		})
	})

	// 处理 404 请求
	g.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 json
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "路由未定义",
			})
		}
	})
	// 运行
	err := g.Run(":8000")
	if err != nil {
		return
	}
}
