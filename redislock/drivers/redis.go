package drivers

import (
	"fmt"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis"
)

const (
	redisHost      = "localhost"
	redisPort      = 6379
	redisNamespace = 0
	redisPassword  = ""
)

// RedisPool defines redis pooling
type RedisPool struct {
	client *redis.Client
}

// NewRedisConn will initialize redis
func NewRedisConn() (*RedisPool, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		DB:       redisNamespace,
	}
	client := redis.NewClient(options)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisPool{
		client: client,
	}, nil
}

// GetClient will get redis client
func (r *RedisPool) GetClient() *redis.Client {
	return r.client
}

// GetPoolLocker gets redis mutex locker
func (r *RedisPool) GetPoolLocker() *redislock.Client {
	return redislock.New(r.client)
}
