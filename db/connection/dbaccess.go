package connection

import (
	"database/sql"
	"os"
	"log"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

type ConnectionInfo struct {
	DbEnv      string
	DbUser     string
	DbPassword string
	DbName     string
}

func GetDbConnection() ConnectionInfo {
	return ConnectionInfo{
		DbEnv:      os.Getenv("DB_ENV"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

func DbAccess(con ConnectionInfo) (db *sql.DB) {
	cfg := mysql.Cfg(con.DbEnv, con.DbUser, con.DbPassword)
	cfg.DBName = con.DbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		log.Fatal("Error connecting server")
	}
	return db
}