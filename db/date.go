package db

import (
	"encoding/json"
	"log"

	"./connection"
	"./query"
)

func GetDateData(cd string) (x string, e error) {
	//DB接続情報
	db, err := connection.DbAccess()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query.GetSelectDate(), cd)
	if err != nil {
		log.Fatal(err)
	}
	var date string
	var xaxis []string

	for rows.Next() {
		err := rows.Scan(&date)
		if err != nil {
			log.Fatal(err)
		}
		xaxis = append(xaxis, date)
	}
	jsonBytes, _ := json.Marshal(xaxis)
	return string(jsonBytes), nil
}
