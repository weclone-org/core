package executor

import (
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/db"
	"github.com/Weclone-org/core/log"
	"github.com/Weclone-org/core/node/cdn"
	"github.com/Weclone-org/core/node/root"
	"github.com/Weclone-org/core/node/server"
	"github.com/Weclone-org/core/redis"
	"github.com/Weclone-org/core/restfulapi"
)

func Run() {
	config.ParseInit()
	log.Println("Configuration load successfully")
	log.SetStrLevel(config.Conf.Level)
	log.Println("Set LogLevel: %s", config.Conf.Level)
	log.Println("connect database and redis server...")
	db.DBInit()
	log.Println("connect postgresql database successfully")
	redis.RedisInit()
	log.Println("connect redis server successfully")
	log.Println("Start " + config.Conf.Mode + " mode...")
	switch config.Conf.Mode {
	case "root":
		root.RootInit()
	case "server":
		server.ServerInit()
	case "cdn":
		cdn.CDNInit()
	}
	log.Println("Start RESTful API listen on " + config.Conf.Listen)
	restfulapi.RESTfulAPIInit()
	log.Println("Start core successfully")
}
