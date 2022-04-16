package grades

import (
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/infrastructure"
	testUtil "github.com/tokuchi765/npb-analysis/test"
)

func TestInsertTeamPlayers(t *testing.T) {
	type args struct {
		initial string
		players [][]string
		teamID  string
		year    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手一覧登録",
			args{
				"g",
				[][]string{
					{"93795138", "デラロサ"},
					{"41045138", "戸郷　翔征"},
				},
				"01",
				"2020",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertTeamPlayers(tt.args.initial, tt.args.players, tt.args.year)

			rows, _ := db.Query("SELECT player_id,player_name FROM team_players WHERE year = $1 AND team_id = $2", "2020", tt.args.teamID)

			var actual [][]string
			for rows.Next() {
				var prayerID, playerName string
				rows.Scan(&prayerID, &playerName)
				actual = append(actual, []string{prayerID, playerName})
			}

			assert.ElementsMatch(t, tt.args.players, actual)
		})
	}
}

func TestGradesInteractor_TestReadCareers(t *testing.T) {
	type args struct {
		initial string
		players [][]string
	}
	tests := []struct {
		name           string
		args           args
		wantCareerList []data.CAREER
	}{
		{
			"選手成績読み込み",
			args{
				initial: "b",
				players: [][]string{
					{"/bis/players/01105137.html", "飯田　優也"},
				},
			},
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
	}
	interactor := GradesInteractor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runtimeCurrent, _ := filepath.Abs("../")
			actual := interactor.ReadCareers(runtimeCurrent+"/test/resource/", tt.args.initial, tt.args.players)
			assert.Exactly(t, tt.wantCareerList, actual)
		})
	}
}

func TestInsertCareers(t *testing.T) {
	type args struct {
		playerID string
		careers  []data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績登録",
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
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertCareers(tt.args.careers)
			rows, _ := db.Query("SELECT name,position,pitching_and_batting FROM players WHERE player_id = $1", tt.args.playerID)
			var name, position, pitchingAndBatting string
			for rows.Next() {
				rows.Scan(&name, &position, &pitchingAndBatting)
			}
			assert.Equal(t, tt.args.careers[0].Name, name)
			assert.Equal(t, tt.args.careers[0].Position, position)
			assert.Equal(t, tt.args.careers[0].PitchingAndBatting, pitchingAndBatting)
		})
	}
}

func TestReadGradesMap(t *testing.T) {
	grades := getTestBatterGrades()
	grades.RC = 0.0 // 読み込み時に算出しない値を0にする
	type args struct {
		initial string
		players [][]string
	}
	tests := []struct {
		name       string
		args       args
		pitcherID  string
		batterID   string
		wantPicher data.PICHERGRADES
		wantBatter data.BATTERGRADES
	}{
		{
			"選手成績読み込み",
			args{
				"b",
				[][]string{
					{"/bis/players/53355134.html", "山本　由伸"},
					{"/bis/players/01605136.html", "福田　周平"},
				},
			},
			"53355134",
			"01605136",
			getTestPicherGrades(),
			grades,
		},
	}
	runtimeCurrent, _ := filepath.Abs("../")
	interactor := GradesInteractor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPicherMap, gotBatterMap := interactor.ReadGradesMap(runtimeCurrent+"/test/resource/", tt.args.initial, tt.args.players)
			assert.Exactly(t, tt.wantPicher, gotPicherMap[tt.pitcherID][0])
			assert.Exactly(t, tt.wantBatter, gotBatterMap[tt.batterID][0])
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
		RC:                     36.138172,
	}
}

func TestInsertPicherGrades(t *testing.T) {
	playerID := "53355134"
	picherMap := make(map[string][]data.PICHERGRADES)
	picherGrades := getTestPicherGrades()
	picherMap[playerID] = []data.PICHERGRADES{picherGrades}
	type args struct {
		picherMap map[string][]data.PICHERGRADES
		playerID  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"投手成績登録",
			args{
				picherMap,
				playerID,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertPicherGrades(tt.args.picherMap)
			rows, _ := db.Query("SELECT team,piched,earned_run_average FROM picher_grades WHERE player_id = $1 AND year = $2", tt.args.playerID, "2018")
			var team string
			var piched, earnedRunAverage float64
			for rows.Next() {
				rows.Scan(&team, &piched, &earnedRunAverage)
			}
			assert.Equal(t, picherGrades.Team, team)
			assert.Equal(t, picherGrades.Piched, piched)
			assert.Equal(t, picherGrades.EarnedRunAverage, earnedRunAverage)
		})
	}
}

func TestGradesInteractor_GetPitching(t *testing.T) {
	type args struct {
		playerID string
	}
	tests := []struct {
		name          string
		args          args
		wantPitchings []data.PICHERGRADES
	}{
		{
			"投手成績取得",
			args{
				"53355134",
			},
			[]data.PICHERGRADES{getTestPicherGrades()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			picherMap := make(map[string][]data.PICHERGRADES)
			picherGrades := getTestPicherGrades()
			picherMap[tt.args.playerID] = []data.PICHERGRADES{picherGrades}
			interactor.InsertPicherGrades(picherMap)
			gotPitchings := interactor.GetPitching(tt.args.playerID)
			assert.Equal(t, 0.24827586, gotPitchings[0].BABIP)
			assert.Equal(t, 7.811321, gotPitchings[0].StrikeOutRate)
		})
	}
}

func TestInsertBatterGrades(t *testing.T) {
	playerID := "01605136"
	batterMap := make(map[string][]data.BATTERGRADES)
	grades := getTestBatterGrades()
	batterMap[playerID] = []data.BATTERGRADES{grades}
	type args struct {
		batterMap map[string][]data.BATTERGRADES
		playerID  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"打者成績登録",
			args{
				batterMap,
				playerID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertBatterGrades(tt.args.batterMap, runtimeCurrent)
			rows, _ := db.Query("SELECT team,plate_appearance,single,w_oba,rc FROM batter_grades WHERE player_id = $1 AND year = $2", tt.args.playerID, "2018")
			var team string
			var plateAppearance, single int
			var wOba, rc float64
			for rows.Next() {
				rows.Scan(&team, &plateAppearance, &single, &wOba, &rc)
			}
			assert.Equal(t, grades.Team, team)
			assert.Equal(t, grades.PlateAppearance, plateAppearance)
			assert.Equal(t, 65, single)
			assert.Equal(t, 0.30729485, wOba)
			assert.Equal(t, 36.138172, rc)
		})
	}
}

func TestGradesInteractor_GetBatting(t *testing.T) {
	type args struct {
		playerID string
	}
	tests := []struct {
		name         string
		args         args
		wantBattings []data.BATTERGRADES
	}{
		{
			"打席成績取得",
			args{
				"01605136",
			},
			[]data.BATTERGRADES{getTestBatterGrades()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			batterMap := make(map[string][]data.BATTERGRADES)
			batterMap[tt.args.playerID] = []data.BATTERGRADES{getTestBatterGrades()}
			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertBatterGrades(batterMap, runtimeCurrent)
			gotBattings := interactor.GetBatting(tt.args.playerID)
			assert.Equal(t, 0.30729485, gotBattings[0].Woba)
			assert.Equal(t, 36.138172, gotBattings[0].RC)
			assert.Equal(t, 0.29501915, gotBattings[0].BABIP)
		})
	}
}

func TestGradesInteractor_GetCareer(t *testing.T) {
	type args struct {
		playerID string
		career   data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績取得",
			args{
				"01105137",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertCareers([]data.CAREER{tt.args.career})
			gotCareer := interactor.GetCareer(tt.args.playerID)
			assert.Exactly(t, tt.args.career, gotCareer)
		})
	}
}

func TestGradesInteractor_GetPlayersByTeamIDAndYear(t *testing.T) {
	type args struct {
		teamID string
		year   string
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
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			players := [][]string{
				{"93795138", "デラロサ"},
				{"41045138", "戸郷　翔征"},
			}
			interactor.InsertTeamPlayers("g", players, tt.args.year)
			gotPlayers := interactor.GetPlayersByTeamIDAndYear(tt.args.teamID, tt.args.year)
			assert.ElementsMatch(t, tt.wantPlayers, gotPlayers)
		})
	}
}

func TestGradesInteractor_TestGetPlayers(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}

			runtimeCurrent, _ := filepath.Abs("../")
			gotPlayers := interactor.GetPlayers(runtimeCurrent+"/test/resource/", tt.args.initial, tt.args.year)
			assert.ElementsMatch(t, tt.wantPlayers, gotPlayers)
		})
	}
}

func TestExtractionCareers(t *testing.T) {
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
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertCareers(tt.args.careers)
			interactor.ExtractionCareers(&tt.args.careers)
			assert.Empty(t, tt.args.careers)
		})
	}
}

func TestExtractionPicherGrades(t *testing.T) {
	playerID := "53355134"
	picherMap := make(map[string][]data.PICHERGRADES)
	picherGrades := getTestPicherGrades()
	picherMap[playerID] = []data.PICHERGRADES{picherGrades}
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
				picherMap,
				"12",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			interactor.InsertPicherGrades(tt.args.picherMap)
			interactor.ExtractionPicherGrades(&tt.args.picherMap, tt.args.teamID)
			assert.Empty(t, tt.args.picherMap)
		})
	}
}

func TestExtractionBatterGrades(t *testing.T) {
	playerID := "01605136"
	batterMap := make(map[string][]data.BATTERGRADES)
	grades := getTestBatterGrades()
	batterMap[playerID] = []data.BATTERGRADES{grades}
	type args struct {
		batterMap map[string][]data.BATTERGRADES
		teamID    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"重複打撃成績を削除する",
			args{
				batterMap,
				"12",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, pool := testUtil.CreateContainer()
			defer testUtil.CloseContainer(resource, pool)
			db := testUtil.ConnectDB(resource, pool)
			sqlHandler := new(infrastructure.SQLHandler)
			sqlHandler.Conn = db
			interactor := GradesInteractor{
				GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
				TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
			}
			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertBatterGrades(tt.args.batterMap, runtimeCurrent)
			interactor.ExtractionBatterGrades(&tt.args.batterMap, tt.args.teamID)
			assert.Empty(t, tt.args.batterMap)
		})
	}
}
