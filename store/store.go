package store

import (
	"database/sql"
	"strings"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"

	_ "github.com/mattn/go-sqlite3"
)

const PRICE_TABLE_NAME = "marketdata_historicalpricerecord"

type PriceStore interface {
	Persist(prices []price.Price) error
}

type SQLPriceStore struct {
	dbName string
}

func (s *SQLPriceStore) Persist(prices []price.Price) error {
	db, err := sql.Open("sqlite3", s.dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStr := "INSERT OR REPLACE INTO " + PRICE_TABLE_NAME + " (symbol, date, open, high, low, close, volume) VALUES "
	for i := 0; i < len(prices); i++ {
		sqlStr += "(?, ?, ?, ?, ?, ?, ?),"
	}
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}

	for _, p := range prices {
		date := p.Date.Format("2006-01-02")
		values = append(values, p.Symbol, date, p.Open, p.High, p.Low, p.Close, p.Volume)
	}

	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func NewSQLPriceStore(dbName string) PriceStore {
	return &SQLPriceStore{dbName: dbName}
}
