package executor

import (
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/db"
	"github.com/Weclone-org/core/log"
	"github.com/Weclone-org/core/redis"
)

func Run() {
	config.ParseInit()
	log.Println("Configuration load successfully")
	log.SetStrLevel(config.Conf.Level)
	log.Println("Set LogLevel: %s", config.Conf.Level)
	log.Println("connect database and redis...")
	db.DBInit()
	log.Println("connect postgresql database successfully")
	redis.RedisInit()
	log.Println("connect redis server successfully")
}
