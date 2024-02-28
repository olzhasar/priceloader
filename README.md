# priceloader

The CLI utility to create a local SQLite database of daily stock prices (OHLC). Can be used as a cron job to keep the database up to date.

Data is being fetched from Yahoo Finance using the [finance-go](github.com/piquette/finance-go) library.
