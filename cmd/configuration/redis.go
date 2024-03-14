package configuration

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func (cf ConfigurationSchema) NewRedis() *redis.Client {
	defaultTimeout := 30 * time.Second

	client := redis.NewClient(&redis.Options{
		Addr:         cf.Redis.Address,
		Password:     cf.Redis.Password,
		DB:           cf.Redis.DB,
		DialTimeout:  defaultTimeout,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
		MaxRetries:   cf.Redis.MaxRetry,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Panic(err)
	}

	return client
}
