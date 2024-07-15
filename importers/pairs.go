package importers

import (
	"log/slog"

	"github.com/JeffreyTheTukkr/candlelyzer.com/providers"
	"github.com/JeffreyTheTukkr/candlelyzer.com/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PairsImporter required importer dependencies
type PairsImporter struct {
	db            *pgxpool.Pool
	logger        *slog.Logger
	binanceApiKey string
	binanceSecret string
}

// NewPairsImporter constructor for new pairs importer
func NewPairsImporter(db *pgxpool.Pool, logger *slog.Logger, binanceApiKey string, binanceSecret string) *PairsImporter {
	return &PairsImporter{
		db:            db,
		logger:        logger,
		binanceApiKey: binanceApiKey,
		binanceSecret: binanceSecret,
	}
}

// RunPairsImport run the pair importer
func (pi *PairsImporter) RunPairsImport() {
	// fetch binance pairs
	binanceRepo := providers.NewBinanceRepo(pi.binanceApiKey, pi.binanceSecret)
	pairs, err := binanceRepo.ListAllPairs()
	if err != nil {
		pi.logger.Error("failed to retrieve binance pairs", "error", err)
	}

	// create pairs repository
	pairsRepo := repositories.NewPairRepo(pi.db)

	// import pairs into database
	for _, pair := range pairs {
		err := pairsRepo.UpsertOne(pair)
		if err != nil {
			pi.logger.Error("failed to insert binance pair", "symbol", pair.Base+pair.Quote, "error", err)
		}
	}
}
