package repositories

import (
	"context"
	"time"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CandleRepo required repository dependencies
type CandleRepo struct {
	db *pgxpool.Pool
}

// NewCandleRepo constructor for new candle repository
func NewCandleRepo(db *pgxpool.Pool) *CandleRepo {
	return &CandleRepo{
		db: db,
	}
}

// InsertOne insert a single candle base model
func (cr *CandleRepo) InsertOne(m models.CandleBase) error {
	_, err := cr.db.Exec(context.Background(), "INSERT INTO candles (pair_id, open_time, close_time, open, close, high, low, volume, no_trades) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING;",
		m.Pair, m.OpenTime, m.CloseTime, m.Open, m.Close, m.High, m.Low, m.Volume, m.NoTrade)

	return err
}

// FindById find a single candle model by id
func (cr *CandleRepo) FindById(id uuid.UUID) (models.Candle, error) {
	rows, err := cr.db.Query(context.Background(), "SELECT * FROM candles WHERE id = $1;", id)
	if err != nil {
		return models.Candle{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Candle])
}

// FindLastCloseTime find the last close_time based on the pair_id
// note: close_time is used as this will be passed on to binance to search from the last open_time
// this to prevent receiving the same open_time entry from binance, thus having to check a double entry before inserting into the database
func (cr *CandleRepo) FindLastCloseTime(pairId uuid.UUID) (time.Time, error) {
	rows, err := cr.db.Query(context.Background(), "SELECT close_time FROM candles WHERE pair_id = $1 ORDER BY close_time DESC;", pairId)
	if err != nil {
		return time.Time{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowTo[time.Time])
}
