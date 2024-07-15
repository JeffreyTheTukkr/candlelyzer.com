package models

import (
	"time"

	"github.com/google/uuid"
)

// ExchangeName enum for the exchange type
type ExchangeName string

// enum options for the exchange name
const (
	Binance     ExchangeName = "binance"
	Depreciated ExchangeName = "depreciated"
)

// PairStatus enum for the pair (trading) status
type PairStatus string

// enum options for the pair (trading) statuses
const (
	Active   PairStatus = "active"     // actively trading
	Break    PairStatus = "break"      // temporary maintenance
	Halt     PairStatus = "halt"       // emergency shutdown
	EndOfDay PairStatus = "end_of_day" // end of day
	Delisted PairStatus = "delisted"   // removed from exchange
)

// Pair database model for pair structure
type Pair struct {
	Id        uuid.UUID    `json:"id"`
	Base      string       `json:"base"`
	Quote     string       `json:"quote"`
	Exchange  ExchangeName `json:"exchange"`
	Status    PairStatus   `json:"status"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// PairBase simplified pair model for filters and inserts
type PairBase struct {
	Base     string       `json:"base"`
	Quote    string       `json:"quote"`
	Exchange ExchangeName `json:"exchange"`
	Status   PairStatus   `json:"status"`
}
