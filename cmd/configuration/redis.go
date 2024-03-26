package configuration

import (
	"time"

	"github.com/go-redis/redis"
)

func (cf ConfigurationSchema) NewRedis() (config *redis.Options) {
	defaultTimeout := 30 * time.Second
	config = &redis.Options{
		Addr:         cf.Redis.Address,
		Password:     cf.Redis.Password,
		DB:           cf.Redis.DB,
		DialTimeout:  defaultTimeout,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
		MaxRetries:   cf.Redis.MaxRetry,
	}
	return
}
