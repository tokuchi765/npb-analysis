package csv

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/player"
)

func TestGradesReader_GetPlayers(t *testing.T) {
	type args struct {
		initial string
		year    string
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
				"2020",
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
			gotPlayers := gradesReader.GetPlayers(runtimeCurrent+"/test/resource", tt.args.initial, tt.args.year)
			assert.ElementsMatch(t, tt.wantPlayers, gotPlayers)
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
		wantPicherGradesList []player.PICHERGRADES
		wantBatterGradesList []player.BATTERGRADES
		wantExsist           bool
	}{
		{
			"投手成績読み込み",
			args{
				"b",
				"53355134",
				"山本　由伸",
			},
			[]player.PICHERGRADES{getTestPicherGrades()},
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
			[]player.PICHERGRADES(nil),
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

func getTestPicherGrades() player.PICHERGRADES {
	return player.PICHERGRADES{
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

func getTestBatterGrades() player.BATTERGRADES {
	return player.BATTERGRADES{
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

func TestGradesReader_ReadCareers(t *testing.T) {
	tests := []struct {
		name        string
		wantCareers []player.CAREER
	}{
		{
			"",
			[]player.CAREER{
				{
					PlayerID:           "01005112",
					Name:               "田中 靖洋",
					Position:           "",
					PitchingAndBatting: "右投右打",
					Height:             "183cm",
					Weight:             "88kg",
					Birthday:           "1987年6月21日",
					Career:             "加賀高",
					Draft:              "2005年高校生ドラフト4巡目",
				},
				{
					PlayerID:           "01005130",
					Name:               "高濱 祐仁",
					Position:           "外野手",
					PitchingAndBatting: "右投右打",
					Height:             "185cm",
					Weight:             "88kg",
					Birthday:           "1996年8月8日",
					Career:             "横浜高",
					Draft:              "2014年ドラフト7位",
				},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	gradesReader := new(GradesReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gradesReader.ReadCareers(runtimeCurrent + "/test/resource")
			assert.ElementsMatch(t, tt.wantCareers, actual)
		})
	}
}
