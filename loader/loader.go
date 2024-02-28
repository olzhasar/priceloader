package loader

import (
	"log"
	"time"

	"github.com/olzhasar/portfolio-tracker-priceloader/price"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

type Loader interface {
	Load(symbol string) ([]price.Price, error)
}

type YahooLoader struct{}

func NewYahooLoader() *YahooLoader {
	return &YahooLoader{}
}

func (l *YahooLoader) Load(symbol string) ([]price.Price, error) {
	today := time.Now()
	start := datetime.Datetime{Month: int(today.Month()), Day: today.Day(), Year: today.Year() - 1}
	end := datetime.Datetime{Month: int(today.Month()), Day: today.Day(), Year: today.Year()}

	params := &chart.Params{
		Symbol:   symbol,
		Start:    &start,
		End:      &end,
		Interval: datetime.OneDay,
	}

	var result []price.Price

	i := 0

	iter := chart.Get(params)
	for iter.Next() {
		var p price.Price

		bar := iter.Bar()

		p.Symbol = symbol
		p.Date = time.Unix(int64(bar.Timestamp), 0)
		p.Open = bar.Open
		p.High = bar.High
		p.Low = bar.Low
		p.Close = bar.Close
		p.Volume = bar.Volume

		result = append(result, p)
		i++
	}

	log.Printf("Loaded %d records for %s", i, symbol)

	return result, nil
}
