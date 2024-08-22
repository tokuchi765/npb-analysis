package grades

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tokuchi765/npb-analysis/entity/player"
	mock_reader "github.com/tokuchi765/npb-analysis/interfaces/reader/mock"
	mock_repository "github.com/tokuchi765/npb-analysis/interfaces/repository/mock"
)

func TestInsertCareers(t *testing.T) {
	type args struct {
		csvPath string
		careers []player.Players
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績登録",
			args{
				"csvPath",
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertPlayers(tt.args.careers)

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			mGradesReader.EXPECT().ReadCareers(tt.args.csvPath).Return([]player.CAREER{getCAREER()})

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				GradesReader:     mGradesReader,
			}
			interactor.InsertCareers(tt.args.csvPath)
		})
	}
}

func getCAREER() player.CAREER {
	return player.CAREER{
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
}

func getPicherGrades() player.PitcherGrades {
	return player.PitcherGrades{
		PlayerID:         "53355134",
		Year:             "2018",
		TeamID:           "12",
		Team:             "オリックス",
		Pitched:          54.0,
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
		InningsPitched:   53.3,
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
		Babip:            0.2482758620689655,
		StrikeOutRate:    7.76735459662289,
	}
}

func getTestPicherGrades() player.PICHERGRADES {
	return player.PICHERGRADES{
		Year:             "2018",
		TeamID:           "12",
		Team:             "オリックス",
		Pitched:          54.0,
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
		InningsPitched:   53.1,
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

func getBatterGrades() player.BatterGrades {
	return player.BatterGrades{
		PlayerID:               "01605136",
		Year:                   "2018",
		TeamID:                 "12",
		Team:                   "オリックス",
		Games:                  113,
		PlateAppearance:        345,
		AtBat:                  295,
		Score:                  39,
		Hit:                    78,
		Single:                 65,
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
		StrikeOutRate:          0.09565217391304348,
		GroundedIntoDoublePlay: 2,
		BattingAverage:         0.264,
		SluggingPercentage:     0.32899999999999996,
		OnBasePercentage:       0.34,
		WOba:                   0.30729483282674774,
		RC:                     36.13817391304349,
		Babip:                  0.2950191570881226,
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
		RC:                     36.138172,
	}
}

func TestInsertPicherGrades(t *testing.T) {
	playerID := "53355134"
	picherMap := make(map[string][]player.PICHERGRADES)
	picherGrades := getTestPicherGrades()
	picherMap[playerID] = []player.PICHERGRADES{picherGrades}
	wantPicherMap := make(map[string][]player.PICHERGRADES)
	wantPicherGrades := getTestPicherGrades()
	wantPicherGrades.SetInningsPitched()
	wantPicherGrades.SetBABIP()
	wantPicherGrades.SetStrikeOutRate()
	wantPicherMap[playerID] = []player.PICHERGRADES{wantPicherGrades}
	type args struct {
		picherMap     map[string][]player.PICHERGRADES
		wantPicherMap map[string][]player.PICHERGRADES
		playerID      string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"投手成績登録",
			args{
				picherMap,
				wantPicherMap,
				playerID,
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().InsertPicherGrades(getPicherGrades()).Times(1)

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			mGradesReader.EXPECT().ReadPitcherGrades(gomock.Any()).Return(tt.args.picherMap)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				GradesReader:     mGradesReader,
			}

			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertPicherGrades(runtimeCurrent)
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
		wantPitchings []player.PitcherGrades
	}{
		{
			"投手成績取得",
			args{
				"53355134",
			},
			[]player.PitcherGrades{getPicherGrades()},
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
	batterMap := make(map[string][]player.BATTERGRADES)
	grades := getTestBatterGrades()
	batterMap[playerID] = []player.BATTERGRADES{grades}
	type args struct {
		batterMap map[string][]player.BATTERGRADES
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

			mGradesRepository.EXPECT().InsertBatterGrades(getBatterGrades()).Times(1)

			mGradesReader := mock_reader.NewMockGradesReader(mockCtrl)

			mGradesReader.EXPECT().ReadBatterGrades(gomock.Any()).Return(tt.args.batterMap)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
				GradesReader:     mGradesReader,
			}
			runtimeCurrent, _ := filepath.Abs("../")
			interactor.InsertBatterGrades(runtimeCurrent)
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
		wantBattings []player.BatterGrades
	}{
		{
			"打席成績取得",
			args{
				"01605136",
			},
			[]player.BatterGrades{getBatterGrades()},
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
		career   player.Players
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"選手成績取得",
			args{
				"01105137",
				player.Players{
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
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().GetPlayers(tt.args.playerID).Return(tt.args.career)

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

func TestGradesInteractor_SearchCareerByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantCareer []player.Players
	}{
		{
			"選手名検索",
			args{
				"飯田",
			},
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
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mGradesRepository := mock_repository.NewMockGradesRepository(mockCtrl)

			mGradesRepository.EXPECT().SearchCareerByName(tt.args.name).Return(tt.wantCareer)

			interactor := GradesInteractor{
				GradesRepository: mGradesRepository,
			}

			gotCareer := interactor.SearchCareerByName(tt.args.name)
			assert.Exactly(t, tt.wantCareer, gotCareer)
		})
	}
}
