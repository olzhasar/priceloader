package store

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"
	dec "github.com/shopspring/decimal"

	_ "github.com/mattn/go-sqlite3"
)

const TEST_DB = "test_price.db"

func TestSQLPriceStorePersist(t *testing.T) {
	t.Run("persists prices to the database", func(t *testing.T) {
		dropTestDatabase(TEST_DB)

		store, err := NewSQLPriceStore(TEST_DB)
		if err != nil {
			t.Fatalf("Error creating SQLPriceStore: %s", err)
		}

		defer store.Close()

		prices := []price.Price{
			{Symbol: "AAPL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(100, 0), High: dec.New(101, 0), Low: dec.New(102, 0), Close: dec.New(103, 0), Volume: 1111},
			{Symbol: "MSFT", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(200, 0), High: dec.New(201, 0), Low: dec.New(202, 0), Close: dec.New(203, 0), Volume: 1111},
		}

		err = store.Persist(prices)

		if err != nil {
			t.Fatalf("Error persisting prices: %s", err)
		}

		persistedPrices := loadPersistedPrices()

		assertPricesEqual(t, prices, persistedPrices)
	})

	t.Run("overwrites existing records", func(t *testing.T) {
		dropTestDatabase(TEST_DB)

		store, err := NewSQLPriceStore(TEST_DB)
		if err != nil {
			t.Fatalf("Error creating SQLPriceStore: %s", err)
		}
		defer store.Close()

		old_prices := []price.Price{
			{Symbol: "AAPL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(100, 0), High: dec.New(101, 0), Low: dec.New(102, 0), Close: dec.New(103, 0), Volume: 1111},
		}

		store.Persist(old_prices)

		new_prices := []price.Price{
			{Symbol: "AAPL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(999, 0), High: dec.New(999, 0), Low: dec.New(999, 0), Close: dec.New(103, 0), Volume: 1111},
			{Symbol: "MSFT", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(300, 0), High: dec.New(301, 0), Low: dec.New(302, 0), Close: dec.New(203, 0), Volume: 1111},
		}

		err = store.Persist(new_prices)
		if err != nil {
			t.Fatalf("Error persisting new prices: %s", err)
		}

		persistedPrices := loadPersistedPrices()

		assertPricesEqual(t, new_prices, persistedPrices)
	})
}

func loadPersistedPrices() []price.Price {
	db, err := sql.Open("sqlite3", TEST_DB)
	if err != nil {
		panic(err)
	}

	query := "SELECT symbol, date, open, high, low, close, volume FROM " + PRICE_TABLE_NAME

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var result []price.Price
	for rows.Next() {
		var p price.Price
		err := rows.Scan(&p.Symbol, &p.Date, &p.Open, &p.High, &p.Low, &p.Close, &p.Volume)
		if err != nil {
			panic(err)
		}
		result = append(result, p)
	}

	return result
}

func assertPricesEqual(t testing.TB, expected, actual []price.Price) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func dropTestDatabase(dbName string) {
	os.Remove(dbName)
}
