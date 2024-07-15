package providers

import (
	"testing"

	"github.com/JeffreyTheTukkr/candlelyzer.com/models"
)

func Test_matchBinancePairStatus(t *testing.T) {
	tests := []struct {
		name   string
		status string
		want   models.PairStatus
	}{
		{"trading status", "TRADING", models.Active},
		{"break status", "BREAK", models.Break},
		{"halt status", "HALT", models.Halt},
		{"end of day status", "END_OF_DAY", models.EndOfDay},
		{"default", "DEFAULT", models.Delisted},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchBinancePairStatus(tt.status); got != tt.want {
				t.Errorf("matchBinancePairStatus() returned %v, wants %v", got, tt.want)
			}
		})
	}
}
