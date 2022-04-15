package team

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamBatting_SetBABIP(t *testing.T) {
	tests := []struct {
		name        string
		teamBatting *TeamBatting
		wantBABIP   float64
	}{
		{
			"BABIP算出",
			&TeamBatting{
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
			tt.teamBatting.SetBABIP()
			assert.Equal(t, tt.wantBABIP, tt.teamBatting.BABIP)
		})
	}
}

func TestTeamPitching_SetBABIP(t *testing.T) {
	tests := []struct {
		name         string
		teamPitching *TeamPitching
		wantBABIP    float64
	}{
		{
			"被BABIP算出",
			&TeamPitching{
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
			tt.teamPitching.SetBABIP()
			assert.Equal(t, tt.wantBABIP, tt.teamPitching.BABIP)
		})
	}
}

func TestTeamPitching_SetStrikeOutRate(t *testing.T) {
	tests := []struct {
		name                 string
		teamPitching         *TeamPitching
		wantSetStrikeOutRate float64
	}{
		{
			"奪三振率算出",
			&TeamPitching{
				StrikeOut:      10,
				InningsPitched: 30,
			},
			3.0,
		},
		{
			"奪三振率がNaN",
			&TeamPitching{
				StrikeOut:      0,
				InningsPitched: 0,
			},
			0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.teamPitching.SetStrikeOutRate()
			assert.Equal(t, tt.wantSetStrikeOutRate, tt.teamPitching.StrikeOutRate)
		})
	}
}
