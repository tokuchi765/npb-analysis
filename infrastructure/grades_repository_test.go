package infrastructure

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/entity/sqlwrapper"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestGradesRepository_InsertPicherGrades_GetPitchings(t *testing.T) {
	type args struct {
		playerID string
		pitcher  data.PICHERGRADES
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"投手成績登録と取得",
			args{
				"53355134",
				createPicherGrades("2020", "01", "チーム名", 54.0, 4.0, 2.0, 1.0, 32.0, 36.0, 2.0, 3.0, 1.0, 0.667, 213.0, 53.0, 40.0, 4.0, 16.0, 2.0, 46.0, 2.0, 10.0, 19.0, 17.0, 2.89, 0.3, 3.6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertPicherGrades(tt.args.playerID, tt.args.pitcher)
			actual := repository.GetPitchings(tt.args.playerID)
			assert.ElementsMatch(t, []data.PICHERGRADES{tt.args.pitcher}, actual)
		})
	}
}

func createPicherGradesList() []data.PICHERGRADES {
	return []data.PICHERGRADES{
		createPicherGrades("2020", "01", "チーム名", 54.0, 4.0, 2.0, 1.0, 32.0, 36.0, 2.0, 3.0, 1.0, 0.667, 213.0, 53.0, 40.0, 4.0, 16.0, 2.0, 46.0, 2.0, 10.0, 19.0, 17.0, 2.89, 0.3, 3.6),
	}
}

func createPicherGrades(year string, teamID string, team string, piched float64, win float64, lose float64, save float64, hold float64, holdPoint float64, completeGame float64, shutout float64, noWalks float64, winningRate float64, batter float64, inningsPitched float64, hit float64, homeRun float64, baseOnBalls float64, hitByPitches float64, strikeOut float64, wildPitches float64, balk float64, runsAllowed float64, earnedRun float64, earnedRunAverage float64, babip float64, strikeOutRate float64) data.PICHERGRADES {
	return data.PICHERGRADES{
		Year:             year,
		TeamID:           teamID,
		Team:             team,
		Piched:           piched,
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
		BABIP:            babip,
		StrikeOutRate:    strikeOutRate,
	}
}

func TestGradesRepository_InsertBatterGrades_GetBattings(t *testing.T) {
	type args struct {
		playerID string
		batting  data.BATTERGRADES
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"打者成績登録と取得",
			args{
				"01605136",
				createBatterGrades("2018", "12", "オリックス", 113, 345, 295, 39, 78, 0, 8, 4, 1, 97, 15, 16, 9, 16, 0, 31, 3, 33, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertBatterGrades(tt.args.playerID, tt.args.batting)
			actual := repository.GetBattings(tt.args.playerID)
			assert.Equal(t, []data.BATTERGRADES{tt.args.batting}, actual)
		})
	}
}

func createBatterGradesList() []data.BATTERGRADES {
	return []data.BATTERGRADES{
		createBatterGrades("2018", "12", "オリックス", 113, 345, 295, 39, 78, 0, 8, 4, 1, 97, 15, 16, 9, 16, 0, 31, 3, 33, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3),
	}
}

func createBatterGrades(Year string, TeamID string, Team string, Games int, PlateAppearance int, AtBat int, Score int, Hit int, Single int, Double int, Triple int, HomeRun int, BaseHit int, RunsBattedIn int, StolenBase int, CaughtStealing int, SacrificeHits int, SacrificeFlies int, BaseOnBalls int, HitByPitches int, StrikeOut int, StrikeOutRate float64, GroundedIntoDoublePlay int, BattingAverage float64, SluggingPercentage float64, OnBasePercentage float64, Woba float64, RC float64, BABIP float64) data.BATTERGRADES {
	return data.BATTERGRADES{
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
		StrikeOutRate:          sqlwrapper.NullFloat64{NullFloat64: sql.NullFloat64{Float64: StrikeOutRate, Valid: true}},
		GroundedIntoDoublePlay: GroundedIntoDoublePlay,
		BattingAverage:         BattingAverage,
		SluggingPercentage:     SluggingPercentage,
		OnBasePercentage:       OnBasePercentage,
		Woba:                   Woba,
		RC:                     RC,
		BABIP:                  BABIP,
	}
}

func TestGradesRepository_InsertCareers_GetCareer(t *testing.T) {
	type args struct {
		playerID string
		careers  []data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績登録と取得",
			args{
				"01105137",
				[]data.CAREER{
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
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertCareers(tt.args.careers)
			actual := repository.GetCareer(tt.args.playerID)
			assert.Exactly(t, tt.args.careers[0], actual)
		})
	}
}

func TestGradesRepository_GetPlayersByTeamIDAndYear(t *testing.T) {
	type args struct {
		teamID   string
		teamName string
		year     string
	}
	tests := []struct {
		name        string
		args        args
		wantPlayers []data.PLAYER
	}{
		{
			"選手一覧取得",
			args{
				"01",
				"Giants",
				"2020",
			},
			[]data.PLAYER{
				{
					Year:     "2020",
					TeamID:   "01",
					PlayerID: "93795138",
					Team:     "Giants",
					Name:     "デラロサ",
				},
				{
					Year:     "2020",
					TeamID:   "01",
					PlayerID: "41045138",
					Team:     "Giants",
					Name:     "戸郷　翔征",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			players := [][]string{
				{"93795138", "デラロサ"},
				{"41045138", "戸郷　翔征"},
			}
			repository.InsertTeamPlayers(tt.args.teamID, tt.args.teamName, players, tt.args.year)
			actual := repository.GetPlayersByTeamIDAndYear(tt.args.teamID, tt.args.year)
			assert.ElementsMatch(t, tt.wantPlayers, actual)
		})
	}
}

func TestGradesRepository_ExtractionCareers(t *testing.T) {
	type args struct {
		careers []data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"重複Careerを削除",
			args{
				[]data.CAREER{
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
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertCareers(tt.args.careers)
			repository.ExtractionCareers(&tt.args.careers)
			assert.Empty(t, tt.args.careers)
		})
	}
}

func TestGradesRepository_ExtractionPicherGrades(t *testing.T) {
	type args struct {
		picherMap map[string][]data.PICHERGRADES
		teamID    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"重複投手成績を削除する",
			args{
				map[string][]data.PICHERGRADES{
					"53355134": createPicherGradesList(),
				},
				"01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertPicherGrades("53355134", createPicherGrades("2020", "01", "チーム名", 54.0, 4.0, 2.0, 1.0, 32.0, 36.0, 2.0, 3.0, 1.0, 0.667, 213.0, 53.0, 40.0, 4.0, 16.0, 2.0, 46.0, 2.0, 10.0, 19.0, 17.0, 2.89, 0.3, 3.6))
			repository.ExtractionPicherGrades(&tt.args.picherMap, tt.args.teamID)
			assert.Empty(t, tt.args.picherMap)
		})
	}
}

func TestGradesRepository_ExtractionBatterGrades(t *testing.T) {
	batter := createBatterGrades("2018", "12", "オリックス", 113, 345, 295, 39, 78, 0, 8, 4, 1, 97, 15, 16, 9, 16, 0, 31, 3, 33, 0.3, 2, 0.264, 0.328, 0.34, 0.351, 60.2, 0.3)
	prayerID := "01605136"
	type args struct {
		prayerID string
		batter   data.BATTERGRADES
		teamID   string
	}
	tests := []struct {
		name      string
		args      args
		batterMap map[string][]data.BATTERGRADES
	}{
		{
			"重複打撃成績を削除する",
			args{
				prayerID,
				batter,
				"12",
			},
			map[string][]data.BATTERGRADES{
				prayerID: {batter},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := GradesRepository{SQLHandler: *sqlHandler}
			repository.InsertBatterGrades(tt.args.prayerID, tt.args.batter)
			repository.ExtractionBatterGrades(&tt.batterMap, tt.args.teamID)
			assert.Empty(t, tt.batterMap)
		})
	}
}

func TestGradesRepository_SearchCareerByName(t *testing.T) {
	type args struct {
		name    string
		careers []data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手名検索",
			args{
				"飯田",
				[]data.CAREER{
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
					},
				},
			},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	repository := GradesRepository{SQLHandler: *sqlHandler}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository.InsertCareers(tt.args.careers)
			actual := repository.SearchCareerByName(tt.args.name)
			assert.ElementsMatch(t, tt.args.careers, actual)
		})
	}
}
