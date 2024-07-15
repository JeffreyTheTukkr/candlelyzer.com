package providers

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
)

// Binance struct to contain implemented functions
type Binance struct {
	Client *binance.Client
}

// NewBinanceRepo struct to create a new binance repository
func NewBinanceRepo(apiKey string, secretKey string) *Binance {
	client := binance.NewClient(apiKey, secretKey)

	return &Binance{
		Client: client,
	}

}

// ListAllPairs return a list of all binance pairs
func (b *Binance) ListAllPairs() ([]binance.Symbol, error) {
	exchange, err := b.Client.NewExchangeInfoService().Do(context.Background())

	return exchange.Symbols, err
}

// FetchCandleData return candle data for a single pair
func (b *Binance) FetchCandleData(pair string, since time.Time) ([]*binance.Kline, error) {
	// override since time for dates before 2010
	if since.UnixMilli() < time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC).UnixMilli() {
		since = time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	return b.Client.NewKlinesService().Symbol(pair).StartTime(since.UnixMilli()).Limit(1000).Interval("1m").Do(context.Background())
}
