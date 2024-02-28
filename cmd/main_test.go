package main

import (
	"reflect"
	"testing"
	"time"

	dec "github.com/shopspring/decimal"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"
)

func TestUpdatePrices(t *testing.T) {
	all_prices := []price.Price{
		price.Price{Symbol: "AAPL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(100, 0), High: dec.New(100, 0), Low: dec.New(100, 0), Close: dec.New(100, 0), Volume: 1111},
		price.Price{Symbol: "AAPL", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(101, 0), High: dec.New(101, 0), Low: dec.New(101, 0), Close: dec.New(101, 0), Volume: 1111},
		price.Price{Symbol: "GOOGL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(102, 0), High: dec.New(102, 0), Low: dec.New(102, 0), Close: dec.New(102, 0), Volume: 1111},
		price.Price{Symbol: "GOOGL", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(103, 0), High: dec.New(103, 0), Low: dec.New(103, 0), Close: dec.New(103, 0), Volume: 1111},
		price.Price{Symbol: "MSFT", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(104, 0), High: dec.New(104, 0), Low: dec.New(104, 0), Close: dec.New(104, 0), Volume: 1111},
	}
	prices := all_prices[0:4]

	all_tickers := []string{"AAPL", "GOOGL", "MSFT"}
	tickers := all_tickers[0:2]

	t.Run("Updates specified tickers", func(t *testing.T) {

		prices_map := map[string][]price.Price{
			"AAPL":  prices[0:2],
			"GOOGL": prices[2:4],
		}

		loader := &StubLoader{prices: prices_map}
		priceStore := &StubPriceStore{}

		UpdatePrices(loader, priceStore, tickers...)

		persisted := priceStore.GetPrices()

		if !reflect.DeepEqual(prices, persisted) {
			t.Errorf("Expected %v, got %v", prices, persisted)
		}
	})
}

type StubLoader struct {
	prices map[string][]price.Price
}

func (s *StubLoader) Load(ticker string) ([]price.Price, error) {
	return s.prices[ticker], nil
}

type StubPriceStore struct {
	prices []price.Price
}

func (s *StubPriceStore) Persist(prices []price.Price) error {
	s.prices = append(s.prices, prices...)
	return nil
}

func (s *StubPriceStore) GetPrices() []price.Price {
	return s.prices
}
