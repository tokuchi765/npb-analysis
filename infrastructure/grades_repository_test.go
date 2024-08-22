package infrastructure

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/player"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestGradesRepository_InsertPicherGrades_GetPitchings(t *testing.T) {
	tests := []struct {
		name    string
		pitcher player.PitcherGrades
	}{
		{
			"投手成績登録と取得",
			createPicherGradesMapping("00001", "2020", "01", "チーム名", 54.0, 4.0, 2.0, 1.0, 32.0, 36.0, 2.0, 3.0, 1.0, 0.667, 213.0, 53.0, 40.0, 4.0, 16.0, 2.0, 46.0, 2.0, 10.0, 19.0, 17.0, 2.89, 0.3, 3.6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			gorm := testUtil.ConnectGormDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.GormDB = gorm
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertPicherGrades(tt.pitcher)
			actual := repository.GetPitchings(tt.pitcher.PlayerID)
			assert.ElementsMatch(t, []player.PitcherGrades{tt.pitcher}, actual)
		})
	}
}

func createPicherGradesMapping(playerID string, year string, teamID string, team string, pitched float64, win float64, lose float64, save float64, hold float64, holdPoint float64, completeGame float64, shutout float64, noWalks float64, winningRate float64, batter float64, inningsPitched float64, hit float64, homeRun float64, baseOnBalls float64, hitByPitches float64, strikeOut float64, wildPitches float64, balk float64, runsAllowed float64, earnedRun float64, earnedRunAverage float64, babip float64, strikeOutRate float64) player.PitcherGrades {
	return player.PitcherGrades{
		PlayerID:         playerID,
		Year:             year,
		TeamID:           teamID,
		Team:             team,
		Pitched:          pitched,
		Win:              win,
		Lose:             lose,
		Save:             save,
		Hold:             hold,
		HoldPoint:        holdPoint,
		CompleteGame:     completeGame,
		Shutout:          shutout,
		NoWalks:          noWalks,
		WinningRate:      winningRate,
		Batter:           batter,
		InningsPitched:   inningsPitched,
		Hit:              hit,
		HomeRun:          homeRun,
		BaseOnBalls:      baseOnBalls,
		HitByPitches:     hitByPitches,
		StrikeOut:        strikeOut,
		WildPitches:      wildPitches,
		Balk:             balk,
		RunsAllowed:      runsAllowed,
		EarnedRun:        earnedRun,
		EarnedRunAverage: earnedRunAverage,
		Babip:            babip,
		StrikeOutRate:    strikeOutRate,
	}
}

func TestGradesRepository_InsertBatterGrades_GetBattings(t *testing.T) {
	tests := []struct {
		name    string
		batting player.BatterGrades
	}{
		{
			"打者成績登録と取得",
			createBatterGradesMapping("01605136", "2018", "12", "オリックス", 113, 345, 295, 39, 78, 0, 8, 4, 1, 97, 15, 16, 9, 16, 0, 31, 3, 33, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			gorm := testUtil.ConnectGormDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.GormDB = gorm
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertBatterGrades(tt.batting)
			actual := repository.GetBattings(tt.batting.PlayerID)
			assert.Equal(t, []player.BatterGrades{tt.batting}, actual)
		})
	}
}

func createBatterGradesMapping(PlayerID string, Year string, TeamID string, Team string, Games int, PlateAppearance int, AtBat int, Score int, Hit int, Single int, Double int, Triple int, HomeRun int, BaseHit int, RunsBattedIn int, StolenBase int, CaughtStealing int, SacrificeHits int, SacrificeFlies int, BaseOnBalls int, HitByPitches int, StrikeOut int, StrikeOutRate float64, GroundedIntoDoublePlay int, BattingAverage float64, SluggingPercentage float64, OnBasePercentage float64, Woba float64, RC float64, BABIP float64) player.BatterGrades {
	return player.BatterGrades{
		PlayerID:               PlayerID,
		Year:                   Year,
		TeamID:                 TeamID,
		Team:                   Team,
		Games:                  Games,
		PlateAppearance:        PlateAppearance,
		AtBat:                  AtBat,
		Score:                  Score,
		Hit:                    Hit,
		Single:                 Single,
		Double:                 Double,
		Triple:                 Triple,
		HomeRun:                HomeRun,
		BaseHit:                BaseHit,
		RunsBattedIn:           RunsBattedIn,
		StolenBase:             StolenBase,
		CaughtStealing:         CaughtStealing,
		SacrificeHits:          SacrificeHits,
		SacrificeFlies:         SacrificeFlies,
		BaseOnBalls:            BaseOnBalls,
		HitByPitches:           HitByPitches,
		StrikeOut:              StrikeOut,
		StrikeOutRate:          StrikeOutRate,
		GroundedIntoDoublePlay: GroundedIntoDoublePlay,
		BattingAverage:         BattingAverage,
		SluggingPercentage:     SluggingPercentage,
		OnBasePercentage:       OnBasePercentage,
		WOba:                   Woba,
		RC:                     RC,
		Babip:                  BABIP,
	}
}

func TestGradesRepository_InsertCareers_GetCareer(t *testing.T) {
	type args struct {
		playerID string
		players  []player.Players
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績登録と取得",
			args{
				"01105137",
				[]player.Players{
					{
						PlayerID:           "01105137",
						Name:               "飯田　優也",
						Position:           "投手",
						PitchingAndBatting: "左投左打",
						Height:             "187cm",
						Weight:             "92kg",
						Birthday:           "1990年11月27日",
						Career:             "神戸弘陵高 - 東京農業大生産学部",
						Draft:              "2012年育成選手ドラフト3位",
						SearchName:         "飯田優也",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			gorm := testUtil.ConnectGormDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.GormDB = gorm
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertPlayers(tt.args.players)
			actual := repository.GetPlayers(tt.args.playerID)
			assert.Exactly(t, tt.args.players[0], actual)
		})
	}
}

func TestGradesRepository_SearchCareerByName(t *testing.T) {
	type args struct {
		name    string
		players []player.Players
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手名検索",
			args{
				"飯田",
				[]player.Players{
					{
						PlayerID:           "01105137",
						Name:               "飯田　優也",
						Position:           "投手",
						PitchingAndBatting: "左投左打",
						Height:             "187cm",
						Weight:             "92kg",
						Birthday:           "1990年11月27日",
						Career:             "神戸弘陵高 - 東京農業大生産学部",
						Draft:              "2012年育成選手ドラフト3位",
						SearchName:         "飯田優也",
					},
				},
			},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectGormDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.GormDB = db
	repository := GradesRepository{SQLHandler: *sqlHandler}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository.InsertPlayers(tt.args.players)
			actual := repository.SearchCareerByName(tt.args.name)
			assert.ElementsMatch(t, tt.args.players, actual)
		})
	}
}
