package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/database"
	"strings"
)

func init() {
	// 自定义规则 not_exits
	govalidator.AddCustomRule("no_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数 表名：如: users
		tableName := rng[0]
		// 第二个参数 字段名：如: name
		dbField := rng[1]

		// 第三个参数 排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接SQL
		query := database.DB.Table(tableName).Where(dbField+" = ?", requestValue)

		// 如果传第三个参数，加上 SQL where 过滤
		if len(exceptID) > 0 {
			query = query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过 数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})
}
