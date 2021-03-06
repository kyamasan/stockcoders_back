package db

import (
	"encoding/json"
	"log"

	"./query"

	"./connection"
)

type StockData struct {
	Cd     string
	Prices []Price
}

type Price struct {
	Date        string     `json:"x"`
	PriceRecord [4]float32 `json:"y"`
}

func GetData(cd string) (data string, e error) {
	db, err := connection.DbAccess()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query.GetSelectData(), cd)
	if err != nil {
		log.Fatal(err)
	}

	var prices []Price
	var date string
	var openingPrice float32
	var highPrice float32
	var lowPrice float32
	var closingPrice float32

	for rows.Next() {
		err := rows.Scan(&date, &openingPrice, &highPrice, &lowPrice, &closingPrice)
		if err != nil {
			log.Fatal(err)
		}

		price := Price{
			Date:        date,
			PriceRecord: [4]float32{openingPrice, highPrice, lowPrice, closingPrice},
		}
		prices = append(prices, price)
	}
	jsonBytes, _ := json.Marshal(prices)
	return string(jsonBytes), nil
}
