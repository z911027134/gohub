package routes

import (
	"github.com/gin-gonic/gin"
	"gohub.com/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
		authGroup := v1.Group("auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
