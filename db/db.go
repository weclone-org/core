package db

import (
	"database/sql"
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/log"
	_ "github.com/lib/pq"
	"os"
	"time"
)

var (
	db *sql.DB
)

func DBPollingPing() {
	for {
		time.Sleep(time.Duration(3) * time.Second)
		err := db.Ping()
		if err != nil {
			log.Throwln("DB", "postgresql database ping connect failed", err)
			os.Exit(1)
		}
	}
}

func DBInit() {
	dsn := "host=" + config.Conf.PostgreSQL.Host + " port=" + config.Conf.PostgreSQL.Port + " user=" + config.Conf.PostgreSQL.DBUser + " password=" + config.Conf.PostgreSQL.DBPass + " dbname=" + config.Conf.PostgreSQL.DBName + " sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Throwln("DB", "postgresql database initial connect failed", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Throwln("DB", "postgresql database initial connect failed", err)
		os.Exit(1)
	}
	go DBPollingPing()
}
