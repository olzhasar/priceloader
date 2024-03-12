package store

import (
	"database/sql"
	"strings"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"

	_ "github.com/mattn/go-sqlite3"
)

const PRICE_TABLE_NAME = "prices"

type PriceStore interface {
	Persist(prices []price.Price) error
	Close() error
}

type SQLPriceStore struct {
	db *sql.DB
}

func (s *SQLPriceStore) Persist(prices []price.Price) error {
	tx, err := s.db.Begin()
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

func (s *SQLPriceStore) createTable() error {
	sqlStmt := `CREATE TABLE IF NOT EXISTS ` + PRICE_TABLE_NAME + ` (
	id integer not null primary key,
	symbol text	 not null,
	date text not null,
	open numeric not null,
	high numeric not null,
	low numeric not null,
	close numeric not null,
	volume integer not null,
	CONSTRAINT "uc_symbol_date" UNIQUE ("symbol", "date")
	);`

	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLPriceStore) Close() error {
	return s.db.Close()
}

func NewSQLPriceStore(dbName string) (PriceStore, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	store := &SQLPriceStore{db: db}

	err = store.createTable()
	if err != nil {
		return nil, err
	}

	return store, nil
}
