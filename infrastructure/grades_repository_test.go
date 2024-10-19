package infrastructure

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/entity/sqlwrapper"
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

func TestGradesRepository_SearchBatterGrades(t *testing.T) {
	test1 := createBatterGradesResult("00001", "テスト1", "2018", "オリックス", 100, 400, 380, 40, 120, 60, 20, 10, 30, 20, 15, 30, 10, 10, 5, 10, 3, 15, 0.2, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3)
	test1_total := createBatterGradesResult("00001", "テスト1", "nan", "", 100, 400, 380, 40, 120, 60, 20, 10, 30, 20, 15, 30, 10, 10, 5, 10, 3, 30, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3)
	test2 := createBatterGradesResult("00002", "テスト2", "2019", "北海道日本ハム", 120, 500, 480, 40, 160, 70, 40, 20, 30, 20, 15, 40, 10, 10, 5, 15, 3, 20, 0.27, 2, 0.281, 0.369, 0.41, 0.382, 70.2, 0.34)
	test2_total := createBatterGradesResult("00002", "テスト2", "nan", "", 120, 500, 480, 40, 160, 70, 40, 20, 30, 20, 15, 40, 10, 10, 5, 15, 3, 30, 0.3, 2, 0.281, 0.369, 0.41, 0.382, 70.2, 0.34)
	test3 := createBatterGradesResult("00003", "テスト3", "2020", "北海道日本ハム", 200, 600, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 6, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43)
	test3_2 := createBatterGradesResult("00003", "テスト3", "2021", "北海道日本ハム", 200, 610, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 6, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43)
	test3_total := createBatterGradesResult("00003", "テスト3", "nan", "", 200, 600, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 2, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43)

	tests := []struct {
		name        string
		condition   player.SearchBatterGradesCondition
		wantResults []player.SearchBatterGradesResult
	}{
		{
			"条件なしによる検索ができない",
			player.SearchBatterGradesCondition{},
			[]player.SearchBatterGradesResult{},
		},
		{
			"通年検索が可能",
			player.SearchBatterGradesCondition{
				Total:     true,
				TotalYear: 2,
			},
			[]player.SearchBatterGradesResult{test3_total},
		},
		{
			"期間による検索が可能",
			player.SearchBatterGradesCondition{
				FromYear: 2018,
				ToYear:   2018,
			},
			[]player.SearchBatterGradesResult{test1},
		},
		{
			"チームIDによる検索が可能",
			player.SearchBatterGradesCondition{
				TeamID: "12",
			},
			[]player.SearchBatterGradesResult{test1},
		},
		{
			"打席数、以上による検索が可能",
			player.SearchBatterGradesCondition{
				PlateAppearanceThresholdType: player.GreaterOrEqual,
				PlateAppearance:              610,
			},
			[]player.SearchBatterGradesResult{test3_2},
		},
		{
			"打席数、以下による検索が可能",
			player.SearchBatterGradesCondition{
				PlateAppearanceThresholdType: player.LessOrEqual,
				PlateAppearance:              500,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"出塁率、以上による検索が可能",
			player.SearchBatterGradesCondition{
				OnBasePercentageThresholdType: player.GreaterOrEqual,
				OnBasePercentage:              0.438,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"出塁率、以下による検索が可能",
			player.SearchBatterGradesCondition{
				OnBasePercentageThresholdType: player.LessOrEqual,
				OnBasePercentage:              0.437,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"打率、以上による検索が可能",
			player.SearchBatterGradesCondition{
				BattingAverageThresholdType: player.GreaterOrEqual,
				BattingAverage:              0.33,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"打率、以下による検索が可能",
			player.SearchBatterGradesCondition{
				BattingAverageThresholdType: player.LessOrEqual,
				BattingAverage:              0.281,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"長打率、以上による検索が可能",
			player.SearchBatterGradesCondition{
				SluggingPercentageThresholdType: player.GreaterOrEqual,
				SluggingPercentage:              0.369,
			},
			[]player.SearchBatterGradesResult{test2, test3, test3_2},
		},
		{
			"長打率、以下による検索が可能",
			player.SearchBatterGradesCondition{
				SluggingPercentageThresholdType: player.LessOrEqual,
				SluggingPercentage:              0.368,
			},
			[]player.SearchBatterGradesResult{test1},
		},
		{
			"本塁打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				HomeRunThresholdType: player.GreaterOrEqual,
				HomeRun:              40,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"本塁打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				HomeRunThresholdType: player.LessOrEqual,
				HomeRun:              39,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"四球、以上による検索が可能",
			player.SearchBatterGradesCondition{
				BaseOnBallsThresholdType: player.GreaterOrEqual,
				BaseOnBalls:              30,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"四球、以下による検索が可能",
			player.SearchBatterGradesCondition{
				BaseOnBallsThresholdType: player.LessOrEqual,
				BaseOnBalls:              15,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"安打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				HitThresholdType: player.GreaterOrEqual,
				Hit:              200,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"安打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				HitThresholdType: player.LessOrEqual,
				Hit:              160,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"単打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				SingleThresholdType: player.GreaterOrEqual,
				Single:              90,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"単打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				SingleThresholdType: player.LessOrEqual,
				Single:              70,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"二塁打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				DoubleThresholdType: player.GreaterOrEqual,
				Double:              50,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"二塁打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				DoubleThresholdType: player.LessOrEqual,
				Double:              40,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"三塁打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				TripleThresholdType: player.GreaterOrEqual,
				Triple:              30,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"三塁打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				TripleThresholdType: player.LessOrEqual,
				Triple:              20,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"三振、以上による検索が可能",
			player.SearchBatterGradesCondition{
				StrikeOutThresholdType: player.GreaterOrEqual,
				StrikeOut:              30,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"三振、以下による検索が可能",
			player.SearchBatterGradesCondition{
				StrikeOutThresholdType: player.LessOrEqual,
				StrikeOut:              29,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"三振率、以上による検索が可能",
			player.SearchBatterGradesCondition{
				StrikeOutRateThresholdType: player.GreaterOrEqual,
				StrikeOutRate:              0.3,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"三振率、以下による検索が可能",
			player.SearchBatterGradesCondition{
				StrikeOutRateThresholdType: player.LessOrEqual,
				StrikeOutRate:              0.29,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"盗塁、以上による検索が可能",
			player.SearchBatterGradesCondition{
				StolenBaseThresholdType: player.GreaterOrEqual,
				StolenBase:              50,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"盗塁、以下による検索が可能",
			player.SearchBatterGradesCondition{
				StolenBaseThresholdType: player.LessOrEqual,
				StolenBase:              40,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"併殺打、以上による検索が可能",
			player.SearchBatterGradesCondition{
				GroundedIntoDoublePlayThresholdType: player.GreaterOrEqual,
				GroundedIntoDoublePlay:              6,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"併殺打、以下による検索が可能",
			player.SearchBatterGradesCondition{
				GroundedIntoDoublePlayThresholdType: player.LessOrEqual,
				GroundedIntoDoublePlay:              5,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"加重出塁率、以上による検索が可能",
			player.SearchBatterGradesCondition{
				WObaThresholdType: player.GreaterOrEqual,
				WOba:              0.438,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"加重出塁率、以下による検索が可能",
			player.SearchBatterGradesCondition{
				WObaThresholdType: player.LessOrEqual,
				WOba:              0.437,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"創出得点、以上による検索が可能",
			player.SearchBatterGradesCondition{
				RCThresholdType: player.GreaterOrEqual,
				RC:              81.2,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"創出得点、以下による検索が可能",
			player.SearchBatterGradesCondition{
				RCThresholdType: player.LessOrEqual,
				RC:              81.1,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"BABIP、以上による検索が可能",
			player.SearchBatterGradesCondition{
				BabipThresholdType: player.GreaterOrEqual,
				Babip:              0.43,
			},
			[]player.SearchBatterGradesResult{test3, test3_2},
		},
		{
			"BABIP、以下による検索が可能",
			player.SearchBatterGradesCondition{
				BabipThresholdType: player.LessOrEqual,
				Babip:              0.42,
			},
			[]player.SearchBatterGradesResult{test1, test2},
		},
		{
			"通算と他の条件で、以上による検索が可能",
			player.SearchBatterGradesCondition{
				Total:              true,
				TotalYear:          1,
				BabipThresholdType: player.GreaterOrEqual,
				Babip:              0.34,
			},
			[]player.SearchBatterGradesResult{test2_total, test3_total},
		},
		{
			"通算と他の条件で、以下による検索が可能",
			player.SearchBatterGradesCondition{
				Total:              true,
				TotalYear:          1,
				BabipThresholdType: player.LessOrEqual,
				Babip:              0.33,
			},
			[]player.SearchBatterGradesResult{test1_total},
		},
		{
			"期間と他の条件で、以上による検索が可能",
			player.SearchBatterGradesCondition{
				FromYear:           2019,
				ToYear:             2020,
				BabipThresholdType: player.GreaterOrEqual,
				Babip:              0.35,
			},
			[]player.SearchBatterGradesResult{test3},
		},
		{
			"期間と他の条件で、以下による検索が可能",
			player.SearchBatterGradesCondition{
				FromYear:           2019,
				ToYear:             2020,
				BabipThresholdType: player.LessOrEqual,
				Babip:              0.35,
			},
			[]player.SearchBatterGradesResult{test2},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectGormDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.GormDB = db
	repository := GradesRepository{SQLHandler: *sqlHandler}
	repository.InsertBatterGrades(createBatterGradesMapping("00001", "2018", "12", "オリックス", 100, 400, 380, 40, 120, 60, 20, 10, 30, 20, 15, 30, 10, 10, 5, 10, 3, 15, 0.2, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3))
	repository.InsertBatterGrades(createBatterGradesMapping("00001", "nan", "13", "", 100, 400, 380, 40, 120, 60, 20, 10, 30, 20, 15, 30, 10, 10, 5, 10, 3, 30, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3))
	repository.InsertBatterGrades(createBatterGradesMapping("00002", "2019", "11", "北海道日本ハム", 120, 500, 480, 40, 160, 70, 40, 20, 30, 20, 15, 40, 10, 10, 5, 15, 3, 20, 0.27, 2, 0.281, 0.369, 0.41, 0.382, 70.2, 0.34))
	repository.InsertBatterGrades(createBatterGradesMapping("00002", "nan", "13", "", 120, 500, 480, 40, 160, 70, 40, 20, 30, 20, 15, 40, 10, 10, 5, 15, 3, 30, 0.3, 2, 0.281, 0.369, 0.41, 0.382, 70.2, 0.34))
	repository.InsertBatterGrades(createBatterGradesMapping("00003", "2020", "11", "北海道日本ハム", 200, 600, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 6, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43))
	repository.InsertBatterGrades(createBatterGradesMapping("00003", "2021", "11", "北海道日本ハム", 200, 610, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 6, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43))
	repository.InsertBatterGrades(createBatterGradesMapping("00003", "nan", "13", "", 200, 600, 550, 50, 200, 90, 50, 30, 40, 30, 15, 50, 10, 10, 5, 30, 3, 30, 0.3, 2, 0.330, 0.400, 0.52, 0.438, 81.2, 0.43))
	players := []player.Players{
		{PlayerID: "00001", Name: "テスト1"},
		{PlayerID: "00002", Name: "テスト2"},
		{PlayerID: "00003", Name: "テスト3"},
	}
	repository.InsertPlayers(players)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := repository.SearchBatterGrades(tt.condition)
			assert.ElementsMatch(t, tt.wantResults, actual)
		})
	}
}

func createBatterGradesResult(PlayerID string, Name string, Year string, Team string, Games int, PlateAppearance int, AtBat int, Score int, Hit int, Single int, Double int, Triple int, HomeRun int, BaseHit int, RunsBattedIn int, StolenBase int, CaughtStealing int, SacrificeHits int, SacrificeFlies int, BaseOnBalls int, HitByPitches int, StrikeOut int, StrikeOutRate float64, GroundedIntoDoublePlay int, BattingAverage float64, SluggingPercentage float64, OnBasePercentage float64, Woba float64, RC float64, BABIP float64) player.SearchBatterGradesResult {
	return player.SearchBatterGradesResult{
		PlayerID:               PlayerID,
		Name:                   Name,
		Year:                   Year,
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
		StrikeOutRate:          sqlwrapper.NullFloat64{NullFloat64: sql.NullFloat64{Float64: StrikeOutRate, Valid: true}},
		GroundedIntoDoublePlay: GroundedIntoDoublePlay,
		BattingAverage:         BattingAverage,
		SluggingPercentage:     SluggingPercentage,
		OnBasePercentage:       OnBasePercentage,
		WOba:                   Woba,
		RC:                     RC,
		Babip:                  BABIP,
	}
}

func TestGradesRepository_SearchPicherGrades(t *testing.T) {
	test1 := createPicherGradesResult("00001", "2018", "テスト1", "テストチーム1", 34.0, 2, 2, 1, 25, 25, 2, 3, 1.0, 0.475, 156.0, 43.0, 27, 4, 12, 2.0, 35, 2.0, 10.0, 19.0, 17.0, 1.89, 0.3, 2.3)
	test1_total := createPicherGradesResult("00001", "nan", "テスト1", "", 34.0, 2, 2, 1, 25, 25, 2, 3, 1.0, 0.475, 156.0, 43.0, 27, 4, 12, 2.0, 35, 2.0, 10.0, 19.0, 17.0, 1.89, 0.3, 2.3)
	test2 := createPicherGradesResult("00002", "2019", "テスト2", "テストチーム3", 40.0, 4, 5, 5, 32, 36, 6, 5, 3.0, 0.617, 183.0, 53.0, 40, 6, 16, 5.0, 46, 4.0, 16.0, 37.0, 23.0, 2.22, 0.5, 3.6)
	test2_total := createPicherGradesResult("00002", "nan", "テスト2", "", 40.0, 4, 5, 5, 32, 36, 6, 5, 3.0, 0.617, 183.0, 53.0, 40, 6, 16, 5.0, 46, 4.0, 16.0, 37.0, 23.0, 2.22, 0.5, 3.6)
	test3 := createPicherGradesResult("00003", "2020", "テスト3", "テストチーム3", 48.0, 6, 7, 7, 36, 40, 8, 7, 5.0, 0.667, 213.0, 56.0, 43, 8, 23, 7.0, 51, 7.0, 23.0, 43.0, 31.0, 2.89, 0.6, 4.3)
	test3_2 := createPicherGradesResult("00003", "2021", "テスト3", "テストチーム3", 54.0, 8, 9, 9, 42, 46, 9, 9, 7.0, 0.725, 233.0, 62.0, 50, 9, 27, 9.0, 57, 9.0, 26.0, 47.0, 34.0, 3.89, 0.8, 5.1)
	test3_total := createPicherGradesResult("00003", "nan", "テスト3", "", 54.0, 8, 9, 9, 42, 46, 9, 9, 7.0, 0.725, 233.0, 62.0, 50, 9.0, 27, 9, 57, 9.0, 26.0, 47.0, 34.0, 3.89, 0.8, 5.1)
	tests := []struct {
		name        string
		condition   player.SearchPitcherGradesCondition
		wantResults []player.SearchPitcherGradesResult
	}{
		{
			"条件なしによる検索ができない",
			player.SearchPitcherGradesCondition{},
			[]player.SearchPitcherGradesResult{},
		},
		{
			"通年検索が可能",
			player.SearchPitcherGradesCondition{
				Total:     true,
				TotalYear: 2,
			},
			[]player.SearchPitcherGradesResult{test3_total},
		},
		{
			"期間による検索が可能",
			player.SearchPitcherGradesCondition{
				FromYear: 2018,
				ToYear:   2018,
			},
			[]player.SearchPitcherGradesResult{test1},
		},
		{
			"チームIDによる検索が可能",
			player.SearchPitcherGradesCondition{
				TeamID: "01",
			},
			[]player.SearchPitcherGradesResult{test1},
		},
		{
			"登板数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				PitchedThresholdType: player.GreaterOrEqual,
				Pitched:              48,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"登板数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				PitchedThresholdType: player.LessOrEqual,
				Pitched:              47,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"投球回数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				InningsPitchedThresholdType: player.GreaterOrEqual,
				InningsPitched:              56,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"投球回数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				InningsPitchedThresholdType: player.LessOrEqual,
				InningsPitched:              55,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"防御率、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				EarnedRunAverageThresholdType: player.GreaterOrEqual,
				EarnedRunAverage:              2.89,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"防御率、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				EarnedRunAverageThresholdType: player.LessOrEqual,
				EarnedRunAverage:              2.88,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"被BABIP、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				BabipThresholdType: player.GreaterOrEqual,
				Babip:              0.6,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"被BABIP、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				BabipThresholdType: player.LessOrEqual,
				Babip:              0.59,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"奪三振率、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				StrikeOutRateThresholdType: player.GreaterOrEqual,
				StrikeOutRate:              4.3,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"奪三振率、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				StrikeOutRateThresholdType: player.LessOrEqual,
				StrikeOutRate:              4.2,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"奪三振数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				StrikeOutThresholdType: player.GreaterOrEqual,
				StrikeOut:              51,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"奪三振数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				StrikeOutThresholdType: player.LessOrEqual,
				StrikeOut:              50,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"被安打数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				HitThresholdType: player.GreaterOrEqual,
				Hit:              43,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"被安打数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				HitThresholdType: player.LessOrEqual,
				Hit:              42,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"四球数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				BaseOnBallsThresholdType: player.GreaterOrEqual,
				BaseOnBalls:              23,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"四球数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				BaseOnBallsThresholdType: player.LessOrEqual,
				BaseOnBalls:              22,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"被ホームラン数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				HomeRunThresholdType: player.GreaterOrEqual,
				HomeRun:              8,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"被ホームラン数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				HomeRunThresholdType: player.LessOrEqual,
				HomeRun:              7,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"勝利数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				WinThresholdType: player.GreaterOrEqual,
				Win:              6,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"勝利数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				WinThresholdType: player.LessOrEqual,
				Win:              5,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"敗北数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				LoseThresholdType: player.GreaterOrEqual,
				Lose:              7,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"敗北数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				LoseThresholdType: player.LessOrEqual,
				Lose:              6,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"セーブ数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				SaveThresholdType: player.GreaterOrEqual,
				Save:              7,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"セーブ数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				SaveThresholdType: player.LessOrEqual,
				Save:              6,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"ホールド数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				HoldThresholdType: player.GreaterOrEqual,
				Hold:              36,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"ホールド数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				HoldThresholdType: player.LessOrEqual,
				Hold:              35,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"ホールドポイント数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				HoldPointThresholdType: player.GreaterOrEqual,
				HoldPoint:              40,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"ホールドポイント数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				HoldPointThresholdType: player.LessOrEqual,
				HoldPoint:              39,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"完投数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				CompleteGameThresholdType: player.GreaterOrEqual,
				CompleteGame:              8,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"完投数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				CompleteGameThresholdType: player.LessOrEqual,
				CompleteGame:              7,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"完封数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				ShutoutThresholdType: player.GreaterOrEqual,
				Shutout:              7,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"完封数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				ShutoutThresholdType: player.LessOrEqual,
				Shutout:              6,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"勝率数、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				WinningRateThresholdType: player.GreaterOrEqual,
				WinningRate:              0.667,
			},
			[]player.SearchPitcherGradesResult{test3, test3_2},
		},
		{
			"勝率数、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				WinningRateThresholdType: player.LessOrEqual,
				WinningRate:              0.666,
			},
			[]player.SearchPitcherGradesResult{test1, test2},
		},
		{
			"通算と他の条件で、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				Total:                    true,
				TotalYear:                1,
				WinningRateThresholdType: player.GreaterOrEqual,
				WinningRate:              0.617,
			},
			[]player.SearchPitcherGradesResult{test2_total, test3_total},
		},
		{
			"通算と他の条件で、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				Total:                    true,
				TotalYear:                1,
				WinningRateThresholdType: player.LessOrEqual,
				WinningRate:              0.616,
			},
			[]player.SearchPitcherGradesResult{test1_total},
		},
		{
			"期間と他の条件で、以上による検索が可能",
			player.SearchPitcherGradesCondition{
				FromYear:                 2019,
				ToYear:                   2020,
				WinningRateThresholdType: player.GreaterOrEqual,
				WinningRate:              0.618,
			},
			[]player.SearchPitcherGradesResult{test3},
		},
		{
			"期間と他の条件で、以下による検索が可能",
			player.SearchPitcherGradesCondition{
				FromYear:                 2019,
				ToYear:                   2020,
				WinningRateThresholdType: player.LessOrEqual,
				WinningRate:              0.617,
			},
			[]player.SearchPitcherGradesResult{test2},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectGormDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.GormDB = db
	repository := GradesRepository{SQLHandler: *sqlHandler}
	repository.InsertPicherGrades(createPicherGradesMapping("00001", "2018", "01", "テストチーム1", 34.0, 2.0, 2.0, 1.0, 25.0, 25.0, 2.0, 3.0, 1.0, 0.475, 156.0, 43.0, 27.0, 4.0, 12.0, 2.0, 35.0, 2.0, 10.0, 19.0, 17.0, 1.89, 0.3, 2.3))
	repository.InsertPicherGrades(createPicherGradesMapping("00001", "nan", "13", "", 34.0, 2.0, 2.0, 1.0, 25.0, 25.0, 2.0, 3.0, 1.0, 0.475, 156.0, 43.0, 27.0, 4.0, 12.0, 2.0, 35.0, 2.0, 10.0, 19.0, 17.0, 1.89, 0.3, 2.3))
	repository.InsertPicherGrades(createPicherGradesMapping("00002", "2019", "03", "テストチーム3", 40.0, 4.0, 5.0, 5.0, 32.0, 36.0, 6.0, 5.0, 3.0, 0.617, 183.0, 53.0, 40.0, 6.0, 16.0, 5.0, 46.0, 4.0, 16.0, 37.0, 23.0, 2.22, 0.5, 3.6))
	repository.InsertPicherGrades(createPicherGradesMapping("00002", "nan", "13", "", 40.0, 4.0, 5.0, 5.0, 32.0, 36.0, 6.0, 5.0, 3.0, 0.617, 183.0, 53.0, 40.0, 6.0, 16.0, 5.0, 46.0, 4.0, 16.0, 37.0, 23.0, 2.22, 0.5, 3.6))
	repository.InsertPicherGrades(createPicherGradesMapping("00003", "2020", "03", "テストチーム3", 48.0, 6.0, 7.0, 7.0, 36.0, 40.0, 8.0, 7.0, 5.0, 0.667, 213.0, 56.0, 43.0, 8.0, 23.0, 7.0, 51.0, 7.0, 23.0, 43.0, 31.0, 2.89, 0.6, 4.3))
	repository.InsertPicherGrades(createPicherGradesMapping("00003", "2021", "03", "テストチーム3", 54.0, 8.0, 9.0, 9.0, 42.0, 46.0, 9.0, 9.0, 7.0, 0.725, 233.0, 62.0, 50.0, 9.0, 27.0, 9.0, 57.0, 9.0, 26.0, 47.0, 34.0, 3.89, 0.8, 5.1))
	repository.InsertPicherGrades(createPicherGradesMapping("00003", "nan", "13", "", 54.0, 8.0, 9.0, 9.0, 42.0, 46.0, 9.0, 9.0, 7.0, 0.725, 233.0, 62.0, 50.0, 9.0, 27.0, 9.0, 57.0, 9.0, 26.0, 47.0, 34.0, 3.89, 0.8, 5.1))
	players := []player.Players{
		{PlayerID: "00001", Name: "テスト1"},
		{PlayerID: "00002", Name: "テスト2"},
		{PlayerID: "00003", Name: "テスト3"},
	}
	repository.InsertPlayers(players)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := repository.SearchPicherGrades(tt.condition)
			assert.ElementsMatch(t, tt.wantResults, actual)
		})
	}
}

func createPicherGradesResult(playerID string, year string, name string, team string, pitched float64, win int, lose int, save int, hold int, holdPoint int, completeGame int, shutout int, noWalks float64, winningRate float64, batter float64, inningsPitched float64, hit int, homeRun int, baseOnBalls int, hitByPitches float64, strikeOut int, wildPitches float64, balk float64, runsAllowed float64, earnedRun float64, earnedRunAverage float64, babip float64, strikeOutRate float64) player.SearchPitcherGradesResult {
	return player.SearchPitcherGradesResult{
		PlayerID:         playerID,
		Name:             name,
		Year:             year,
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
