package main

import (
	"github.com/JeffreyTheTukkr/candlelyzer.com/databases"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

// Application struct to contain global submodules
type Application struct {
	DB *pgxpool.Pool
}

// NewApplication return a new application instance
func NewApplication(db *pgxpool.Pool) *Application {
	return &Application{
		DB: db,
	}
}

// main bootstrap application
func main() {
	// create psql database instance
	db, dbErr := databases.CreatePsqlPool()
	if dbErr != nil {
		panic(dbErr)
	}

	// create new application
	_ = NewApplication(db)
}
