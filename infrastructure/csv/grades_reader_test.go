package csv

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/player"
	data "github.com/tokuchi765/npb-analysis/entity/player"
)

func TestGradesReader_GetPlayers(t *testing.T) {
	type args struct {
		initial string
	}
	tests := []struct {
		name        string
		args        args
		wantPlayers [][]string
	}{
		{
			"選手一覧",
			args{
				"g",
			},
			[][]string{
				{"/bis/players/93795138.html", "デラロサ"},
				{"/bis/players/41045138.html", "戸郷　翔征"},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	gradesReader := new(GradesReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlayers := gradesReader.GetPlayers(runtimeCurrent+"/test/resource", tt.args.initial)
			assert.ElementsMatch(t, tt.wantPlayers, gotPlayers)
		})
	}
}

func TestGradesReader_ReadCareer(t *testing.T) {
	type args struct {
		initial    string
		playerID   string
		playerName string
	}
	tests := []struct {
		name       string
		args       args
		wantCareer data.CAREER
	}{
		{
			"選手成績読み込み",
			args{
				initial:    "b",
				playerID:   "01105137",
				playerName: "飯田　優也",
			},
			data.CAREER{
				PlayerID:           "01105137",
				Name:               "飯田　優也",
				Position:           "投手",
				PitchingAndBatting: "左投左打",
				Height:             "187cm",
				Weight:             "92kg",
				Birthday:           "1990年11月27日",
				Career:             "神戸弘陵高 - 東京農業大生産学部",
				Draft:              "2012年育成選手ドラフト3位",
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	gradesReader := new(GradesReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, exsist := gradesReader.ReadCareer(runtimeCurrent+"/test/resource", tt.args.initial, tt.args.playerID, tt.args.playerName)
			assert.Exactly(t, tt.wantCareer, actual)
			assert.Equal(t, true, exsist)
		})
	}
}

func TestGradesReader_ReadGrades(t *testing.T) {
	type args struct {
		initial    string
		playerID   string
		playerName string
	}
	tests := []struct {
		name                 string
		args                 args
		wantPicherGradesList []data.PICHERGRADES
		wantBatterGradesList []data.BATTERGRADES
		wantExsist           bool
	}{
		{
			"投手成績読み込み",
			args{
				"b",
				"53355134",
				"山本　由伸",
			},
			[]data.PICHERGRADES{getTestPicherGrades()},
			[]player.BATTERGRADES(nil),
			true,
		},
		{
			"野手成績読み込み",
			args{
				"b",
				"01605136",
				"福田　周平",
			},
			[]data.PICHERGRADES(nil),
			[]player.BATTERGRADES{getTestBatterGrades()},
			true,
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	gradesReader := new(GradesReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPicherGradesList, gotBatterGradesList, gotExsist := gradesReader.ReadGrades(runtimeCurrent+"/test/resource", tt.args.initial, tt.args.playerID, tt.args.playerName)
			assert.Exactly(t, tt.wantPicherGradesList, gotPicherGradesList)
			assert.Exactly(t, tt.wantBatterGradesList, gotBatterGradesList)
			assert.Equal(t, tt.wantExsist, gotExsist)
		})
	}
}

func getTestPicherGrades() data.PICHERGRADES {
	return data.PICHERGRADES{
		Year:             "2018",
		TeamID:           "12",
		Team:             "オリックス",
		Piched:           54.0,
		Win:              4.0,
		Lose:             2.0,
		Save:             1.0,
		Hold:             32.0,
		HoldPoint:        36.0,
		CompleteGame:     0.0,
		Shutout:          0.0,
		NoWalks:          0.0,
		WinningRate:      0.667,
		Batter:           213.0,
		InningsPitched:   53.0,
		Hit:              40.0,
		HomeRun:          4.0,
		BaseOnBalls:      16.0,
		HitByPitches:     2.0,
		StrikeOut:        46.0,
		WildPitches:      2.0,
		Balk:             0.0,
		RunsAllowed:      19.0,
		EarnedRun:        17.0,
		EarnedRunAverage: 2.89,
	}
}

func getTestBatterGrades() data.BATTERGRADES {
	return data.BATTERGRADES{
		Year:                   "2018",
		TeamID:                 "12",
		Team:                   "オリックス",
		Games:                  113,
		PlateAppearance:        345,
		AtBat:                  295,
		Score:                  39,
		Hit:                    78,
		Single:                 0,
		Double:                 8,
		Triple:                 4,
		HomeRun:                1,
		BaseHit:                97,
		RunsBattedIn:           15,
		StolenBase:             16,
		CaughtStealing:         9,
		SacrificeHits:          16,
		SacrificeFlies:         0,
		BaseOnBalls:            31,
		HitByPitches:           3,
		StrikeOut:              33,
		GroundedIntoDoublePlay: 2,
		BattingAverage:         0.264,
		SluggingPercentage:     0.32899999999999996,
		OnBasePercentage:       0.34,
		Woba:                   0.0,
	}
}
