package databases

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RunPsqlMigrations run all database migrations
func RunPsqlMigrations(db *pgxpool.Pool) error {
	// retrieve all migration files
	files, filesErr := os.ReadDir("migrations")
	if filesErr != nil {
		fmt.Printf("unable to read migrations directory: %v\n", filesErr)
		return filesErr
	}

	// create migrations table if not exists
	_, migrationsErr := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS migrations (id SERIAL PRIMARY KEY, name VARCHAR(255), key INT NOT NULL UNIQUE, created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);")
	if migrationsErr != nil {
		fmt.Printf("unable to create migrations table: %v\n", migrationsErr)
		return migrationsErr
	}

	// retrieve already migrated keys from database
	rows, _ := db.Query(context.Background(), "SELECT key FROM migrations;")
	keys, _ := pgx.CollectRows(rows, pgx.RowTo[int])

	// loop over migration files
	for _, file := range files {
		// set variables
		key := file.Name()[:3]                       // first 3 digits of filename is the key
		keyInt, keyErr := strconv.Atoi(key)          // transform key to integer
		name := file.Name()[:len(file.Name())-4][4:] // skip first (`001-`) and last (`.sql`) 4 characters

		if keyErr != nil {
			fmt.Printf("failed to parse migration file key: %v\n", keyErr)
			return keyErr
		}

		// skip migrations if it already exists in the migrations table
		if slices.Contains(keys, keyInt) {
			continue
		}

		// read file contents
		content, contentErr := os.ReadFile(filepath.Join("migrations", file.Name()))
		if contentErr != nil {
			fmt.Printf("unable to read migration file: %v\n", contentErr)
			return contentErr
		}

		// execute file contents
		_, execErr := db.Exec(context.Background(), string(content))
		if execErr != nil {
			fmt.Printf("unable to execute migration file: %v\n", execErr)
			return execErr
		}

		// insert migration into database
		_, insertErr := db.Exec(context.Background(), "INSERT INTO migrations (name, key) VALUES ($1, $2);", name, key)
		if insertErr != nil {
			fmt.Printf("unable to insert migration file: %v\n", insertErr)
			return insertErr
		}
	}

	return nil
}
