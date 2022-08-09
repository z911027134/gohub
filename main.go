package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub.com/bootstrap"
)

func main() {

	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":9501")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
