package csv

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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
