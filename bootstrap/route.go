package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub.com/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	RegisterGlobalMiddleware(r)
	// 注册路由
	routes.RegisterApiRoutes(r)
	// 配置404路由
	Setup404Route(r)
}

func Setup404Route(r *gin.Engine) {
	// 404 router
	r.NoRoute(func(context *gin.Context) {
		acceptStr := context.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			context.String(http.StatusNotFound, "page not found.")
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "route not defined,please check it.",
			})
		}
	})
}

func RegisterGlobalMiddleware(r *gin.Engine) {
	r.Use(gin.Logger(), gin.Recovery())
}
