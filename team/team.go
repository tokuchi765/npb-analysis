package team

import (
	"strconv"

	teamData "github.com/tokuchi765/npb-analysis/entity/team"
	csvReader "github.com/tokuchi765/npb-analysis/infrastructure/csv"
	"github.com/tokuchi765/npb-analysis/interfaces/repository"
)

// TeamInteractor チーム情報処理のInteractor
type TeamInteractor struct {
	repository.TeamRepository
	csvReader.TeamReader
}

// InsertPythagoreanExpectation ピタゴラス勝率をDBに登録します。
func (Interactor *TeamInteractor) InsertPythagoreanExpectation(years []int, teamBattingMap map[string][]teamData.TeamBatting, teamPitchingMap map[string][]teamData.TeamPitching) {
	for _, year := range years {
		strYear := strconv.Itoa(year)
		teamBattings := teamBattingMap[strYear]
		teamPitchings := teamPitchingMap[strYear]
		Interactor.TeamRepository.InsertPythagoreanExpectation(teamBattings, teamPitchings)
	}
}

// GetTeamPitching 引数で受け取った年に紐づくチーム投手成績を取得します。
func (Interactor *TeamInteractor) GetTeamPitching(years []int) (teamPitchingMap map[string][]teamData.TeamPitching) {
	return Interactor.TeamRepository.GetTeamPitchings(years)
}

// GetTeamPitchingByTeamIDAndYear 引数で受け取った年とチームIDに紐づくチーム投手成績を取得します。
func (Interactor *TeamInteractor) GetTeamPitchingByTeamIDAndYear(year string, teamID string) (teamPitching teamData.TeamPitching) {
	return Interactor.TeamRepository.GetTeamPitchingByTeamIDAndYear(year, teamID)
}

// GetTeamPitchingMax チーム投手成績の各項目の最大値を取得する。
func (Interactor *TeamInteractor) GetTeamPitchingMax() (maxStrikeOutRate float64, maxRunsAllowed int) {
	return Interactor.TeamRepository.GetTeamPitchingMax()
}

// GetTeamBatting 引数で受け取った年に紐づくチーム打撃成績を取得します。
func (Interactor *TeamInteractor) GetTeamBatting(years []int) (teamBattingMap map[string][]teamData.TeamBatting) {
	return Interactor.TeamRepository.GetTeamBattings(years)
}

// GetTeamBattingByTeamIDAndYear 引数で受け取った年とチームIDに紐づくチーム打撃成績を取得します。
func (Interactor *TeamInteractor) GetTeamBattingByTeamIDAndYear(teamID string, year string) (teamBatting teamData.TeamBatting) {
	return Interactor.TeamRepository.GetTeamBattingByTeamIDAndYear(teamID, year)
}

// GetTeamBattingMax チーム打撃成績の各項目の最大値を取得する。
func (Interactor *TeamInteractor) GetTeamBattingMax() (maxHomeRun int, maxSluggingPercentage float64, maxOnBasePercentage float64) {
	return Interactor.TeamRepository.GetTeamBattingMax()
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func (Interactor *TeamInteractor) GetTeamStats(years []int) (teamStatsMap map[string][]teamData.TeamLeagueStats) {
	return Interactor.TeamRepository.GetTeamStats(years)
}

// InsertSeasonLeagueStats チームごとのシーズン成績をDBに登録する
func (Interactor *TeamInteractor) InsertSeasonLeagueStats(csvPath string, years []int) {
	for _, year := range years {
		cTeamLeagueStats, _ := Interactor.TeamReader.ReadTeamLeagueStats(csvPath, "c", strconv.Itoa(year))
		pTeamLeagueStats, _ := Interactor.TeamReader.ReadTeamLeagueStats(csvPath, "p", strconv.Itoa(year))

		Interactor.setManager(csvPath, &cTeamLeagueStats)
		Interactor.setManager(csvPath, &pTeamLeagueStats)

		// DBに登録する
		Interactor.TeamRepository.InsertTeamLeagueStats(append(cTeamLeagueStats, pTeamLeagueStats...))
	}
}

func (Interactor *TeamInteractor) setManager(csvPath string, teamLeagueStatsList *[]teamData.TeamLeagueStats) {
	teamLeagueStatses := *teamLeagueStatsList
	for i, teamLeagueStats := range teamLeagueStatses {
		teamLeagueStatses[i].Manager = Interactor.ReadManager(csvPath, teamLeagueStats.TeamID, teamLeagueStats.Year)
	}
}

// InsertSeasonMatchResults 各チームの対戦成績をDBに登録する
func (Interactor *TeamInteractor) InsertSeasonMatchResults(csvPath string, years []int) {
	for _, year := range years {
		_, cTeamMatchResults := Interactor.TeamReader.ReadTeamLeagueStats(csvPath, "c", strconv.Itoa(year))
		_, pTeamMatchResults := Interactor.TeamReader.ReadTeamLeagueStats(csvPath, "p", strconv.Itoa(year))

		if year != 2020 {
			cTeamExchangeMatchResults := Interactor.TeamReader.ReadTeamExchangeStats(csvPath, "c", strconv.Itoa(year))
			cTeamMatchResults = append(cTeamMatchResults, cTeamExchangeMatchResults...)
			pTeamExchangeMatchResults := Interactor.TeamReader.ReadTeamExchangeStats(csvPath, "p", strconv.Itoa(year))
			pTeamMatchResults = append(pTeamMatchResults, pTeamExchangeMatchResults...)
		}

		// DBに登録する
		Interactor.TeamRepository.InsertMatchResults(append(cTeamMatchResults, pTeamMatchResults...))
	}
}

// InsertTeamPitchings チーム投手成績をDBに登録する
func (Interactor *TeamInteractor) InsertTeamPitchings(csvPath string, league string, years []int) {
	for _, year := range years {
		teamPitching := Interactor.ReadTeamPitching(csvPath, league, strconv.Itoa(year))
		for _, pitching := range teamPitching {
			pitching.SetBABIP()
			pitching.SetStrikeOutRate()
			Interactor.TeamRepository.InsertTeamPitchings(pitching)
		}
	}
}

// InsertTeamBattings チーム打撃成績をDBに登録する
func (Interactor *TeamInteractor) InsertTeamBattings(csvPath string, league string, years []int) {
	for _, year := range years {
		teamBatting := Interactor.ReadTeamBatting(csvPath, league, strconv.Itoa(year))
		for _, batting := range teamBatting {
			batting.SetBABIP()
			Interactor.TeamRepository.InsertTeamBattings(batting)
		}
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
