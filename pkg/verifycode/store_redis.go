package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (r *RedisStore) Set(id string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便测试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return r.RedisClient.Set(r.KeyPrefix+id, value, ExpireTime)
}

func (r *RedisStore) Get(id string, clear bool) string {
	key := r.KeyPrefix + id
	val := r.RedisClient.Get(key)
	if clear {
		r.RedisClient.Del(key)
	}
	return val
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
