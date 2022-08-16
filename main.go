package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub.com/bootstrap"
	baseConfig "gohub.com/config"
	"gohub.com/pkg/config"
)

func init() {
	baseConfig.Initialize()
}

func main() {

	// 读取命令行的当前环境变量
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=test 加载的是 .env.test 文件")
	flag.Parse()
	config.InitConfig(env)
	r := gin.New()
	bootstrap.SetupDB()
	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
