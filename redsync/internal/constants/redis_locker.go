package constants

import (
	"time"
)

const (
	// RedisLockerTries set redis locker tries
	RedisLockerTries = 5
	// RedisLockerExpiry set redis locker expiration time
	RedisLockerExpiry = 10 * time.Second
)
