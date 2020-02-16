package db

import (
	"stockcoder/server/db/connection"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetData(t *testing.T) {
	godotenv.Load("../.env")
	db, err := connection.DbAccess(connection.GetDbConnection())
	defer db.Close()
	if err != nil {
		//DB接続に失敗
		t.Fatalf("Error connecting server, got %q", err)
	}

	dummy, err := db.Query(getDummyData())
	if dummy == nil {
		//DB接続には成功したが、クエリ実行に失敗
		t.Fatalf("Error Query, got %q", err)
	}
}

func getDummyData() string {
	return `SELECT 1`
}
