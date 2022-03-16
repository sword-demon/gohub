// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			authGroup := v1.Group("/auth")
			{
				suc := new(auth.SignupController)
				// 判断手机是否已注册
				authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
				// 判断 Email 是否已注册
				authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			}
		})
	}
}
