package update

import (
	"log"

	"github.com/olzhasar/portfolio-tracker-priceloader/loader"
	"github.com/olzhasar/portfolio-tracker-priceloader/store"
)

func UpdateSingle(loader loader.Loader, priceStore store.PriceStore, ticker string) error {
	log.Printf("Loading prices for %s", ticker)
	prices, err := loader.Load(ticker)
	if err != nil {
		log.Printf("Error loading prices for %s:\n%s", ticker, err)
		return err
	}

	if prices == nil {
		log.Printf("ERROR: No prices found for %s", ticker)
		return err
	}

	err = priceStore.Persist(prices)
	if err != nil {
		log.Printf("Error persisting prices for %s:\n%s", ticker, err)
		return err
	}

	return nil
}

func UpdateMultiple(loader loader.Loader, priceStore store.PriceStore, tickers ...string) error {
	for _, ticker := range tickers {
		UpdateSingle(loader, priceStore, ticker)
	}

	return nil
}
