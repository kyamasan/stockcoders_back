package query

func GetSelectData() string {
	return `SELECT  date,
	CAST(open AS Decimal(20,2)),
	CAST(high AS Decimal(20,2)),
	CAST(low AS Decimal(20,2)),
	CAST(close AS Decimal(20,2))
	FROM stockdata
	WHERE CD = ?
	ORDER BY date`
}

func GetSelectVolume() string {
	return `SELECT Volume AS volume
	FROM StockData
	WHERE CD = ? 
	ORDER BY date`
}

func GetSelectDate() string {
	return `SELECT Date AS date
	FROM StockData
	WHERE CD = ? 
	ORDER BY date`
}

func CheckPriceData() string {
	return `SELECT COUNT(*) FROM stockdata 
	WHERE cd = ? and date = ?
	LIMIT 1;`
}

func AddPriceData() string {
	return `INSERT INTO stockdata 
	(cd, date, open, high, low, close)
	VALUES (?, ?, ?, ?, ?, ?);`
}
func UpdatePriceData() string {
	return `UPDATE stockdata 
	SET open = ?,  high = ?, low = ?, close = ?
	WHERE cd = ? and date = ?;`
}
func RemovePriceData() string {
	return `DELETE FROM stockdata 
	WHERE cd = ? and date = ?;`
}
