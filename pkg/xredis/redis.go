package xredis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

func NewRdsClient() (client *redis.Client,err error) {
	client = redis.NewClient(
		&redis.Options{
			Addr:         viper.GetString("redis.local.addr"),
			Password:     "",
			DB:           0,
			DialTimeout:  600 * time.Millisecond,
			ReadTimeout:  600 * time.Millisecond,
			WriteTimeout: 600 * time.Millisecond,
			MaxRetries:   64,
			PoolSize:     10,
			MinIdleConns: 5,
		})
	err = client.Ping().Err()
	return
}
