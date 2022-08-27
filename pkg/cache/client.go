package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

type ClientRedis struct {
	Redis *redis.Client
}

type ClientEntity struct {
	Value string
}

func NewClientRedis(redis *redis.Client) *ClientRedis {
	return &ClientRedis{Redis: redis}
}

func (c *ClientRedis) Add(entity ClientEntity) {
	c.Redis.RPush(RedisCashNameClients, entity.Value)
}

func (c *ClientRedis) Delete(entity ClientEntity) {
	c.Redis.LRem(RedisCashNameClients, 0, entity.Value)
}

func (c *ClientRedis) GetClients() ([]ClientEntity, error) {
	result, err := c.Redis.LRange(RedisCashNameClients, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get client data from redis: %v", err)
	}
	var clients []ClientEntity
	for _, client := range result {
		clients = append(clients, ClientEntity{client})
	}
	return clients, err
}
