package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PairRepo required repository dependencies
type PairRepo struct {
	db *pgxpool.Pool
}

// NewPairRepo constructor for new pair repository
func NewPairRepo(db *pgxpool.Pool) *PairRepo {
	return &PairRepo{
		db: db,
	}
}

// InsertOne insert a single pair base model
func (pr *PairRepo) InsertOne(m models.PairBase) error {
	_, err := pr.db.Exec(context.Background(), "INSERT INTO pairs (base, quote, exchange, status) VALUES ($1, $2, $3, $4);",
		m.Base, m.Quote, m.Exchange, m.Status)

	return err
}

// FindById find a single pair model by id
func (pr *PairRepo) FindById(id uuid.UUID) (models.Pair, error) {
	rows, err := pr.db.Query(context.Background(), "SELECT * FROM pairs WHERE id = $1;", id)
	if err != nil {
		return models.Pair{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Pair])
}

// FindActivelyImporting find multiple pairs with importing status true
func (pr *PairRepo) FindActivelyImporting() ([]models.Pair, error) {
	rows, err := pr.db.Query(context.Background(), "SELECT * FROM pairs WHERE importing = true AND status = 'active';")
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Pair])
}

// UpsertOne update or insert a single pair model
func (pr *PairRepo) UpsertOne(m models.PairBase) error {
	_, err := pr.db.Exec(context.Background(), "INSERT INTO pairs (base, quote, exchange, status) VALUES ($1, $2, $3, $4) "+
		"ON CONFLICT (base, quote, exchange) DO UPDATE "+
		"SET status = EXCLUDED.status WHERE pairs.status IS DISTINCT FROM EXCLUDED.status;",
		m.Base, m.Quote, m.Exchange, m.Status)

	return err
}

// UpdateStatusByIds update multiple pair models status by ids
func (pr *PairRepo) UpdateStatusByIds(ids []uuid.UUID, status models.PairStatus) error {
	idList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")
	_, err := pr.db.Exec(context.Background(), "UPDATE pairs SET status = $1 WHERE id IN ($2);", status, idList)

	return err
}
