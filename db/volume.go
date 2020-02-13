package db

import (
	"encoding/json"
	"log"
	"stockcoder/server/db/connection"
	"stockcoder/server/db/query"
)

func GetVolumeData(cd string) (y string, e error) {
	db := connection.DbAccess(connection.GetDbConnection())
	defer db.Close()

	rows, err := db.Query(query.GetSelectVolume(), cd)
	if err != nil {
		log.Fatal("Wrong Query")
	}
	var volume float32
	var yaxis []float32

	for rows.Next() {
		err := rows.Scan(&volume)
		if err != nil {
			log.Fatal("Scan Failure")
		}
		yaxis = append(yaxis, volume)
	}
	jsonBytes, _ := json.Marshal(yaxis)
	return string(jsonBytes), nil
}