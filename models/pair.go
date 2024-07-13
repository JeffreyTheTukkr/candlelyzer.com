package models

import "time"

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
	Id        uint32       `json:"id"`
	Base      string       `json:"base"`
	Quote     string       `json:"quote"`
	Exchange  ExchangeName `json:"exchange"`
	Status    PairStatus   `json:"status"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// PairFilters model to contain the model filters
type PairFilters struct {
	Base     string       `json:"base"`
	Quote    string       `json:"quote"`
	Exchange ExchangeName `json:"exchange"`
	Status   PairStatus   `json:"status"`
}
