package cache

import "github.com/go-redis/redis"

type Client interface {
	Add(ClientEntity)
	Delete(ClientEntity)
	GetClients() ([]ClientEntity, error)
}

type ClientCache struct {
	Client
}

func NewCache(redis *redis.Client) *ClientCache {
	return &ClientCache{
		Client: NewClientRedis(redis),
	}
}
