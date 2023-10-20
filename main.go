package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rxy/bootstrap"
	btsConfig "go-rxy/config"
	"go-rxy/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化gin实例
	g := gin.New()

	// 初始化路由和全局中间件
	bootstrap.SetupRoute(g)

	// 运行
	err := g.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
