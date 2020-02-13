package db

import (
	"encoding/json"
	"log"
	"stockcoder/server/db/connection"
	"stockcoder/server/db/query"
)

func GetDateData(cd string) (x string, e error) {
	db := connection.DbAccess(connection.GetDbConnection())
	defer db.Close()

	rows, err := db.Query(query.GetSelectDate(), cd)
	if err != nil {
		log.Fatal("Wrong Query")
	}
	var date string
	var xaxis []string

	for rows.Next() {
		err := rows.Scan(&date)
		if err != nil {
			log.Fatal("Scan Failure")
		}
		xaxis = append(xaxis, date)
	}
	jsonBytes, _ := json.Marshal(xaxis)
	return string(jsonBytes), nil
}