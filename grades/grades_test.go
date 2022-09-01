package grades

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	mock_reader "github.com/tokuchi765/npb-analysis/interfaces/reader/mock"
	mock_repository "github.com/tokuchi765/npb-analysis/interfaces/repository/mock"
)

func TestInsertTeamPlayers(t *testing.T) {
	type args struct {
		initial  string
		players  [][]string
		teamID   string
		year     string
		teamName string
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
				"Giants",
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertTeamPlayers(tt.args.teamID, tt.args.teamName, tt.args.players, tt.args.year)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamName(tt.args.teamID).Return(tt.args.teamName)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			interactor.InsertTeamPlayers(tt.args.initial, tt.args.players, tt.args.year)
		})
	}
}

func TestGradesInteractor_TestReadCareers(t *testing.T) {
	career := data.CAREER{
		PlayerID:           "01105137",
		Name:               "飯田　優也",
		Position:           "投手",
		PitchingAndBatting: "左投左打",
		Height:             "187cm",
		Weight:             "92kg",
		Birthday:           "1990年11月27日",
		Career:             "神戸弘陵高 - 東京農業大生産学部",
		Draft:              "2012年育成選手ドラフト3位",
	}
	type args struct {
		initial    string
		player     string
		id         string
		playerName string
		players    [][]string
		career     data.CAREER
	}
	tests := []struct {
		name           string
		args           args
		wantCareerList []data.CAREER
	}{
		{
			"選手成績読み込み",
			args{
				initial:    "b",
				player:     "/bis/players/01105137.html",
				id:         "01105137",
				playerName: "飯田　優也",
				players: [][]string{
					{"/bis/players/01105137.html", "飯田　優也"},
				},
				career: career,
			},
			[]data.CAREER{career},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			csvPath := "csvPath"
			mGradesReader.EXPECT().ReadCareer(csvPath, tt.args.initial, tt.args.id, tt.args.playerName).Return(tt.args.career, true)

			interactor := GradesInteractor{
				GradesReader: mGradesReader,
			}

			actual := interactor.ReadCareers(csvPath, tt.args.initial, tt.args.players)

			assert.ElementsMatch(t, tt.wantCareerList, actual)
		})
	}
}

func TestInsertCareers(t *testing.T) {
	type args struct {
		careers []data.CAREER
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績登録",
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertCareers(tt.args.careers)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}
			interactor.InsertCareers(tt.args.careers)
		})
	}
}

func TestReadGradesMap(t *testing.T) {
	type args struct {
		initial  string
		playName string
		players  [][]string
		playerID string
	}
	tests := []struct {
		name       string
		args       args
		wantPicher []data.PICHERGRADES
		wantBatter []data.BATTERGRADES
	}{
		{
			name: "野手選手成績読み込み",
			args: args{
				"b",
				"テストプレイヤー1",
				[][]string{
					{"/bis/players/01605136.html", "テストプレイヤー1"},
				},
				"01605136",
			},
			wantBatter: []data.BATTERGRADES{
				{
					Year:   "2020",
					TeamID: "09",
					Games:  144,
				},
			},
		},
		{
			name: "投手選手成績読み込み",
			args: args{
				"b",
				"テストプレイヤー2",
				[][]string{
					{"/bis/players/53355134.html", "テストプレイヤー2"},
				},
				"53355134",
			},
			wantPicher: []data.PICHERGRADES{
				{
					Year:   "2020",
					TeamID: "09",
					Piched: 60,
				},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			csvPath := "csvPath"
			mGradesReader.EXPECT().ReadGrades(csvPath, tt.args.initial, tt.args.playerID, tt.args.playName).Return(tt.wantPicher, tt.wantBatter, true)

			interactor := GradesInteractor{
				GradesReader: mGradesReader,
			}

			gotPicherMap, gotBatterMap := interactor.ReadGradesMap(csvPath, tt.args.initial, tt.args.players)

			assert.Exactly(t, tt.wantPicher, gotPicherMap[tt.args.playerID])
			assert.Exactly(t, tt.wantBatter, gotBatterMap[tt.args.playerID])
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
	agsPicherGrades := getTestPicherGrades()
	agsPicherGrades.SetBABIP()
	agsPicherGrades.SetStrikeOutRate()
	picherMap[playerID] = []data.PICHERGRADES{picherGrades}
	type args struct {
		picherMap       map[string][]data.PICHERGRADES
		playerID        string
		agsPicherGrades data.PICHERGRADES
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
				agsPicherGrades,
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertPicherGrades(tt.args.playerID, tt.args.agsPicherGrades).Times(1)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			interactor.InsertPicherGrades(tt.args.picherMap)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().GetPitchings(tt.args.playerID).Return(tt.wantPitchings)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			gotPitchings := interactor.GetPitching(tt.args.playerID)
			assert.ElementsMatch(t, tt.wantPitchings, gotPitchings)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertBatterGrades(tt.args.playerID, gomock.Any())

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}
			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertBatterGrades(tt.args.batterMap, runtimeCurrent)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().GetBattings(tt.args.playerID).Return(tt.wantBattings)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			gotBattings := interactor.GetBatting(tt.args.playerID)
			assert.ElementsMatch(t, tt.wantBattings, gotBattings)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().GetCareer(tt.args.playerID).Return(tt.args.career)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().GetPlayersByTeamIDAndYear(tt.args.teamID, tt.args.year).Return(tt.wantPlayers)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			csvPath := "csvPath"
			mGradesReader.EXPECT().GetPlayers(csvPath, tt.args.initial, tt.args.year).Return(tt.wantPlayers)

			interactor := GradesInteractor{
				GradesReader: mGradesReader,
			}

			gotPlayers := interactor.GetPlayers(csvPath, tt.args.initial, tt.args.year)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().ExtractionCareers(&tt.args.careers)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			interactor.ExtractionCareers(&tt.args.careers)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().ExtractionPicherGrades(&tt.args.picherMap, tt.args.teamID)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			interactor.ExtractionPicherGrades(&tt.args.picherMap, tt.args.teamID)
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().ExtractionBatterGrades(&tt.args.batterMap, tt.args.teamID)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				TeamRepository:   mTeamRepository,
			}

			interactor.ExtractionBatterGrades(&tt.args.batterMap, tt.args.teamID)
		})
	}
}
