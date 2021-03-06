package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cikupin/redis-mutex-lock/redsync/drivers"
	"github.com/cikupin/redis-mutex-lock/redsync/internal/constants"
	"github.com/cikupin/redis-mutex-lock/redsync/internal/models"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
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
	conn := c.redisDriver.GetConn()
	defer conn.Close()

	jsonStr, err := json.Marshal(c.user)
	if err != nil {
		return err
	}

	_, err = redis.String(conn.Do("SETEX", key, constants.UserCacheTTL, jsonStr))
	if err != nil {
		return err
	}
	return nil
}

// GetCacheWithThunderingHerd will get data from cache
// If data not exist, update with thundering herd
func (c *CacheRepo) GetCacheWithThunderingHerd(key string) (models.User, error) {
	var (
		user models.User
		conn redis.Conn
	)

	conn = c.redisDriver.GetConn()
	defer conn.Close()

	lockKey := fmt.Sprintf("lock-key-%s", key)
	mtx := c.redisDriver.GetPoolLocker().NewMutex(
		lockKey,
		redsync.SetTries(constants.RedisLockerTries),
		redsync.SetExpiry(constants.RedisLockerExpiry),
	)

	userCache, err := redis.String(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return user, err
	}

	if err == redis.ErrNil {
		for {
			err := mtx.Lock()
			if err == nil {
				break
			}
		}
		defer mtx.Unlock()

		userCache, err = redis.String(conn.Do("GET", key))
		if err != nil && err != redis.ErrNil {
			return user, err
		}

		if err == redis.ErrNil {
			log.Println("<<<<<<<<<< [ THUNDERING HERD ] UPDATING DATA >>>>>>>>>>>>")
			time.Sleep(constants.RedisLockerExpiry)
			_ = c.UpdateCache(key)
			log.Println("<<<<<<<<<< [ THUNDERING HERD ] FINISHED UPDATING DATA >>>>>>>>>>>>")

			userCache, _ = redis.String(conn.Do("GET", key))
		}
	}

	err = json.Unmarshal([]byte(userCache), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
