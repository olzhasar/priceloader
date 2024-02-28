package price

import (
	"time"

	"github.com/shopspring/decimal"
)

type Price struct {
	Symbol string          `json:"symbol"`
	Date   time.Time       `json:"date"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume int             `json:"volume"`
}

func New(symbol string, date time.Time, open, high, low, close decimal.Decimal, volume int) Price {
	return Price{
		Symbol: symbol, Date: date,
		Open: open, High: high, Low: low, Close: close,
		Volume: volume,
	}
}
