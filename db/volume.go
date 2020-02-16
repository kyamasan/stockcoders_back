package db

import (
	"encoding/json"
	"log"
	"stockcoder/server/db/connection"
	"stockcoder/server/db/query"
)

func GetVolumeData(cd string) (y string, e error) {
	//DB接続情報
	db, err := connection.DbAccess(connection.GetDbConnection())
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query.GetSelectVolume(), cd)
	if err != nil {
		log.Fatal(err)
	}
	var volume float32
	var yaxis []float32

	for rows.Next() {
		err := rows.Scan(&volume)
		if err != nil {
			log.Fatal(err)
		}
		yaxis = append(yaxis, volume)
	}
	jsonBytes, _ := json.Marshal(yaxis)
	return string(jsonBytes), nil
}
