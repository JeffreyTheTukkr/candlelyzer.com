package models

import (
	"time"

	"github.com/google/uuid"
)

// Candle database model for candle structure
type Candle struct {
	Id        uuid.UUID `json:"id"`
	Pair      uuid.UUID `json:"pair"`
	OpenTime  time.Time `json:"open_time"`
	CloseTime time.Time `json:"close_time"`
	Open      float64   `json:"open"`
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Volume    float64   `json:"volume"`
	NoTrade   uint64    `json:"no_trade"`
	CreatedAt time.Time `json:"created_at"`
}
