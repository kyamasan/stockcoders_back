package db

import (
	"database/sql"
	"log"

	"./connection"
	"./query"
)

func RemovePriceData(cd string, date string) (res sql.Result) {
	//DB接続情報
	db, err := connection.DbAccess()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var count int

	err = db.QueryRow(query.CheckPriceData(), cd, date).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count != 0 {
		res, err = db.Exec(query.RemovePriceData(), cd, date)
		if err != nil {
			log.Fatal(err)
		}
	}

	return res
}
