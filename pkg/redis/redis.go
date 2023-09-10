package cachedRedis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/kaustubh-pandey-kp/storage-utility/constants"
	"github.com/kaustubh-pandey-kp/storage-utility/pkg/logger"
)

var (
	Pool *redis.Pool
)

/*
*
Init redis connection
*/
func RedisConnection(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     constants.REDIS_CONNECTION_POOL,
		IdleTimeout: 1000 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				//Logger.Panicf("redis-connection-failed - Unable to connecto redis server %s,  host: %s", err.Error(), constants.REDIS_HOST)
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

/*
*
Initialising redis connection
*/
func InitRedis() {

	// initialise pool for test commands
	redisHost := constants.CACHED_REDIS_HOST
	Pool = RedisConnection(redisHost)
	err := Pool.TestOnBorrow(Pool.Get(), time.Time{})

	if err != nil {
		logger.Errorf("Unable to connect to redis server %s,  host: %s", err.Error(), constants.CACHED_REDIS_HOST)
	} else {
		logger.Infof("Connection to redis host: %s established successfully", constants.CACHED_REDIS_HOST)
	}

}
