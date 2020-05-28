package db

import (
	"database/sql"
	"log"

	"./connection"
	"./query"
)

func AddPriceData(cd string, date string, start int, high int, low int, close int) (res sql.Result) {
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
		res, err = db.Exec(query.UpdatePriceData(), start, high, low, close, cd, date)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		res, err = db.Exec(query.AddPriceData(), cd, date, start, high, low, close)
		if err != nil {
			log.Fatal(err)
		}
	}

	return res
}
