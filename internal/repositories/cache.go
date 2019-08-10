package repositories

import (
	"encoding/json"

	"github.com/cikupin/redis-mutex-lock/drivers"
	"github.com/cikupin/redis-mutex-lock/internal/models"
	"github.com/gomodule/redigo/redis"
)

const (
	userCacheKey = "dummy_user"
)

// ICache defines interface for cache
type ICache interface {
	GetCache(key string) (models.User, error)
	UpdateCache(key string) error
}

// CacheRepo defines cache repost=sitory struct
type CacheRepo struct {
	redisDriver *drivers.RedisPool
	user        *models.User
}

// NewCacheRepo will initialize new instance of cache repo
func NewCacheRepo(redis *drivers.RedisPool) *CacheRepo {
	return &CacheRepo{
		redisDriver: redis,
		user:        models.NewUser(),
	}
}

// GetCache will get cache value by key
func (c *CacheRepo) GetCache(key string) (models.User, error) {
	var (
		user models.User
		conn redis.Conn
	)

	conn = c.redisDriver.GetConn()
	defer conn.Close()

	userCache, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(userCache), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateCache will update cache value
func (c *CacheRepo) UpdateCache(key string) error {
	return nil
}
