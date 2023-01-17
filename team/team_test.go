package team

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	teamData "github.com/tokuchi765/npb-analysis/entity/team"
	mock_reader "github.com/tokuchi765/npb-analysis/interfaces/reader/mock"
	mock_repository "github.com/tokuchi765/npb-analysis/interfaces/repository/mock"
)

func TestTeamInteractor_InsertPythagoreanExpectation(t *testing.T) {
	type args struct {
		years           []int
		teamBattingMap  map[string][]teamData.TeamBatting
		teamPitchingMap map[string][]teamData.TeamPitching
	}
	teamBatting := teamData.TeamBatting{TeamID: "01", Year: "2020", Score: 100}
	teamPitching := teamData.TeamPitching{TeamID: "01", Year: "2020", RunsAllowed: 100}
	tests := []struct {
		name string
		args args
	}{
		{
			"ピタゴラス勝率登録テスト",
			args{
				years: []int{2020},
				teamBattingMap: map[string][]teamData.TeamBatting{"2020": {
					teamBatting,
				}},
				teamPitchingMap: map[string][]teamData.TeamPitching{"2020": {
					teamPitching,
				}},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().InsertPythagoreanExpectation([]teamData.TeamBatting{teamBatting}, []teamData.TeamPitching{teamPitching})

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			interactor.InsertPythagoreanExpectation(tt.args.years, tt.args.teamBattingMap, tt.args.teamPitchingMap)
		})
	}
}

func TestTeamInteractor_GetTeamStats(t *testing.T) {
	type args struct {
		teamLeagueStats map[string][]teamData.TeamLeagueStats
		years           []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チーム成績取得",
			args{
				teamLeagueStats: map[string][]teamData.TeamLeagueStats{"2020": {{TeamID: "01"}}},
				years:           []int{2020},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamStats(tt.args.years).Return(tt.args.teamLeagueStats)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			actual := interactor.GetTeamStats(tt.args.years)
			assert.Exactly(t, tt.args.teamLeagueStats, actual)
		})
	}
}

func TestTeamInteractor_InsertSeasonMatchResults(t *testing.T) {
	type args struct {
		year                      string
		years                     []int
		cTeamMatchResults         []teamData.TeamMatchResults
		pTeamMatchResults         []teamData.TeamMatchResults
		cTeamExchangeMatchResults []teamData.TeamMatchResults
		pTeamExchangeMatchResults []teamData.TeamMatchResults
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"リーグ対戦成績登録確認",
			args{
				year:  "2020",
				years: []int{2020},
				cTeamMatchResults: []teamData.TeamMatchResults{
					{TeamID: "01", Year: "2020", CompetitiveTeamID: "02", VsType: "league", Win: 7, Lose: 8, Draw: 3},
				},
				pTeamMatchResults: []teamData.TeamMatchResults{
					{TeamID: "10", Year: "2020", CompetitiveTeamID: "06", VsType: "league", Win: 7, Lose: 5, Draw: 3},
				},
			},
		},
		{
			"交流戦対戦成績登録確認",
			args{
				year:  "2005",
				years: []int{2005},
				cTeamMatchResults: []teamData.TeamMatchResults{
					{TeamID: "02", Year: "2005", CompetitiveTeamID: "06", VsType: "league", Win: 7, Lose: 8, Draw: 3},
				},
				pTeamMatchResults: []teamData.TeamMatchResults{
					{TeamID: "09", Year: "2005", CompetitiveTeamID: "07", VsType: "league", Win: 7, Lose: 5, Draw: 3},
				},
				cTeamExchangeMatchResults: []teamData.TeamMatchResults{
					{TeamID: "02", Year: "2005", CompetitiveTeamID: "10", VsType: "exchange", Win: 2, Lose: 1, Draw: 1},
				},
				pTeamExchangeMatchResults: []teamData.TeamMatchResults{
					{TeamID: "09", Year: "2005", CompetitiveTeamID: "04", VsType: "exchange", Win: 1, Lose: 2, Draw: 1},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			csvPath := "csvpath"
			mTeamReader.EXPECT().ReadTeamLeagueStats(csvPath, "c", tt.args.year).Return([]teamData.TeamLeagueStats{}, tt.args.cTeamMatchResults)
			mTeamReader.EXPECT().ReadTeamLeagueStats(csvPath, "p", tt.args.year).Return([]teamData.TeamLeagueStats{}, tt.args.pTeamMatchResults)

			if tt.args.year != "2020" {
				mTeamReader.EXPECT().ReadTeamExchangeStats(csvPath, "c", tt.args.year).Return(tt.args.cTeamExchangeMatchResults)
				mTeamReader.EXPECT().ReadTeamExchangeStats(csvPath, "p", tt.args.year).Return(tt.args.pTeamExchangeMatchResults)

				tt.args.cTeamMatchResults = append(tt.args.cTeamMatchResults, tt.args.cTeamExchangeMatchResults...)
				tt.args.pTeamMatchResults = append(tt.args.pTeamMatchResults, tt.args.pTeamExchangeMatchResults...)
			}

			mTeamRepository.EXPECT().InsertMatchResults(append(tt.args.cTeamMatchResults, tt.args.pTeamMatchResults...))

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			interactor.InsertSeasonMatchResults(csvPath, tt.args.years)

			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_InsertSeasonLeagueStats(t *testing.T) {
	type args struct {
		year             string
		years            []int
		cTeamLeagueStats []teamData.TeamLeagueStats
		pTeamLeagueStats []teamData.TeamLeagueStats
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チームシーズン成績登録確認",
			args{
				year:  "2015",
				years: []int{2015},
				cTeamLeagueStats: []teamData.TeamLeagueStats{
					{TeamID: "01", Year: "2015", Games: 144, Win: 80, Lose: 60, Draw: 4},
				},
				pTeamLeagueStats: []teamData.TeamLeagueStats{
					{TeamID: "09", Year: "2015", Games: 144, Win: 60, Lose: 80, Draw: 4},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			csvPath := "csvpath"

			mTeamReader.EXPECT().ReadTeamLeagueStats(csvPath, "c", tt.args.year).Return(tt.args.cTeamLeagueStats, []teamData.TeamMatchResults{})
			mTeamReader.EXPECT().ReadTeamLeagueStats(csvPath, "p", tt.args.year).Return(tt.args.pTeamLeagueStats, []teamData.TeamMatchResults{})

			cManager := "セントラル監督"
			mTeamReader.EXPECT().ReadManager(csvPath, tt.args.cTeamLeagueStats[0].TeamID, tt.args.year).Return(cManager)
			pManager := "パシフィック監督"
			mTeamReader.EXPECT().ReadManager(csvPath, tt.args.pTeamLeagueStats[0].TeamID, tt.args.year).Return(pManager)

			tt.args.cTeamLeagueStats[0].Manager = cManager
			tt.args.pTeamLeagueStats[0].Manager = pManager
			mTeamRepository.EXPECT().InsertTeamLeagueStats(append(tt.args.cTeamLeagueStats, tt.args.pTeamLeagueStats...))

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			interactor.InsertSeasonLeagueStats(csvPath, tt.args.years)
			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_InsertTeamPitchings(t *testing.T) {
	type args struct {
		year         string
		leage        string
		years        []int
		teamPitching []teamData.TeamPitching
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"投手成績登録",
			args{
				year:  "2005",
				leage: "central",
				years: []int{2005},
				teamPitching: []teamData.TeamPitching{
					{TeamID: "04", Year: "2005", EarnedRunAverage: 0.35},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			csvPath := "csvpath"
			mTeamReader.EXPECT().ReadTeamPitching(csvPath, tt.args.leage, tt.args.year).Return(tt.args.teamPitching)

			mTeamRepository.EXPECT().InsertTeamPitchings(tt.args.teamPitching[0])

			interactor.InsertTeamPitchings(csvPath, tt.args.leage, tt.args.years)

			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_GetTeamPitching(t *testing.T) {
	type args struct {
		years           []int
		teamPitchingMap map[string][]teamData.TeamPitching
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"投手成績登録取得",
			args{
				years: []int{2005},
				teamPitchingMap: map[string][]teamData.TeamPitching{"2005": {
					{TeamID: "04", Year: "2005", EarnedRunAverage: 0.35},
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			mTeamRepository.EXPECT().GetTeamPitchings(tt.args.years).Return(tt.args.teamPitchingMap)

			actual := interactor.GetTeamPitching(tt.args.years)
			assert.Exactly(t, tt.args.teamPitchingMap, actual)

			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_InsertTeamBattings(t *testing.T) {
	type args struct {
		year        string
		years       []int
		league      string
		teamBatting []teamData.TeamBatting
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"打撃成績登録",
			args{
				year:   "2005",
				years:  []int{2005},
				league: "central",
				teamBatting: []teamData.TeamBatting{
					{TeamID: "05", Year: "2005", BattingAverage: 0.28},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			expect := tt.args.teamBatting[0]
			expect.SetStrikeOutRate()
			mTeamRepository.EXPECT().InsertTeamBattings(expect)

			csvPath := "csvPath"
			mTeamReader.EXPECT().ReadTeamBatting(csvPath, tt.args.league, tt.args.year).Return(tt.args.teamBatting)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			interactor.InsertTeamBattings(csvPath, tt.args.league, tt.args.years)

			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_GetTeamBatting(t *testing.T) {
	type args struct {
		years          []int
		teamBattingMap map[string][]teamData.TeamBatting
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"打撃成績取得",
			args{
				years: []int{2005},
				teamBattingMap: map[string][]teamData.TeamBatting{
					"2005": {
						{TeamID: "05", Year: "2005", BattingAverage: 0.28}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamReader := mock_reader.NewMockTeamReader(mockCtrl)

			mTeamRepository.EXPECT().GetTeamBattings(tt.args.years).Return(tt.args.teamBattingMap)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
				TeamReader:     mTeamReader,
			}

			actual := interactor.GetTeamBatting(tt.args.years)

			assert.Exactly(t, tt.args.teamBattingMap, actual)

			mockCtrl.Finish()
		})
	}
}

func TestTeamInteractor_GetTeamPitchingByTeamIDAndYear(t *testing.T) {
	type args struct {
		teamID       string
		year         string
		teamPitching teamData.TeamPitching
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"チーム投手成績取得（チームIDと年指定）",
			args{
				teamID: "04",
				year:   "2005",
				teamPitching: teamData.TeamPitching{
					TeamID:           "04",
					Year:             "2005",
					EarnedRunAverage: 3.5,
				},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			mTeamRepository.EXPECT().GetTeamPitchingByTeamIDAndYear(tt.args.year, tt.args.teamID).Return(tt.args.teamPitching)

			actual := interactor.GetTeamPitchingByTeamIDAndYear(tt.args.year, tt.args.teamID)

			assert.Exactly(t, tt.args.teamPitching, actual)
		})
	}
}

func TestTeamInteractor_GetTeamBattingByTeamIDAndYear(t *testing.T) {
	type args struct {
		teamID      string
		year        string
		teamBatting teamData.TeamBatting
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"セリーグ打撃成績登録確認",
			args{
				teamID: "06",
				year:   "2005",
				teamBatting: teamData.TeamBatting{
					TeamID:         "06",
					Year:           "2005",
					BattingAverage: 0.3,
				},
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			mTeamRepository.EXPECT().GetTeamBattingByTeamIDAndYear(tt.args.teamID, tt.args.year).Return(tt.args.teamBatting)

			actual := interactor.GetTeamBattingByTeamIDAndYear(tt.args.teamID, tt.args.year)

			assert.Exactly(t, tt.args.teamBatting, actual)
		})
	}
}

func TestTeamInteractor_GetTeamPitchingMax(t *testing.T) {
	tests := []struct {
		name                 string
		wantMaxStrikeOutRate float64
		wantMaxRunsAllowed   int
	}{
		{
			name:                 "チーム投手成績の各項目の最大値取得",
			wantMaxStrikeOutRate: 8.6,
			wantMaxRunsAllowed:   750,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamPitchingMax().Return(tt.wantMaxStrikeOutRate, tt.wantMaxRunsAllowed)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			gotMaxStrikeOutRate, gotMaxRunsAllowed := interactor.GetTeamPitchingMax()

			assert.Equal(t, gotMaxStrikeOutRate, tt.wantMaxStrikeOutRate)
			assert.Equal(t, gotMaxRunsAllowed, tt.wantMaxRunsAllowed)
		})
	}
}

func TestTeamInteractor_GetTeamPitchingMin(t *testing.T) {
	tests := []struct {
		name                 string
		wantMinStrikeOutRate float64
		wantMinRunsAllowed   int
	}{
		{
			name:                 "チーム投手成績の各項目の最小値取得",
			wantMinStrikeOutRate: 5.4,
			wantMinRunsAllowed:   456,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamPitchingMin().Return(tt.wantMinStrikeOutRate, tt.wantMinRunsAllowed)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			gotMinStrikeOutRate, gotMinRunsAllowed := interactor.GetTeamPitchingMin()

			assert.Equal(t, gotMinStrikeOutRate, tt.wantMinStrikeOutRate)
			assert.Equal(t, gotMinRunsAllowed, tt.wantMinRunsAllowed)
		})
	}
}

func TestTeamInteractor_GetTeamBattingMax(t *testing.T) {
	tests := []struct {
		name                      string
		wantMaxHomeRun            int
		wantMaxSluggingPercentage float64
		wantMaxOnBasePercentage   float64
	}{
		{
			name:                      "チーム打撃成績の各項目の最大値取得",
			wantMaxHomeRun:            580,
			wantMaxSluggingPercentage: 0.48,
			wantMaxOnBasePercentage:   0.36,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamBattingMax().Return(tt.wantMaxHomeRun, tt.wantMaxSluggingPercentage, tt.wantMaxOnBasePercentage)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			gotMaxHomeRun, gotMaxSluggingPercentage, gotMaxOnBasePercentage := interactor.GetTeamBattingMax()

			assert.Equal(t, gotMaxHomeRun, tt.wantMaxHomeRun)
			assert.Equal(t, gotMaxSluggingPercentage, tt.wantMaxSluggingPercentage)
			assert.Equal(t, gotMaxOnBasePercentage, tt.wantMaxOnBasePercentage)
		})
	}
}

func TestTeamInteractor_GetTeamBattingMin(t *testing.T) {
	tests := []struct {
		name                      string
		wantMinHomeRun            int
		wantMinSluggingPercentage float64
		wantMinOnBasePercentage   float64
	}{
		{
			name:                      "チーム打撃成績の各項目の最小値取得",
			wantMinHomeRun:            123,
			wantMinSluggingPercentage: 0.29,
			wantMinOnBasePercentage:   0.23,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mTeamRepository := mock_repository.NewMockTeamRepository(mockCtrl)

			mTeamRepository.EXPECT().GetTeamBattingMin().Return(tt.wantMinHomeRun, tt.wantMinSluggingPercentage, tt.wantMinOnBasePercentage)

			interactor := TeamInteractor{
				TeamRepository: mTeamRepository,
			}

			gotMinHomeRun, gotMinSluggingPercentage, gotMinOnBasePercentage := interactor.GetTeamBattingMin()

			assert.Equal(t, gotMinHomeRun, tt.wantMinHomeRun)
			assert.Equal(t, gotMinSluggingPercentage, tt.wantMinSluggingPercentage)
			assert.Equal(t, gotMinOnBasePercentage, tt.wantMinOnBasePercentage)
		})
	}
}
