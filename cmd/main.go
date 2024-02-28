package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/olzhasar/portfolio-tracker-priceloader/loader"
	"github.com/olzhasar/portfolio-tracker-priceloader/store"
	"github.com/olzhasar/portfolio-tracker-priceloader/update"
)

func main() {
	var dbPath string

	flag.StringVar(&dbPath, "db", "", "Path to the database file")
	flag.Parse()

	if dbPath == "" {
		fmt.Println("db path is required")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("db file does not exist")
		os.Exit(1)
	}

	dbAbsPath, err := filepath.Abs(dbPath)
	if err != nil {
		fmt.Println("Invalid path: ", err)
		os.Exit(1)
	}

	priceStore := store.NewSQLPriceStore(dbAbsPath)
	loader := loader.NewYahooLoader()

	err = update.UpdateMultiple(loader, priceStore, "AAPL", "MSFT", "GOOG")

	if err != nil {
		fmt.Println("Error updating prices: ", err)
		os.Exit(1)
	}
}
