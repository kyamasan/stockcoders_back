package query

func GetSelectVolume() string {
	return `SELECT Volume AS volume
	FROM StockData
	WHERE CD = ? 
	ORDER BY date`
}

func GetSelectData() string {
	return `SELECT  Date,
	CAST(OpeningPrice AS Decimal(20,2)),
	CAST(HighPrice AS Decimal(20,2)),
	CAST(LowPrice AS Decimal(20,2)),
	CAST(ClosingPrice AS Decimal(20,2))
	FROM StockData
	WHERE CD = ?
	ORDER BY Date`
}

func GetSelectDate() string {
	return `SELECT Date AS date
	FROM StockData
	WHERE CD = ? 
	ORDER BY date`
}