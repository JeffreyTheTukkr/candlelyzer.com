package databases

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CreatePsqlPool initializes a new postgresql pool
func CreatePsqlPool() (*pgxpool.Pool, error) {
	connString := getConnectionString()

	// create connection pool and connect to database
	dbPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Printf("unable to connect to psql database: %v\n", err)
		return nil, err
	}

	// ping database to ensure valid connection
	err = dbPool.Ping(context.Background())
	if err != nil {
		fmt.Printf("unable to ping psql database: %v\n", err)
		return nil, err
	}

	return dbPool, nil
}

// getConnectionString returns the database connection string based on the environment variables
func getConnectionString() string {
	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASS")
	dbName := os.Getenv("PG_NAME")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
