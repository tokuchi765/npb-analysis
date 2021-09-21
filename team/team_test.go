package team

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	teamData "github.com/tokuchi765/npb-analysis/entity/team"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestInsertPythagoreanExpectation(t *testing.T) {
	type args struct {
		years           []int
		teamBattingMap  map[string][]teamData.TeamBatting
		teamPitchingMap map[string][]teamData.TeamPitching
		expected        float64
		teamID          string
		year            string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"ピタゴラス勝率登録テスト",
			args{
				years: []int{2020},
				teamBattingMap: map[string][]teamData.TeamBatting{"2020": {
					{TeamID: "01", Year: "2020", Score: 100},
				}},
				teamPitchingMap: map[string][]teamData.TeamPitching{"2020": {
					{TeamID: "01", Year: "2020", RunsAllowed: 100},
				}},
				expected: 0.5,
				teamID:   "01",
				year:     "2020",
			},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertDefaultTeamStats(tt.args.teamID, tt.args.year, db)
			InsertPythagoreanExpectation(tt.args.years, tt.args.teamBattingMap, tt.args.teamPitchingMap, db)
			actual := GetTeamStats([]int{2020}, db)
			assert.Exactly(t, tt.args.expected, actual["2020"][0].PythagoreanExpectation)
		})
	}
}

func Test_insertPythagoreanExpectation(t *testing.T) {
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
			"ピタゴラス勝率登録テスト",
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

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertDefaultTeamStats(tt.args.teamID, tt.args.year, db)
			insertPythagoreanExpectation(tt.args.teamBattings, tt.args.teamPitchings, db)
			actual2 := GetTeamStats([]int{2020}, db)
			assert.Exactly(t, tt.args.expected, actual2["2020"][0].PythagoreanExpectation)
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

func TestInsertSeasonMatchResults(t *testing.T) {
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
			"リーグ対戦成績登録確認",
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
			"交流戦対戦成績登録確認",
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
	runtimeCurrent, _ := filepath.Abs("../")
	InsertSeasonMatchResults(runtimeCurrent+"/test/resource", db)
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

func TestInsertSeasonLeagueStats(t *testing.T) {
	type args struct {
		teamID          string
		year            string
		expectedManager string
		expectedGames   int
		expectedWin     int
		expectedLose    int
		expectedDraw    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チームシーズン成績登録確認",
			args{
				teamID:          "03",
				year:            "2005",
				expectedManager: "岡田　彰布",
				expectedGames:   146,
				expectedWin:     87,
				expectedLose:    54,
				expectedDraw:    5,
			},
		},
	}

	resource, pool := testUtil.CreateContainer()
	defer testUtil.CloseContainer(resource, pool)
	db := testUtil.ConnectDB(resource, pool)
	runtimeCurrent, _ := filepath.Abs("../")
	InsertSeasonLeagueStats(runtimeCurrent+"/test/resource", db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func TestGetTeamID(t *testing.T) {
	type args struct {
		teamNames []string
	}
	tests := []struct {
		name       string
		args       args
		wantTeamID string
	}{
		{
			"ジャイアンツのチームIDを取得",
			args{teamNames: []string{"巨 人", "巨　人", "読 売ジャイアンツ", "読　売ジャイアンツ", "読　売", "g"}},
			"01",
		},
		{
			"ベイスターズのチームIDを取得",
			args{teamNames: []string{"横 浜", "横 浜ベイスターズ", "DeNA", "横浜DeNAベイスターズ", "横　浜", "横浜DeNA", "db"}},
			"02",
		},
		{
			"タイガースのチームIDを取得",
			args{teamNames: []string{"阪 神", "阪　神", "阪 神タイガース", "阪　神タイガース", "t"}},
			"03",
		},
		{
			"カープのチームIDを取得",
			args{teamNames: []string{"広 島", "広　島", "広島東洋カープ", "広島東洋", "c"}},
			"04",
		},
		{
			"ドラゴンズのチームIDを取得",
			args{teamNames: []string{"中 日", "中　日", "中 日ドラゴンズ", "中　日ドラゴンズ", "d"}},
			"05",
		},
		{
			"ヤクルトのチームIDを取得",
			args{teamNames: []string{"ヤクルト", "ヤクルトスワローズ", "東京ヤクルトスワローズ", "東京ヤクルト", "s"}},
			"06",
		},
		{
			"ライオンズのチームIDを取得",
			args{teamNames: []string{"西 武", "西　武", "西 武ライオンズ", "埼玉西武ライオンズ", "埼玉西武", "l"}},
			"07",
		},
		{
			"ホークスのチームIDを取得",
			args{teamNames: []string{"ソフトバンク", "福岡ソフトバンクホークス", "福岡ソフトバンク", "福岡ダイエー", "h"}},
			"08",
		},
		{
			"イーグルスのチームIDを取得",
			args{teamNames: []string{"楽 天", "楽　天", "東北楽天ゴールデンイーグルス", "東北楽天", "e"}},
			"09",
		},
		{
			"マリーンズのチームIDを取得",
			args{teamNames: []string{"ロッテ", "千葉ロッテマリーンズ", "千葉ロッテ", "m"}},
			"10",
		},
		{
			"ファイターズのチームIDを取得",
			args{teamNames: []string{"日本ハム", "北海道日本ハムファイターズ", "北海道日本ハム", "f"}},
			"11",
		},
		{
			"バファローズのチームIDを取得",
			args{teamNames: []string{"オリックス", "オリックスバファローズ", "大阪近鉄", "b"}},
			"12",
		},
		{
			"予期せぬチーム名",
			args{teamNames: []string{"unknown"}},
			"13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, teamName := range tt.args.teamNames {
				assert.Equal(t, tt.wantTeamID, GetTeamID(teamName))
			}
		})
	}
}

func TestInsertTeamPitchings(t *testing.T) {
	type args struct {
		teamID                   string
		year                     string
		leage                    string
		expectedEarnedRunAverage float64
		expectedGames            int
		expectedWin              int
		expectedLose             int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"セリーグ投手成績登録確認",
			args{
				teamID:                   "04",
				year:                     "2005",
				leage:                    "central",
				expectedEarnedRunAverage: 4.8,
				expectedGames:            146,
				expectedWin:              58,
				expectedLose:             84,
			},
		},
		{
			"パリーグ投手成績登録確認",
			args{
				teamID:                   "08",
				year:                     "2005",
				leage:                    "pacific",
				expectedEarnedRunAverage: 3.46,
				expectedGames:            136,
				expectedWin:              89,
				expectedLose:             45,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)

			runtimeCurrent, _ := filepath.Abs("../")
			InsertTeamPitchings(runtimeCurrent+"/test/resource", tt.args.leage, db)

			pitching := GetTeamPitching([]int{2005}, db)["2005"][0]

			assert.Equal(t, tt.args.expectedEarnedRunAverage, pitching.EarnedRunAverage)
			assert.Equal(t, tt.args.expectedGames, pitching.Games)
			assert.Equal(t, tt.args.expectedWin, pitching.Win)
			assert.Equal(t, tt.args.expectedLose, pitching.Lose)

			testUtil.CloseContainer(resource, pool)
		})
	}
}

func TestInsertTeamBattings(t *testing.T) {
	type args struct {
		teamID                  string
		year                    string
		league                  string
		expectedBattingAverage  float64
		expectedGames           int
		expectedPlateAppearance int
		expectedAtBat           int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"セリーグ打撃成績登録確認",
			args{
				teamID:                  "06",
				year:                    "2005",
				league:                  "central",
				expectedBattingAverage:  0.276,
				expectedGames:           146,
				expectedPlateAppearance: 5523,
				expectedAtBat:           5033,
			},
		},
		{
			"パリーグ打撃成績登録確認",
			args{
				teamID:                  "09",
				year:                    "2005",
				league:                  "pacific",
				expectedBattingAverage:  0.255,
				expectedGames:           136,
				expectedPlateAppearance: 5068,
				expectedAtBat:           4577,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			db := testUtil.ConnectDB(resource, pool)

			runtimeCurrent, _ := filepath.Abs("../")
			InsertTeamBattings(runtimeCurrent+"/test/resource", tt.args.league, db)

			batting := GetTeamBatting([]int{2005}, db)["2005"][0]

			assert.Equal(t, tt.args.expectedBattingAverage, batting.BattingAverage)
			assert.Equal(t, tt.args.expectedGames, batting.Games)
			assert.Equal(t, tt.args.expectedPlateAppearance, batting.PlateAppearance)
			assert.Equal(t, tt.args.expectedAtBat, batting.AtBat)

			testUtil.CloseContainer(resource, pool)
		})
	}
}
