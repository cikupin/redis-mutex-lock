package drivers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

const (
	redisHost      = "localhost"
	redisPort      = 6379
	redisNamespace = 0
	redisPassword  = ""
)

// RedisOption defines redis options struct
type RedisOption struct {
	Host      string
	Port      int
	Password  string
	Namespace string
}

// RedisPool defines redis pooling
type RedisPool struct {
	pool *redis.Pool
}

// NewRedisConn will initialize redis
func NewRedisConn() (*RedisPool, error) {
	redisPassword := redis.DialPassword(redisPassword)

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisHost, redisPort), redisPassword)
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("SELECT", redisNamespace); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	_, err := pool.Get().Do("PING")
	if err != nil {
		return nil, err
	}

	return &RedisPool{
		pool: pool,
	}, nil
}

// GetPool will get redis pool
func (r *RedisPool) GetPool() *redis.Pool {
	if r.pool == nil {
		log.Fatalln(errors.New("error get redis pool"))
	}
	return r.pool
}

// GetPoolLocker gets redis mutex locker
func (r *RedisPool) GetPoolLocker() *redsync.Redsync {
	var arrPool []redsync.Pool
	arrPool = append(arrPool, r.GetPool())

	return redsync.New(arrPool)
}
