package cron

import (
	"github.com/JeffreyTheTukkr/candlelyzer.com/importers"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

// ImportManager required importer manager dependencies
type ImportManager struct {
	db            *pgxpool.Pool
	logger        *slog.Logger
	binanceApiKey string
	binanceSecret string
}

// NewImportManager constructor for new import manager
func NewImportManager(db *pgxpool.Pool, logger *slog.Logger, binanceApiKey string, binanceSecret string) *ImportManager {
	return &ImportManager{
		db:            db,
		logger:        logger,
		binanceApiKey: binanceApiKey,
		binanceSecret: binanceSecret,
	}
}

// StartImportManager start the cron import manager
func (im *ImportManager) StartImportManager() {
	cron, err := gocron.NewScheduler()
	if err != nil {
		im.logger.Error("failed to start import manager", "error", err)
		panic(err)
	}

	// create importer instances
	pairsImporter := importers.NewPairsImporter(im.db, im.logger, im.binanceApiKey, im.binanceSecret)
	candlesImporter := importers.NewCandlesImporter(im.db, im.logger, im.binanceApiKey, im.binanceSecret)

	// add pair importer job
	pairJob, _ := cron.NewJob(gocron.CronJob("30 * */8 * * *", true), gocron.NewTask(func() {
		im.logger.Info("running pairs importer...")
		pairsImporter.RunPairsImport()
	}))

	// add candle importer job
	candleJob, _ := cron.NewJob(gocron.CronJob("0 * * * * *", true), gocron.NewTask(func() {
		im.logger.Info("running candle importer...")
		candlesImporter.RunCandlesImport()
	}))

	// start cron manager and run imports
	cron.Start()
	_ = pairJob.RunNow()
	_ = candleJob.RunNow()
	select {}
}
