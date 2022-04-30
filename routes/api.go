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
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送短信
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 发送邮件
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}
	}
}
