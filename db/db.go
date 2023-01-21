package db

import (
	"database/sql"
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/log"
	_ "github.com/lib/pq"
	"os"
)

var (
	db *sql.DB
)

func DBInit() {
	dsn := "host=" + config.Conf.PostgreSQL.Host + " port=" + config.Conf.PostgreSQL.Port + " user=" + config.Conf.PostgreSQL.DBUser + " password=" + config.Conf.PostgreSQL.DBPass + " dbname=" + config.Conf.PostgreSQL.DBName + " sslmode=disable"
	dbt, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Throwln("DB", "postgresql database connect failed", err)
		os.Exit(1)
	}
	err = dbt.Ping()
	if err != nil {
		log.Throwln("DB", "postgresql database connect failed", err)
		os.Exit(1)
	}
	db = dbt
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
}
