# priceloader

This is a simple CLI utility to maintain a local SQLite database of daily stock prices (OHLC). Can be used as a cron job to keep the database up to date.

Data is being fetched from Yahoo Finance using the [finance-go](github.com/piquette/finance-go) library.

## Requirements

- Go 1.22+

## Installation

Build the binary using the provided Makefile:

```bash
make build
```

## Usage

Load the daily prices for AAPL and MSFT into the database `main.db`:

```bash
./priceloader -db main.db AAPL MSFT
```

## Database schema

The application creates a table `prices` with the following schema:

```sql
CREATE TABLE prices (
	id integer not null primary key,
	symbol text not null,
	date text not null,
	open numeric not null,
	high numeric not null,
	low numeric not null,
	close numeric not null,
	volume integer not null,
	CONSTRAINT "uc_symbol_date" UNIQUE ("symbol", "date")
);
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## TODO

- [ ] Load prices concurrently
