package providers

import (
	"context"
	"strconv"
	"time"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
	"github.com/adshao/go-binance/v2"
)

// BinanceRepo required provider dependencies
type BinanceRepo struct {
	client *binance.Client
}

// NewBinanceRepo constructor for new binance repository
func NewBinanceRepo(apiKey string, secretKey string) *BinanceRepo {
	client := binance.NewClient(apiKey, secretKey)

	return &BinanceRepo{
		client: client,
	}
}

// ListAllPairs return a list of all binance pairs
func (br *BinanceRepo) ListAllPairs() ([]models.PairBase, error) {
	// retrieve data
	exchange, err := br.client.NewExchangeInfoService().Do(context.Background())

	// map data to standardized pair
	pairs := make([]models.PairBase, len(exchange.Symbols))
	for i, symbol := range exchange.Symbols {
		pairs[i] = models.PairBase{
			Base:     symbol.BaseAsset,
			Quote:    symbol.QuoteAsset,
			Exchange: models.Binance,
			Status:   matchBinancePairStatus(symbol.Status),
		}
	}

	return pairs, err
}

// FetchCandleData return candle data for a single pair
func (br *BinanceRepo) FetchCandleData(pair string, since time.Time) ([]models.CandleBase, error) {
	// override since time for dates before 2010
	if since.UnixMilli() < time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC).UnixMilli() {
		since = time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	// retrieve data
	data, err := br.client.NewKlinesService().Symbol(pair).StartTime(since.UnixMilli()).Limit(1000).Interval("1m").Do(context.Background())

	// map data to standardized candle
	candles := make([]models.CandleBase, len(data))
	for i, candle := range data {
		openF, _ := strconv.ParseFloat(candle.Open, 64)
		closeF, _ := strconv.ParseFloat(candle.Close, 64)
		highF, _ := strconv.ParseFloat(candle.High, 64)
		lowF, _ := strconv.ParseFloat(candle.Low, 64)
		volumeF, _ := strconv.ParseFloat(candle.Volume, 64)

		// append data
		candles[i] = models.CandleBase{
			OpenTime:  time.Unix(candle.OpenTime, 0),
			CloseTime: time.Unix(candle.CloseTime, 0),
			Open:      openF,
			Close:     closeF,
			High:      highF,
			Low:       lowF,
			Volume:    volumeF,
			NoTrade:   uint64(candle.TradeNum),
		}
	}

	return candles, err
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
