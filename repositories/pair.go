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

// Pair repository manager for the pair model
type Pair struct {
	db *pgxpool.Pool
}

// NewPair repository constructor for a new pair model repo
func NewPair(db *pgxpool.Pool) *Pair {
	return &Pair{
		db: db,
	}
}

// InsertOne insert a single pair base model
func (p *Pair) InsertOne(m models.PairBase) error {
	_, err := p.db.Exec(context.Background(), "INSERT INTO pairs (base, quote, exchange, status) VALUES ($1, $2, $3, $4);",
		m.Base, m.Quote, m.Exchange, m.Status)

	return err
}

// FindById find a single pair model by id
func (p *Pair) FindById(id uuid.UUID) (models.Pair, error) {
	rows, err := p.db.Query(context.Background(), "SELECT * FROM pairs WHERE id = $1;", id)
	if err != nil {
		return models.Pair{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Pair])
}

// UpsertOne update or insert a single pair model
func (p *Pair) UpsertOne(m models.PairBase) error {
	_, err := p.db.Exec(context.Background(), "INSERT INTO pairs (base, quote, exchange, status) VALUES ($1, $2, $3, $4) "+
		"ON CONFLICT (base, quote, exchange) DO UPDATE "+
		"SET status = EXCLUDED.status WHERE pairs.status IS DISTINCT FROM EXCLUDED.status;",
		m.Base, m.Quote, m.Exchange, m.Status)

	return err
}

// UpdateStatusByIds update multiple pair models status by ids
func (p *Pair) UpdateStatusByIds(ids []uuid.UUID, status models.PairStatus) error {
	idList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")
	_, err := p.db.Exec(context.Background(), "UPDATE pairs SET status = $1 WHERE id IN ($2);", status, idList)

	return err
}
