package connection

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//GCPのMySQL接続に必要な情報
type ConnectionInfo struct {
	DbUser     string
	DbPassword string
	DbEnv      string
	DbName     string
	DbPort     string
}

func getDbConnection() ConnectionInfo {
	return ConnectionInfo{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbEnv:      os.Getenv("DB_ENV"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}

func DbAccess() (db *sql.DB, e error) {
	con := getDbConnection()
	db, err := sql.Open("mysql", con.DbUser+":"+con.DbPassword+"@tcp("+con.DbEnv+":"+con.DbPort+")/"+con.DbName)
	return db, err
}
