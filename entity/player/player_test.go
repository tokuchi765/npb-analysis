package player

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/sqlwrapper"
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

func TestBATTERGRADES_SetBABIP(t *testing.T) {
	tests := []struct {
		name         string
		batterGrades *BATTERGRADES
		wantBABIP    float64
	}{
		{
			"BABIP算出",
			&BATTERGRADES{
				AtBat:          500,
				Hit:            240,
				HomeRun:        20,
				StrikeOut:      50,
				SacrificeFlies: 10,
			},
			0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.batterGrades.SetBABIP()
			assert.Equal(t, tt.wantBABIP, tt.batterGrades.BABIP)
		})
	}
}

func TestPICHERGRADES_SetBABIP(t *testing.T) {
	tests := []struct {
		name         string
		picherGrades *PICHERGRADES
		wantBABIP    float64
	}{
		{
			"被BABIP算出",
			&PICHERGRADES{
				Batter:       500,
				Hit:          240,
				HomeRun:      20,
				StrikeOut:    20,
				BaseOnBalls:  10,
				HitByPitches: 10,
			},
			0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.picherGrades.SetBABIP()
			assert.Equal(t, tt.wantBABIP, tt.picherGrades.BABIP)
		})
	}
}

func TestPICHERGRADES_SetStrikeOutRate(t *testing.T) {
	tests := []struct {
		name                 string
		picherGrades         *PICHERGRADES
		wantSetStrikeOutRate float64
	}{
		{
			"奪三振率算出",
			&PICHERGRADES{
				StrikeOut:      10,
				InningsPitched: 30,
			},
			3.0,
		},
		{
			"奪三振率がNaN",
			&PICHERGRADES{
				StrikeOut:      0,
				InningsPitched: 0,
			},
			0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.picherGrades.SetStrikeOutRate()
			assert.Equal(t, tt.wantSetStrikeOutRate, tt.picherGrades.StrikeOutRate)
		})
	}
}

func TestBATTERGRADES_SetStrikeOutRate(t *testing.T) {
	tests := []struct {
		name              string
		batterGrades      *BATTERGRADES
		wantStrikeOutRate sqlwrapper.NullFloat64
	}{
		{
			"三振率算出",
			&BATTERGRADES{
				StrikeOut:       10,
				PlateAppearance: 100,
			},
			sqlwrapper.NullFloat64{
				NullFloat64: sql.NullFloat64{
					Float64: 0.1,
					Valid:   true,
				},
			},
		},
		{
			"三振率がNaN",
			&BATTERGRADES{
				StrikeOut:       0,
				PlateAppearance: 0,
			},
			sqlwrapper.NullFloat64{
				NullFloat64: sql.NullFloat64{
					Float64: 0.0,
					Valid:   true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.batterGrades.SetStrikeOutRate()
			assert.Equal(t, tt.wantStrikeOutRate, tt.batterGrades.StrikeOutRate)
		})
	}
}
