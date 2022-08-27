package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const (
	RedisCashNameClients = "clients"
	RedisExpireTime      = 60 * time.Second
	RedisConnectAddress  = "127.0.0.1:6379"
)

func NewRedis() (*redis.Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     RedisConnectAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rc.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis client is not responsing: %v", err)
	}

	_, err = rc.Expire(RedisCashNameClients, RedisExpireTime).Result()
	if err != nil {
		return nil, fmt.Errorf("redis client can't expire data: %v", err)
	}

	return rc, nil
}
