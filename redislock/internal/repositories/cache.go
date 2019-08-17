package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bsm/redislock"
	"github.com/cikupin/redis-mutex-lock/redislock/drivers"
	"github.com/cikupin/redis-mutex-lock/redislock/internal/constants"
	"github.com/cikupin/redis-mutex-lock/redislock/internal/models"
	"github.com/go-redis/redis"
)

// ICache defines interface for cache
type ICache interface {
	GetCache(key string) (models.User, error)
	UpdateCache(key string) error
	GetCacheWithThunderingHerd(key string) (models.User, error)
}

// CacheRepo defines cache repost=sitory struct
type CacheRepo struct {
	redisDriver *drivers.RedisPool
	redisClient *redis.Client
	user        *models.User
	// mx          sync.Mutex
}

// NewCacheRepo will initialize new instance of cache repo
func NewCacheRepo(redis *drivers.RedisPool) *CacheRepo {
	return &CacheRepo{
		redisDriver: redis,
		redisClient: redis.GetClient(),
		user:        models.NewUser(),
	}
}

// GetCache will get cache value by key
func (c *CacheRepo) GetCache(key string) (models.User, error) {
	var user models.User

	userCache, err := c.redisClient.Get(key).Result()
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
	jsonStr, err := json.Marshal(c.user)
	if err != nil {
		return err
	}

	err = c.redisClient.Set(key, jsonStr, constants.UserCacheTTL*time.Second).Err()
	return err
}

// GetCacheWithThunderingHerd will get data from cache
// If data not exist, update with thundering herd
func (c *CacheRepo) GetCacheWithThunderingHerd(key string) (models.User, error) {
	var (
		user   models.User
		locker *redislock.Lock
		err    error
	)

	lockKey := fmt.Sprintf("lock-key-%s", key)
	opts := &redislock.Options{
		RetryCount: constants.RedisLockerTries,
	}

	for {
		locker, err = c.redisDriver.GetPoolLocker().Obtain(lockKey, constants.RedisLockerExpiry, opts)
		if err == nil {
			break
		}
	}
	defer locker.Release()

	userCache, err := c.redisClient.Get(key).Result()
	if err != nil && err != redis.Nil {
		return user, err
	}

	if err == redis.Nil {
		log.Println("<<<<<<<<<< [ THUNDERING HERD ] UPDATING DATA >>>>>>>>>>>>")
		time.Sleep(constants.RedisLockerExpiry)
		_ = c.UpdateCache(key)
		log.Println("<<<<<<<<<< [ THUNDERING HERD ] FINISHED UPDATING DATA >>>>>>>>>>>>")

		userCache, _ = c.redisClient.Get(key).Result()
	}

	err = json.Unmarshal([]byte(userCache), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
