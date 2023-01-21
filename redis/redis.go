package redis

import (
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/log"
	"github.com/gomodule/redigo/redis"
	"os"
)

var (
	redisConn redis.Conn
)

func RedisInit() {
	conn, err := redis.Dial("tcp", config.Conf.Redis.Host+":"+config.Conf.Redis.Port,
		redis.DialUsername(config.Conf.Redis.User),
		redis.DialPassword(config.Conf.Redis.Pass),
	)
	if err != nil {
		log.Throwln("Redis", "Redis Server connect failed", err)
		os.Exit(1)
	}
	redisConn = conn
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
		}
	}(redisConn)
}
