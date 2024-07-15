package providers

import (
	"context"
	"time"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
	"github.com/adshao/go-binance/v2"
)

// BinanceRepo required provider dependencies
type BinanceRepo struct {
	Client *binance.Client
}

// NewBinanceRepo constructor for new binance repository
func NewBinanceRepo(apiKey string, secretKey string) *BinanceRepo {
	client := binance.NewClient(apiKey, secretKey)

	return &BinanceRepo{
		Client: client,
	}
}

// ListAllPairs return a list of all binance pairs
func (br *BinanceRepo) ListAllPairs() ([]models.PairBase, error) {
	exchange, err := br.Client.NewExchangeInfoService().Do(context.Background())

	// map binance data to standardized pair
	var pairs []models.PairBase
	for _, symbol := range exchange.Symbols {
		return append(pairs, models.PairBase{
			Base:     symbol.BaseAsset,
			Quote:    symbol.QuoteAsset,
			Exchange: models.Binance,
			Status:   matchBinancePairStatus(symbol.Status),
		}), nil
	}

	return pairs, err
}

// FetchCandleData return candle data for a single pair
func (br *BinanceRepo) FetchCandleData(pair string, since time.Time) ([]*binance.Kline, error) {
	// override since time for dates before 2010
	if since.UnixMilli() < time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC).UnixMilli() {
		since = time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	return br.Client.NewKlinesService().Symbol(pair).StartTime(since.UnixMilli()).Limit(1000).Interval("1m").Do(context.Background())
}

// matchBinancePairStatus helper to match the pair status from binance to standard
func matchBinancePairStatus(status string) models.PairStatus {
	switch status {
	case "TRADING":
		return models.Active
	case "BREAK":
		return models.Break
	case "HALT":
		return models.Halt
	case "END_OF_DAY":
		return models.EndOfDay
	default:
		return models.Delisted
	}
}
