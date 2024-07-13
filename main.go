package main

import (
	"log/slog"

	"github.com/JeffreyTheTukkr/candlelyzer.com/databases"
	"github.com/JeffreyTheTukkr/candlelyzer.com/loggers"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

// Application struct to contain global submodules
type Application struct {
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

// NewApplication return a new application instance
func NewApplication(db *pgxpool.Pool, logger *slog.Logger) *Application {
	return &Application{
		DB:     db,
		Logger: logger,
	}
}

// main bootstrap application
func main() {
	// create psql database instance
	db, dbErr := databases.NewPsqlPool()
	if dbErr != nil {
		panic(dbErr)
	}

	// run psql database migrations
	migratorErr := databases.RunPsqlMigrations(db)
	if migratorErr != nil {
		panic(migratorErr)
	}

	// create slog logger instance
	logger := loggers.NewSlogLogger()

	// create new application and start
	app := NewApplication(db, logger)
	app.start()
}

// start run the application
func (app *Application) start() {
	app.Logger.Info("starting application..")
}
