package update

import (
	"reflect"
	"testing"
	"time"

	dec "github.com/shopspring/decimal"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"
)

func TestUpdateMultiple(t *testing.T) {
	all_prices := []price.Price{
		price.Price{Symbol: "AAPL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(100, 0), High: dec.New(100, 0), Low: dec.New(100, 0), Close: dec.New(100, 0), Volume: 1111},
		price.Price{Symbol: "AAPL", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(101, 0), High: dec.New(101, 0), Low: dec.New(101, 0), Close: dec.New(101, 0), Volume: 1111},
		price.Price{Symbol: "GOOGL", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Open: dec.New(102, 0), High: dec.New(102, 0), Low: dec.New(102, 0), Close: dec.New(102, 0), Volume: 1111},
		price.Price{Symbol: "GOOGL", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(103, 0), High: dec.New(103, 0), Low: dec.New(103, 0), Close: dec.New(103, 0), Volume: 1111},
		price.Price{Symbol: "MSFT", Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Open: dec.New(104, 0), High: dec.New(104, 0), Low: dec.New(104, 0), Close: dec.New(104, 0), Volume: 1111},
	}
	prices_map := map[string][]price.Price{
		"AAPL":  all_prices[0:2],
		"GOOGL": all_prices[2:4],
		"MSFT":  all_prices[4:5],
	}

	tickers := []string{"AAPL", "GOOGL"}

	loader := &StubLoader{prices: prices_map}
	priceStore := &StubPriceStore{}

	expected := all_prices[0:4]

	UpdateMultiple(loader, priceStore, tickers...)

	persisted := priceStore.GetPrices()

	if !reflect.DeepEqual(expected, persisted) {
		t.Errorf("Expected %v, got %v", expected, persisted)
	}
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
