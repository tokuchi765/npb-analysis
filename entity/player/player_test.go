package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBATTERGRADES_SetRC(t *testing.T) {
	tests := []struct {
		name         string
		batterGrades *BATTERGRADES
		wantRC       float64
	}{
		{
			"RC算出",
			&BATTERGRADES{
				Year:                   "2020",
				TeamID:                 "",
				Team:                   "",
				Games:                  119,
				PlateAppearance:        515,
				AtBat:                  427,
				Score:                  90,
				Hit:                    146,
				Single:                 0,
				Double:                 23,
				Triple:                 5,
				HomeRun:                29,
				BaseHit:                266,
				RunsBattedIn:           86,
				StolenBase:             7,
				CaughtStealing:         2,
				SacrificeHits:          0,
				SacrificeFlies:         3,
				BaseOnBalls:            84,
				HitByPitches:           1,
				StrikeOut:              103,
				GroundedIntoDoublePlay: 2,
				BattingAverage:         0.342,
				SluggingPercentage:     0.623,
				OnBasePercentage:       0.449,
				Woba:                   0.0,
				RC:                     0.0,
			},
			116.04369795037758,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.batterGrades.SetRC()
			assert.Equal(t, tt.wantRC, tt.batterGrades.RC)
		})
	}
}
