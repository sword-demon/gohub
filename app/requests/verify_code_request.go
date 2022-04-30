package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaId     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单 返回长度等于0表示验证成功
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号长度必须为11位",
		},
		"captcha_id":     []string{"required:验证码不能为空"},
		"captcha_answer": []string{"required:验证码不能为空", "digits:验证码长度必须为6位"},
	}

	errs := validate(data, rules, messages)

	_data := data.(VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaId, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码错误")
	}

	return errs
}
