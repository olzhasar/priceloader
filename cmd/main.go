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

	dbAbsPath, err := filepath.Abs(dbPath)
	if err != nil {
		fmt.Println("Invalid path: ", err)
		os.Exit(1)
	}

	tickers := flag.Args()
	if len(tickers) == 0 {
		fmt.Println("At least one ticker symbol is required")
		flag.Usage()
		os.Exit(1)
	}

	priceStore, err := store.NewSQLPriceStore(dbAbsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer priceStore.Close()

	loader := loader.NewYahooLoader()

	err = update.UpdateMultiple(loader, priceStore, tickers...)

	if err != nil {
		fmt.Println("Error updating prices: ", err)
		os.Exit(1)
	}
}
