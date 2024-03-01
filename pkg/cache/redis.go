package cache

import (
	"collector-sidecar-server/pkg/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ ICache = (*defaultRedisCache)(nil)

type defaultRedisCache struct {
	client *redis.Client
}

// Set implements ICache.
func (d *defaultRedisCache) Set(key string, value interface{}) error {
	// judge value type and set to redis
	switch valueType := value.(type) {
	case string:
		return d.client.Set(context.Background(), key, valueType, 0).Err()
	case int:
		return d.client.Set(context.Background(), key, valueType, 0).Err()
	case int64:
		return d.client.Set(context.Background(), key, valueType, 0).Err()
	case float64:
		return d.client.Set(context.Background(), key, valueType, 0).Err()
	case bool:
		return d.client.Set(context.Background(), key, valueType, 0).Err()
	default:
		return d.client.Set(context.Background(), key, value, 0).Err()
	}
}

// Get implements ICache.
func (d *defaultRedisCache) Get(key string) (interface{}, error) {
	return d.client.Get(context.Background(), key).Result()
}

// Close implements ICache.
func (d *defaultRedisCache) Close() {
	if d.client != nil {
		d.client.Close()
	}
}

func NewDefaultRedisCache(c config.RedisConfig) *defaultRedisCache {

	redisClient := redis.NewClient(&redis.Options{
		DB:           c.Db,
		Addr:         c.Addr,
		Password:     c.Password,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
		IdleTimeout:  time.Duration(c.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	return &defaultRedisCache{
		client: redisClient,
	}
}
