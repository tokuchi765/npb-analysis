package csv

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	teamData "github.com/tokuchi765/npb-analysis/entity/team"
)

func TestTeamReader_ReadTeamLeagueStats(t *testing.T) {
	type args struct {
		league string
		year   string
	}
	tests := []struct {
		name                 string
		args                 args
		wantTeamLeagueStats  []teamData.TeamLeagueStats
		wantTeamMatchResults []teamData.TeamMatchResults
	}{
		{
			"パリーグ成績読み込み",
			args{
				"p",
				"2005",
			},
			[]teamData.TeamLeagueStats{
				{TeamID: "10", Year: "2005", Manager: "", Games: 136, Win: 84, Lose: 49, Draw: 3, WinningRate: 0.632, ExchangeWin: 24, ExchangeLose: 11, ExchangeDraw: 1, HomeWin: 38, HomeLose: 30, HomeDraw: 0, LoadWin: 46, LoadLose: 19, LoadDraw: 3, PythagoreanExpectation: 0},
				{TeamID: "08", Year: "2005", Manager: "", Games: 136, Win: 89, Lose: 45, Draw: 2, WinningRate: 0.664, ExchangeWin: 23, ExchangeLose: 12, ExchangeDraw: 1, HomeWin: 45, HomeLose: 21, HomeDraw: 2, LoadWin: 44, LoadLose: 24, LoadDraw: 0, PythagoreanExpectation: 0},
				{TeamID: "07", Year: "2005", Manager: "", Games: 136, Win: 67, Lose: 69, Draw: 0, WinningRate: 0.493, ExchangeWin: 18, ExchangeLose: 18, ExchangeDraw: 0, HomeWin: 36, HomeLose: 32, HomeDraw: 0, LoadWin: 31, LoadLose: 37, LoadDraw: 0, PythagoreanExpectation: 0},
				{TeamID: "12", Year: "2005", Manager: "", Games: 136, Win: 62, Lose: 70, Draw: 4, WinningRate: 0.47, ExchangeWin: 17, ExchangeLose: 16, ExchangeDraw: 3, HomeWin: 30, HomeLose: 37, HomeDraw: 1, LoadWin: 32, LoadLose: 33, LoadDraw: 3, PythagoreanExpectation: 0},
				{TeamID: "11", Year: "2005", Manager: "", Games: 136, Win: 62, Lose: 71, Draw: 3, WinningRate: 0.466, ExchangeWin: 12, ExchangeLose: 22, ExchangeDraw: 2, HomeWin: 34, HomeLose: 31, HomeDraw: 3, LoadWin: 28, LoadLose: 40, LoadDraw: 0, PythagoreanExpectation: 0},
				{TeamID: "09", Year: "2005", Manager: "", Games: 136, Win: 38, Lose: 97, Draw: 1, WinningRate: 0.281, ExchangeWin: 11, ExchangeLose: 25, ExchangeDraw: 0, HomeWin: 21, HomeLose: 46, HomeDraw: 1, LoadWin: 17, LoadLose: 51, LoadDraw: 0, PythagoreanExpectation: 0},
			},
			[]teamData.TeamMatchResults{
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 12, Lose: 8, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 10, Lose: 10, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 14, Lose: 5, Draw: 1},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 11, Lose: 8, Draw: 1},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 13, Lose: 7, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 11, Lose: 9, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 17, Lose: 3, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 10, Lose: 10, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 15, Lose: 5, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 13, Lose: 6, Draw: 1},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 9, Lose: 11, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 13, Lose: 7, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 8, Lose: 12, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 10, Lose: 10, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 9, Lose: 11, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 11, Lose: 9, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 6, Lose: 13, Draw: 1},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 14, Lose: 6, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 7, Lose: 13, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 7, Lose: 13, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 10, Lose: 10, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 5, Lose: 15, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 14, Lose: 6, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 8, Lose: 11, Draw: 1},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 13, Lose: 7, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 7, Lose: 13, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "08", VsType: "league", Win: 3, Lose: 17, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "09", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "10", VsType: "league", Win: 5, Lose: 14, Draw: 1},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "11", VsType: "league", Win: 6, Lose: 14, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "12", VsType: "league", Win: 6, Lose: 14, Draw: 0},
			},
		},
		{"セリーグ成績読み込み",
			args{
				"c",
				"2005",
			},
			[]teamData.TeamLeagueStats{
				{TeamID: "03", Year: "2005", Manager: "", Games: 146, Win: 87, Lose: 54, Draw: 5, WinningRate: 0.617, ExchangeWin: 21, ExchangeLose: 13, ExchangeDraw: 2, HomeWin: 42, HomeLose: 26, HomeDraw: 5, LoadWin: 45, LoadLose: 28, LoadDraw: 0, PythagoreanExpectation: 0},
				{TeamID: "05", Year: "2005", Manager: "", Games: 146, Win: 79, Lose: 66, Draw: 1, WinningRate: 0.545, ExchangeWin: 15, ExchangeLose: 21, ExchangeDraw: 0, HomeWin: 42, HomeLose: 31, HomeDraw: 0, LoadWin: 37, LoadLose: 35, LoadDraw: 1, PythagoreanExpectation: 0},
				{TeamID: "02", Year: "2005", Manager: "", Games: 146, Win: 69, Lose: 70, Draw: 7, WinningRate: 0.496, ExchangeWin: 19, ExchangeLose: 17, ExchangeDraw: 0, HomeWin: 39, HomeLose: 32, HomeDraw: 2, LoadWin: 30, LoadLose: 38, LoadDraw: 5, PythagoreanExpectation: 0},
				{TeamID: "06", Year: "2005", Manager: "", Games: 146, Win: 71, Lose: 73, Draw: 2, WinningRate: 0.493, ExchangeWin: 20, ExchangeLose: 16, ExchangeDraw: 0, HomeWin: 40, HomeLose: 33, HomeDraw: 0, LoadWin: 31, LoadLose: 40, LoadDraw: 2, PythagoreanExpectation: 0},
				{TeamID: "01", Year: "2005", Manager: "", Games: 146, Win: 62, Lose: 80, Draw: 4, WinningRate: 0.437, ExchangeWin: 18, ExchangeLose: 14, ExchangeDraw: 4, HomeWin: 34, HomeLose: 38, HomeDraw: 1, LoadWin: 28, LoadLose: 42, LoadDraw: 3, PythagoreanExpectation: 0},
				{TeamID: "04", Year: "2005", Manager: "", Games: 146, Win: 58, Lose: 84, Draw: 4, WinningRate: 0.408, ExchangeWin: 11, ExchangeLose: 24, ExchangeDraw: 1, HomeWin: 31, HomeLose: 39, HomeDraw: 3, LoadWin: 27, LoadLose: 45, LoadDraw: 1, PythagoreanExpectation: 0},
			},
			[]teamData.TeamMatchResults{
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 14, Lose: 8, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 13, Lose: 6, Draw: 3},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 16, Lose: 6, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 11, Lose: 11, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 12, Lose: 10, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 14, Lose: 8, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 13, Lose: 8, Draw: 1},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 11, Lose: 11, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 14, Lose: 8, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 12, Lose: 10, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 16, Lose: 6, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 6, Lose: 13, Draw: 3},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 9, Lose: 11, Draw: 2},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 8, Lose: 13, Draw: 1},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 11, Lose: 10, Draw: 1},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 10, Lose: 12, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 10, Lose: 11, Draw: 1},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 10, Lose: 12, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 11, Lose: 10, Draw: 1},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 10, Lose: 12, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 6, Lose: 16, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 8, Lose: 14, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 10, Lose: 12, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 8, Lose: 14, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 12, Lose: 10, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "01", VsType: "league", Win: 12, Lose: 10, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "02", VsType: "league", Win: 11, Lose: 9, Draw: 2},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "03", VsType: "league", Win: 6, Lose: 16, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "04", VsType: "league", Win: 0, Lose: 0, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "05", VsType: "league", Win: 8, Lose: 14, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 10, Lose: 11, Draw: 1},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	teamReader := new(TeamReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLeagueStats, actualMatchResults := teamReader.ReadTeamLeagueStats(runtimeCurrent+"/test/resource", tt.args.league, tt.args.year)
			assert.Exactly(t, tt.wantTeamLeagueStats, actualLeagueStats)
			assert.Exactly(t, tt.wantTeamMatchResults, actualMatchResults)
		})
	}
}

func TestTeamReader_ReadManager(t *testing.T) {
	type args struct {
		teamID string
		year   string
	}
	tests := []struct {
		name        string
		args        args
		wantManager string
	}{
		{
			"監督情報読み込み",
			args{
				"01",
				"2006",
			},
			"原　辰徳",
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	teamReader := new(TeamReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := teamReader.ReadManager(runtimeCurrent+"/test/resource", tt.args.teamID, tt.args.year)
			assert.Equal(t, tt.wantManager, actual)
		})
	}
}

func TestTeamReader_ReadTeamExchangeStats(t *testing.T) {
	type args struct {
		league string
		year   string
	}
	tests := []struct {
		name                         string
		args                         args
		wantTeamExchangeMatchResults []teamData.TeamMatchResults
	}{
		{
			"セリーグ交流戦CSV読み込み",
			args{
				"c",
				"2005",
			},
			[]teamData.TeamMatchResults{
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 2, Lose: 3, Draw: 1},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "03", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 4, Lose: 1, Draw: 1},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "05", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 6, Lose: 0, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "02", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "06", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 4, Lose: 0, Draw: 2},
				{TeamID: "01", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 2, Lose: 2, Draw: 2},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "07", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "08", VsType: "exchange", Win: 2, Lose: 3, Draw: 1},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "09", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "11", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "04", Year: "2005", CompetitiveTeamID: "12", VsType: "exchange", Win: 1, Lose: 5, Draw: 0}},
		},
		{
			"パリーグ交流戦CSV読み込み",
			args{
				"p",
				"2005",
			},
			[]teamData.TeamMatchResults{
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 3, Lose: 2, Draw: 1},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "10", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 3, Lose: 2, Draw: 1},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
				{TeamID: "08", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "07", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 2, Lose: 2, Draw: 2},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 1, Lose: 4, Draw: 1},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
				{TeamID: "12", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 0, Lose: 4, Draw: 2},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 5, Lose: 1, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "11", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "01", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "02", VsType: "exchange", Win: 0, Lose: 6, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "03", VsType: "exchange", Win: 1, Lose: 5, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 2, Lose: 4, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "05", VsType: "exchange", Win: 4, Lose: 2, Draw: 0},
				{TeamID: "09", Year: "2005", CompetitiveTeamID: "06", VsType: "exchange", Win: 3, Lose: 3, Draw: 0},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	teamReader := new(TeamReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := teamReader.ReadTeamExchangeStats(runtimeCurrent+"/test/resource", tt.args.league, tt.args.year)
			assert.Exactly(t, tt.wantTeamExchangeMatchResults, actual)
		})
	}
}

func TestTeamReader_ReadTeamPitching(t *testing.T) {
	type args struct {
		league string
		year   string
	}
	tests := []struct {
		name             string
		args             args
		wantTeamPitching []teamData.TeamPitching
	}{
		{
			"セリーグ投手成績CSV読み込み",
			args{
				"central",
				"2005",
			},
			[]teamData.TeamPitching{{TeamID: "04", Year: "2005", EarnedRunAverage: 4.8, Games: 146, Win: 58, Lose: 84, Save: 27, Hold: 49, HoldPoint: 62, CompleteGame: 18, Shutout: 6, NoWalks: 8, WinningRate: 0.408, Batter: 5746, InningsPitched: 1286, Hit: 1379, HomeRun: 171, BaseOnBalls: 539, IntentionalWalk: 18, HitByPitches: 60, StrikeOut: 1041, WildPitches: 57, Balk: 9, RunsAllowed: 779, EarnedRun: 686}},
		},
		{
			"パリーグ投手成績CSV読み込み",
			args{
				"pacific",
				"2005",
			},
			[]teamData.TeamPitching{
				{TeamID: "08", Year: "2005", EarnedRunAverage: 3.46, Games: 136, Win: 89, Lose: 45, Save: 42, Hold: 83, HoldPoint: 102, CompleteGame: 22, Shutout: 14, NoWalks: 7, WinningRate: 0.664, Batter: 5122, InningsPitched: 1222, Hit: 1133, HomeRun: 107, BaseOnBalls: 380, IntentionalWalk: 20, HitByPitches: 63, StrikeOut: 1062, WildPitches: 28, Balk: 2, RunsAllowed: 504, EarnedRun: 470},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	teamReader := new(TeamReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := teamReader.ReadTeamPitching(runtimeCurrent+"/test/resource", tt.args.league, tt.args.year)
			assert.Exactly(t, tt.wantTeamPitching, actual)
		})
	}
}

func TestTeamReader_ReadTeamBatting(t *testing.T) {
	type args struct {
		league string
		year   string
	}
	tests := []struct {
		name            string
		args            args
		wantTeamBatting []teamData.TeamBatting
	}{
		{
			"セリーグ打撃成績CSV読み込み",
			args{
				"central",
				"2005",
			},
			[]teamData.TeamBatting{
				{TeamID: "06", Year: "2005", BattingAverage: 0.276, Games: 146, PlateAppearance: 5523, AtBat: 5033, Score: 591, Hit: 1389, Double: 221, Triple: 15, HomeRun: 128, BaseHit: 2024, RunsBattedIn: 565, StolenBase: 65, CaughtStealing: 25, SacrificeHits: 103, SacrificeFlies: 25, BaseOnBalls: 311, IntentionalWalk: 16, HitByPitches: 51, StrikeOut: 1016, GroundedIntoDoublePlay: 106, SluggingPercentage: 0.402, OnBasePercentage: 0.32299999999999995},
			},
		},
		{
			"パリーグ打撃成績CSV読み込み",
			args{
				"pacific",
				"2005",
			},
			[]teamData.TeamBatting{
				{TeamID: "09", Year: "2005", BattingAverage: 0.255, Games: 136, PlateAppearance: 5068, AtBat: 4577, Score: 504, Hit: 1166, Double: 209, Triple: 16, HomeRun: 88, BaseHit: 1671, RunsBattedIn: 474, StolenBase: 41, CaughtStealing: 34, SacrificeHits: 70, SacrificeFlies: 30, BaseOnBalls: 347, IntentionalWalk: 5, HitByPitches: 44, StrikeOut: 919, GroundedIntoDoublePlay: 124, SluggingPercentage: 0.365, OnBasePercentage: 0.312},
			},
		},
	}
	runtimeCurrent, _ := filepath.Abs("../../")
	teamReader := new(TeamReader)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := teamReader.ReadTeamBatting(runtimeCurrent+"/test/resource", tt.args.league, tt.args.year)
			assert.Exactly(t, tt.wantTeamBatting, actual)
		})
	}
}
