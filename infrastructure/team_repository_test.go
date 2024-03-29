package infrastructure

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/sqlwrapper"
	teamData "github.com/tokuchi765/npb-analysis/entity/team"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestTeamRepository_InsertTeamPitchings_GetTeamPitchings(t *testing.T) {
	type args struct {
		years        []int
		teamPitching teamData.TeamPitching
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チーム投手成績取得と登録",
			args{
				[]int{2020},
				createTeamPitching("01", "2020", 3.4, 143, 60, 60, 60, 60, 60, 60, 60, 60, 3.4, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 0.3, 3.6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamPitchings(tt.args.teamPitching)
			actual := repository.GetTeamPitchings(tt.args.years)

			assert.Exactly(t, tt.args.teamPitching, actual["2020"][0])
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func createTeamPitching(teamID string, year string, earnedRunAverage float64, games int, win int, lose int, save int, hold int, holdPoint int, completeGame int, shutout int, noWalks int, winningRate float64, batter int, inningsPitched int, hit int, homeRun int, baseOnBalls int, intentionalWalk int, hitByPitches int, strikeOut int, wildPitches int, balk int, runsAllowed int, earnedRun int, babip float64, strikeOutRate float64) (teamPitching teamData.TeamPitching) {
	return teamData.TeamPitching{
		TeamID:           teamID,
		Year:             year,
		EarnedRunAverage: earnedRunAverage,
		Games:            games,
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
		IntentionalWalk:  intentionalWalk,
		HitByPitches:     hitByPitches,
		StrikeOut:        strikeOut,
		WildPitches:      wildPitches,
		Balk:             balk,
		RunsAllowed:      runsAllowed,
		EarnedRun:        earnedRun,
		BABIP:            babip,
		StrikeOutRate:    strikeOutRate,
	}
}

func TestTeamInteractor_InsertTeamBattings_GetTeamBatting(t *testing.T) {
	type args struct {
		teamBatting teamData.TeamBatting
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チーム打撃成績取得と登録",
			args{
				teamBatting: createTeamBatting("01", "2005", 0.301, 144, 360, 360, 400, 360, 90, 5, 70, 400, 400, 50, 20, 20, 20, 100, 100, 100, 100, 0.3, 20, 0.21, 0.314, 0.3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamBattings(tt.args.teamBatting)

			batting := repository.GetTeamBattings([]int{2005})["2005"]

			assert.Equal(t, tt.args.teamBatting, batting[0])

			testUtil.CloseContainer(resource, pool)
		})
	}
}

func createTeamBatting(teamID string, year string, battingAverage float64, games int, plateAppearance int, atBat int, score int, hit int, double int, triple int, homeRun int, baseHit int, runsBattedIn int, stolenBase int, caughtStealing int, sacrificeHits int, sacrificeFlies int, baseOnBalls int, intentionalWalk int, hitByPitches int, strikeOut int, strikeOutRate float64, groundedIntoDoublePlay int, sluggingPercentage float64, onBasePercentage float64, babip float64) teamData.TeamBatting {
	return teamData.TeamBatting{
		TeamID:                 teamID,
		Year:                   year,
		BattingAverage:         battingAverage,
		Games:                  games,
		PlateAppearance:        plateAppearance,
		AtBat:                  atBat,
		Score:                  score,
		Hit:                    hit,
		Double:                 double,
		Triple:                 triple,
		HomeRun:                homeRun,
		BaseHit:                baseHit,
		RunsBattedIn:           runsBattedIn,
		StolenBase:             stolenBase,
		CaughtStealing:         caughtStealing,
		SacrificeHits:          sacrificeHits,
		SacrificeFlies:         sacrificeFlies,
		BaseOnBalls:            baseOnBalls,
		IntentionalWalk:        intentionalWalk,
		HitByPitches:           hitByPitches,
		StrikeOut:              strikeOut,
		StrikeOutRate:          sqlwrapper.NullFloat64{NullFloat64: sql.NullFloat64{Float64: strikeOutRate, Valid: true}},
		GroundedIntoDoublePlay: groundedIntoDoublePlay,
		SluggingPercentage:     sluggingPercentage,
		OnBasePercentage:       onBasePercentage,
		BABIP:                  babip,
	}
}

func TestTeamRepository_GetTeamStats(t *testing.T) {
	type args struct {
		teamBattings  []teamData.TeamBatting
		teamPitchings []teamData.TeamPitching
		expected      float64
		teamID        string
		year          string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チームリーグ成績取得",
			args{
				teamBattings: []teamData.TeamBatting{
					{TeamID: "01", Year: "2020", Score: 100},
				},
				teamPitchings: []teamData.TeamPitching{
					{TeamID: "01", Year: "2020", RunsAllowed: 100},
				},
				expected: 0.5,
				teamID:   "01",
				year:     "2020",
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
			repository := TeamRepository{SQLHandler: *sqlHandler}

			insertDefaultTeamStats(tt.args.teamID, tt.args.year, db)
			repository.InsertPythagoreanExpectation(tt.args.teamBattings, tt.args.teamPitchings)
			actual := repository.GetTeamStats([]int{2020})
			assert.Exactly(t, tt.args.expected, actual["2020"][0].PythagoreanExpectation)
		})
	}
}

func insertDefaultTeamStats(teamID string, year string, db *sql.DB) {
	stmt1, _ := db.Prepare("INSERT INTO team_season_stats(team_id, year, manager, games, win, lose, draw, winning_rate, exchange_win, exchange_lose, exchange_draw, home_win, home_lose, home_draw, load_win, load_lose, load_draw, pythagorean_expectation) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)")
	stmt1.Exec(teamID, year, "manager", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.0)
	stmt1.Close()
}

func Test_calcPythagoreanExpectation(t *testing.T) {
	type args struct {
		score       int
		runsAllowed int
		want        float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ピタゴラス勝率を計算",
			args: args{score: 100, runsAllowed: 100, want: 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.args.want, calcPythagoreanExpectation(tt.args.score, tt.args.runsAllowed))
		})
	}
}

func TestTeamRepository_InsertPythagoreanExpectation(t *testing.T) {
	type args struct {
		years         []int
		teamBattings  []teamData.TeamBatting
		teamPitchings []teamData.TeamPitching
		expected      float64
		teamID        string
		year          string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"ピタゴラス勝率登録",
			args{
				years: []int{2020},
				teamBattings: []teamData.TeamBatting{
					{TeamID: "01", Year: "2020", Score: 100},
				},
				teamPitchings: []teamData.TeamPitching{
					{TeamID: "01", Year: "2020", RunsAllowed: 100},
				},
				expected: 0.5,
				teamID:   "01",
				year:     "2020",
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
			repository := TeamRepository{SQLHandler: *sqlHandler}

			insertDefaultTeamStats(tt.args.teamID, tt.args.year, db)
			repository.InsertPythagoreanExpectation(tt.args.teamBattings, tt.args.teamPitchings)
			actual := repository.GetTeamStats([]int{2020})
			assert.Exactly(t, tt.args.expected, actual["2020"][0].PythagoreanExpectation)
		})
	}
}

func TestTeamRepository_InsertTeamLeagueStats(t *testing.T) {
	type args struct {
		teamID          string
		year            string
		expectedManager string
		expectedGames   int
		expectedWin     int
		expectedLose    int
		expectedDraw    int
		teamLeagueStats []teamData.TeamLeagueStats
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チームシーズン成績登録",
			args{
				teamID:          "01",
				year:            "2020",
				expectedManager: "監督",
				expectedGames:   146,
				expectedWin:     60,
				expectedLose:    40,
				expectedDraw:    46,
				teamLeagueStats: []teamData.TeamLeagueStats{
					{
						TeamID:                 "01",
						Year:                   "2020",
						Manager:                "監督",
						Games:                  146,
						Win:                    60,
						Lose:                   40,
						Draw:                   46,
						WinningRate:            0.6,
						ExchangeWin:            10,
						ExchangeLose:           10,
						ExchangeDraw:           5,
						HomeWin:                30,
						HomeLose:               20,
						HomeDraw:               20,
						LoadWin:                30,
						LoadLose:               20,
						LoadDraw:               26,
						PythagoreanExpectation: 0.6,
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
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamLeagueStats(tt.args.teamLeagueStats)
			rows, _ := db.Query("SELECT manager,games,win,lose,draw FROM team_season_stats WHERE team_id = $1 AND year = $2", tt.args.teamID, tt.args.year)
			var manager string
			var games, win, lose, draw int
			for rows.Next() {
				rows.Scan(&manager, &games, &win, &lose, &draw)
			}
			rows.Close()
			assert.Equal(t, tt.args.expectedManager, manager)
			assert.Equal(t, tt.args.expectedGames, games)
			assert.Equal(t, tt.args.expectedWin, win)
			assert.Equal(t, tt.args.expectedLose, lose)
			assert.Equal(t, tt.args.expectedDraw, draw)
		})
	}
}

func TestTeamRepository_InsertMatchResults(t *testing.T) {
	type expectedData struct {
		expectedVsType string
		expectedWin    int
		expectedLose   int
		expectedDraw   int
	}
	type args struct {
		teamID         string
		year           string
		opponentTeamID string
		expectedVsType string
		league         expectedData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"リーグ対戦成績登録",
			args{
				teamID:         "01",
				year:           "2020",
				opponentTeamID: "06",
				league: expectedData{
					expectedVsType: "league",
					expectedWin:    15,
					expectedLose:   6,
					expectedDraw:   3,
				},
			},
		},
		{
			"交流戦対戦成績登録",
			args{
				teamID:         "01",
				year:           "2005",
				opponentTeamID: "12",
				league: expectedData{
					expectedVsType: "exchange",
					expectedWin:    2,
					expectedLose:   2,
					expectedDraw:   2,
				},
			},
		},
		{
			"交流戦対戦成績未登録年",
			args{
				teamID:         "01",
				year:           "2020",
				opponentTeamID: "12",
				league: expectedData{
					expectedVsType: "",
					expectedWin:    0,
					expectedLose:   0,
					expectedDraw:   0,
				},
			},
		},
	}
	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	repository := TeamRepository{SQLHandler: *sqlHandler}

	teamMatchResults := []teamData.TeamMatchResults{
		{
			TeamID:            "01",
			Year:              "2020",
			CompetitiveTeamID: "06",
			VsType:            "league",
			Win:               15,
			Lose:              6,
			Draw:              3,
		},
		{
			TeamID:            "01",
			Year:              "2005",
			CompetitiveTeamID: "12",
			VsType:            "exchange",
			Win:               2,
			Lose:              2,
			Draw:              2,
		},
	}
	repository.InsertMatchResults(teamMatchResults)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows, _ := db.Query("SELECT vs_type,win,lose,draw FROM team_match_results WHERE team_id = $1 AND year = $2 AND competitive_team_id = $3", tt.args.teamID, tt.args.year, tt.args.opponentTeamID)
			var vsType string
			var win, lose, draw int
			for rows.Next() {
				rows.Scan(&vsType, &win, &lose, &draw)
			}
			rows.Close()
			assert.Equal(t, tt.args.league.expectedVsType, vsType)
			assert.Equal(t, tt.args.league.expectedWin, win)
			assert.Equal(t, tt.args.league.expectedLose, lose)
			assert.Equal(t, tt.args.league.expectedDraw, draw)
		})
	}
}

func TestTeamRepository_GetTeamPitchingByTeamIDAndYear(t *testing.T) {
	type args struct {
		year   string
		teamID string
	}
	tests := []struct {
		name             string
		args             args
		wantTeamPitching teamData.TeamPitching
	}{
		{
			"チーム投手成績取得（チームIDと年指定）",
			args{
				"2020",
				"01",
			},
			createTeamPitching("01", "2020", 3.4, 143, 60, 60, 60, 60, 60, 60, 60, 60, 3.4, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 0.3, 3.6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamPitchings(tt.wantTeamPitching)
			actual := repository.GetTeamPitchingByTeamIDAndYear(tt.args.year, tt.args.teamID)
			assert.Exactly(t, tt.wantTeamPitching, actual)
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestTeamRepository_GetTeamBattingByTeamIDAndYear(t *testing.T) {
	type args struct {
		teamID string
		year   string
	}
	tests := []struct {
		name            string
		args            args
		wantTeamBatting teamData.TeamBatting
	}{
		{
			"チーム打撃成績取得（チームIDと年指定）",
			args{
				"01",
				"2005",
			},
			createTeamBatting("01", "2005", 0.301, 144, 360, 360, 400, 360, 90, 5, 70, 400, 400, 50, 20, 20, 20, 100, 100, 100, 100, 0.3, 20, 0.21, 0.314, 0.3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamBattings(tt.wantTeamBatting)
			actual := repository.GetTeamBattingByTeamIDAndYear(tt.args.teamID, tt.args.year)
			assert.Exactly(t, tt.wantTeamBatting, actual)
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestTeamRepository_GetTeamBattingMax(t *testing.T) {
	tests := []struct {
		name                      string
		wantMaxHomeRun            int
		wantMaxSluggingPercentage float64
		wantMaxOnBasePercentage   float64
	}{
		{
			"打撃成績最大値取得",
			90,
			0.67,
			0.531,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamBattings(createTeamBatting("01", "2005", 0, 0, 0, 0, 0, 0, 0, 0, 70, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.67, 0.451, 0))
			repository.InsertTeamBattings(createTeamBatting("05", "2007", 0, 0, 0, 0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.21, 0.314, 0))
			repository.InsertTeamBattings(createTeamBatting("11", "2015", 0, 0, 0, 0, 0, 0, 0, 0, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.30, 0.531, 0))

			maxHomeRun, maxSluggingPercentage, maxOnBasePercentage := repository.GetTeamBattingMax()
			assert.Equal(t, tt.wantMaxHomeRun, maxHomeRun)
			assert.Equal(t, tt.wantMaxSluggingPercentage, maxSluggingPercentage)
			assert.Equal(t, tt.wantMaxOnBasePercentage, maxOnBasePercentage)
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestTeamRepository_GetTeamBattingMin(t *testing.T) {
	tests := []struct {
		name                      string
		wantMinHomeRun            int
		wantMinSluggingPercentage float64
		wantMinOnBasePercentage   float64
	}{
		{
			"打撃成績最小値取得",
			70,
			0.21,
			0.314,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamBattings(createTeamBatting("01", "2005", 0, 0, 0, 0, 0, 0, 0, 0, 70, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.67, 0.451, 0))
			repository.InsertTeamBattings(createTeamBatting("05", "2007", 0, 0, 0, 0, 0, 0, 0, 0, 90, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.21, 0.314, 0))
			repository.InsertTeamBattings(createTeamBatting("11", "2015", 0, 0, 0, 0, 0, 0, 0, 0, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.30, 0.531, 0))

			minHomeRun, minSluggingPercentage, minOnBasePercentage := repository.GetTeamBattingMin()

			assert.Equal(t, tt.wantMinHomeRun, minHomeRun)
			assert.Equal(t, tt.wantMinSluggingPercentage, minSluggingPercentage)
			assert.Equal(t, tt.wantMinOnBasePercentage, minOnBasePercentage)
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestTeamRepository_GetTeamPitchingMax(t *testing.T) {
	tests := []struct {
		name                 string
		wantMaxStrikeOutRate float64
		wantMaxRunsAllowed   int
	}{
		{
			"投手成績最大値取得",
			3.6,
			80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamPitchings(createTeamPitching("01", "2020", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 3.6))
			repository.InsertTeamPitchings(createTeamPitching("06", "2009", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 80, 0, 0, 2.3))

			maxStrikeOutRate, maxRunsAllowed := repository.GetTeamPitchingMax()
			assert.Equal(t, tt.wantMaxStrikeOutRate, maxStrikeOutRate)
			assert.Equal(t, tt.wantMaxRunsAllowed, maxRunsAllowed)
			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestTeamRepository_GetTeamPitchingMin(t *testing.T) {
	tests := []struct {
		name                 string
		wantMinStrikeOutRate float64
		wantMinRunsAllowed   int
	}{
		{
			"投手成績最小値取得",
			2.3,
			60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(SQLHandler)
			sqlHandler.Conn = db
			repository := TeamRepository{SQLHandler: *sqlHandler}

			repository.InsertTeamPitchings(createTeamPitching("01", "2020", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 3.6))
			repository.InsertTeamPitchings(createTeamPitching("06", "2009", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 80, 0, 0, 2.3))

			minStrikeOutRate, minRunsAllowed := repository.GetTeamPitchingMin()

			assert.Equal(t, tt.wantMinStrikeOutRate, minStrikeOutRate)
			assert.Equal(t, tt.wantMinRunsAllowed, minRunsAllowed)
			testUtil.CloseContainer(resource, pool)
		})
	}
}
