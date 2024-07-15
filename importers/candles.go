package importers

import (
	"log/slog"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
	"github.com/JeffreyTheTukkr/candlelyzer.com/providers"
	"github.com/JeffreyTheTukkr/candlelyzer.com/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CandlesImporter required importer dependencies
type CandlesImporter struct {
	db            *pgxpool.Pool
	logger        *slog.Logger
	binanceApiKey string
	binanceSecret string
}

// NewCandlesImporter constructor for new candles importer
func NewCandlesImporter(db *pgxpool.Pool, logger *slog.Logger, binanceApiKey string, binanceSecret string) *CandlesImporter {
	return &CandlesImporter{
		db:            db,
		logger:        logger,
		binanceApiKey: binanceApiKey,
		binanceSecret: binanceSecret,
	}
}

// RunCandlesImport run the candle importer
func (ci *CandlesImporter) RunCandlesImport() {
	// fetch actively importing database pairs
	pairsRepo := repositories.NewPairRepo(ci.db)
	pairs, err := pairsRepo.FindActivelyImporting(models.Binance)
	if err != nil {
		ci.logger.Error("failed to retrieve binance pairs", "error", err)
	}

	// create candles repository
	candleRepo := repositories.NewCandleRepo(ci.db)

	// create binance provider
	binanceProvider := providers.NewBinanceRepo(ci.binanceApiKey, ci.binanceSecret)

	// fetch and import candles for each pair
	for _, pair := range pairs {
		// fetch last candle close time
		lastCloseTime, err := candleRepo.FindLastCloseTime(pair.Id)
		if err != nil {
			ci.logger.Error("failed to retrieve last close time model", "error", err)
			continue
		}

		// retrieve candle data
		candles, err := binanceProvider.FetchCandleData(pair.Base+pair.Quote, lastCloseTime)
		if err != nil {
			ci.logger.Error("failed to fetch last candle close_time", "error", err)
			continue
		}

		// import candle data
		for _, candle := range candles {
			err := candleRepo.InsertOne(models.CandleBase{
				Pair:      pair.Id,
				OpenTime:  candle.OpenTime,
				CloseTime: candle.CloseTime,
				Open:      candle.Open,
				Close:     candle.Close,
				High:      candle.High,
				Low:       candle.Low,
				Volume:    candle.Volume,
				NoTrade:   candle.NoTrade,
			})

			if err != nil {
				ci.logger.Error("failed to insert candle", "symbol", pair.Base+pair.Quote, "error", err)
			}
		}
	}
}
