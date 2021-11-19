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

// GetTeamBatting 引数で受け取った年に紐づくチーム打撃成績を取得します。
func (Interactor *TeamInteractor) GetTeamBatting(years []int) (teamBattingMap map[string][]teamData.TeamBatting) {
	return Interactor.TeamRepository.GetTeamBattings(years)
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func (Interactor *TeamInteractor) GetTeamStats(years []int) (teamStatsMap map[string][]teamData.TeamLeagueStats) {
	return Interactor.TeamRepository.GetTeamStats(years)
}

// InsertSeasonLeagueStats チームごとのシーズン成績をDBに登録する
func (Interactor *TeamInteractor) InsertSeasonLeagueStats(csvPath string) {
	years := makeRange(2005, 2020)
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
func (Interactor *TeamInteractor) InsertSeasonMatchResults(csvPath string) {
	years := makeRange(2005, 2020)
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
func (Interactor *TeamInteractor) InsertTeamPitchings(csvPath string, league string) {
	years := makeRange(2005, 2020)
	for _, year := range years {
		teamPitching := Interactor.ReadTeamPitching(csvPath, league, strconv.Itoa(year))
		Interactor.TeamRepository.InsertTeamPitchings(teamPitching)
	}
}

// InsertTeamBattings チーム打撃成績をDBに登録する
func (Interactor *TeamInteractor) InsertTeamBattings(csvPath string, league string) {
	years := makeRange(2005, 2020)
	for _, year := range years {
		teamBatting := Interactor.ReadTeamBatting(csvPath, league, strconv.Itoa(year))
		Interactor.TeamRepository.InsertTeamBattings(teamBatting)
	}
}

// GetTeamID 引数で受け取ったチーム名とイニシャルからチームIDを取得する
func GetTeamID(teamName string) (teamID string) {
	idDatas := map[string][]string{
		"01": {"巨 人", "巨　人", "読 売ジャイアンツ", "読　売ジャイアンツ", "読　売", "g"},
		"02": {"横 浜", "横 浜ベイスターズ", "DeNA", "横浜DeNAベイスターズ", "横　浜", "横浜DeNA", "db"},
		"03": {"阪 神", "阪　神", "阪 神タイガース", "阪　神タイガース", "t"},
		"04": {"広 島", "広　島", "広島東洋カープ", "広島東洋", "c"},
		"05": {"中 日", "中　日", "中 日ドラゴンズ", "中　日ドラゴンズ", "d"},
		"06": {"ヤクルト", "ヤクルトスワローズ", "東京ヤクルトスワローズ", "東京ヤクルト", "s"},
		"07": {"西 武", "西　武", "西 武ライオンズ", "埼玉西武ライオンズ", "埼玉西武", "l"},
		"08": {"ソフトバンク", "福岡ソフトバンクホークス", "福岡ソフトバンク", "福岡ダイエー", "h"},
		"09": {"楽 天", "楽　天", "東北楽天ゴールデンイーグルス", "東北楽天", "e"},
		"10": {"ロッテ", "千葉ロッテマリーンズ", "千葉ロッテ", "m"},
		"11": {"日本ハム", "北海道日本ハムファイターズ", "北海道日本ハム", "f"},
		"12": {"オリックス", "オリックスバファローズ", "大阪近鉄", "b"},
	}
	for key, idData := range idDatas {
		for _, idName := range idData {
			if idName == teamName {
				return key
			}
		}
	}
	return "13"
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
