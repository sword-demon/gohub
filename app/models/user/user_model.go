// Package user 存放用户 model 相关逻辑
package user

import "gohub/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	// 因为我们不希望将敏感信息输出给用户 使用 `json:"-"` 的方式 JSON 解析器忽略字段
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
