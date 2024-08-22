package infrastructure

import (
	"strconv"

	"github.com/tokuchi765/npb-analysis/entity/team"
)

// TeamRepository チーム成績データアクセスを管理するリポジトリ
type TeamRepository struct {
	SQLHandler
}

const MEMBERS = "MEMBERS"

// InsertTeamPitchings チーム投手成績をDBに登録する
func (Repository *TeamRepository) InsertTeamPitchings(pitching team.TeamPitching) {
	Repository.GormDB.Create(&pitching)
}

// GetTeamPitchings 引数で受け取った年に紐づくチーム投手成績を取得します。
func (Repository *TeamRepository) GetTeamPitchings(years []int) (teamPitchingMap map[string][]team.TeamPitching) {
	teamPitchingMap = make(map[string][]team.TeamPitching)

	for _, year := range years {
		strYear := strconv.Itoa(year)
		var teamPitchings []team.TeamPitching
		Repository.GormDB.Where("year = ?", strYear).Find(&teamPitchings)

		teamPitchingMap[strYear] = teamPitchings
	}
	return teamPitchingMap
}

// GetTeamPitchingByTeamIDAndYear 引数で受け取ったチームIDと年に紐づくチーム投手成績を取得します。
func (Repository *TeamRepository) GetTeamPitchingByTeamIDAndYear(year string, teamID string) (teamPitching team.TeamPitching) {
	Repository.GormDB.Where("year = ?", year).Where("team_id = ?", teamID).Find(&teamPitching)
	return teamPitching
}

// GetTeamPitchingMax チーム投手成績の各項目の最大値を取得する。
func (Repository *TeamRepository) GetTeamPitchingMax() (maxStrikeOutRate float64, maxRunsAllowed int) {
	type MaxResult struct {
		MaxStrikeOutRate float64
		MaxRunsAllowed   int
	}
	var result MaxResult
	Repository.GormDB.Raw("select max(strike_out_rate) as max_strike_out_rate, max(runs_allowed) as max_runs_allowed from team_pitching").Scan(&result)
	return result.MaxStrikeOutRate, result.MaxRunsAllowed
}

// GetTeamPitchingMin チーム投手成績の各項目の最小値を取得する。
func (Repository *TeamRepository) GetTeamPitchingMin() (minStrikeOutRate float64, minRunsAllowed int) {
	type MinResult struct {
		MinStrikeOutRate float64
		MinRunsAllowed   int
	}
	var result MinResult
	Repository.GormDB.Raw("select min(strike_out_rate) as min_strike_out_rate, min(runs_allowed) as min_runs_allowed from team_pitching").Scan(&result)
	return result.MinStrikeOutRate, result.MinRunsAllowed
}

// InsertTeamBattings チーム打撃成績をDBに登録する
func (Repository *TeamRepository) InsertTeamBattings(batting team.TeamBatting) {
	Repository.GormDB.Create(&batting)
}

// GetTeamBattings 引数で受け取った年に紐づくチーム打撃成績を取得します。
func (Repository *TeamRepository) GetTeamBattings(years []int) (teamBattingMap map[string][]team.TeamBatting) {
	teamBattingMap = make(map[string][]team.TeamBatting)
	for _, year := range years {
		strYear := strconv.Itoa(year)
		var teamBattins []team.TeamBatting
		Repository.GormDB.Where("year = ?", strYear).Find(&teamBattins)

		teamBattingMap[strYear] = teamBattins
	}

	return teamBattingMap
}

// GetTeamBattingByTeamIDAndYear 引数で受け取ったチームIDと年に紐づくチーム野手成績を取得します。
func (Repository *TeamRepository) GetTeamBattingByTeamIDAndYear(teamID string, year string) (teamBatting team.TeamBatting) {
	Repository.GormDB.Where("year = ?", year).Where("team_id = ?", teamID).Find(&teamBatting)
	return teamBatting
}

// GetTeamBattingMax チーム打撃成績の各項目の最大値を取得する。
func (Repository *TeamRepository) GetTeamBattingMax() (maxHomeRun int, maxSluggingPercentage float64, maxOnBasePercentage float64) {
	type MaxResult struct {
		MaxHomeRun            int
		MaxSluggingPercentage float64
		MaxOnBasePercentage   float64
	}
	var result MaxResult
	Repository.GormDB.Raw("select max(home_run) as max_home_run, max(slugging_percentage) as max_slugging_percentage, max(on_base_percentage) as max_on_base_percentage from team_batting").Scan(&result)
	return result.MaxHomeRun, result.MaxSluggingPercentage, result.MaxOnBasePercentage
}

// GetTeamBattingMin チーム打撃成績の各項目の最小値を取得する。
func (Repository *TeamRepository) GetTeamBattingMin() (minHomeRun int, minSluggingPercentage float64, minOnBasePercentage float64) {
	type MinResult struct {
		MinHomeRun            int
		MinSluggingPercentage float64
		MinOnBasePercentage   float64
	}
	var result MinResult
	Repository.GormDB.Raw("select min(home_run) as min_home_run, min(slugging_percentage) as min_slugging_percentage, min(on_base_percentage) as min_on_base_percentage from team_batting").Scan(&result)
	return result.MinHomeRun, result.MinSluggingPercentage, result.MinOnBasePercentage
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func (Repository *TeamRepository) GetTeamStats(years []int) (teamStatsMap map[string][]team.TeamSeasonStats) {
	teamStatsMap = make(map[string][]team.TeamSeasonStats)
	for _, year := range years {
		strYear := strconv.Itoa(year)

		var teamStatses []team.TeamSeasonStats
		Repository.GormDB.Where("year = ?", strYear).Find(&teamStatses)

		teamStatsMap[strYear] = teamStatses
	}
	return teamStatsMap
}

// InsertPythagoreanExpectation ピタゴラス勝率をDBに登録します。
func (Repository *TeamRepository) InsertPythagoreanExpectation(teamBattings []team.TeamBatting, teamPitchings []team.TeamPitching) {
	for _, teamBatting := range teamBattings {
		for _, teamPitching := range teamPitchings {
			if teamBatting.TeamID == teamPitching.TeamID {
				pythagoreanExpectation := calcPythagoreanExpectation(teamBatting.Score, teamPitching.RunsAllowed)

				Repository.GormDB.Model(&team.TeamSeasonStats{}).Where("team_id = ?", teamPitching.TeamID).Where("year = ?", teamPitching.Year).Update("pythagorean_expectation", pythagoreanExpectation)
			}
		}
	}
}

func calcPythagoreanExpectation(score int, runsAllowed int) float64 {
	fScore := float64(score)
	fRunsAllowed := float64(runsAllowed)
	return (fScore * fScore) / ((fScore * fScore) + (fRunsAllowed * fRunsAllowed))
}

// InsertTeamSeasonStats チームごとのシーズン成績をDBに登録する
func (Repository *TeamRepository) InsertTeamSeasonStats(teamLeagueStats []team.TeamSeasonStats) {
	Repository.GormDB.Create(&teamLeagueStats)
}

// InsertMatchResults 各チームの対戦成績をDBに登録する
func (Repository *TeamRepository) InsertMatchResults(teamMatchResults []team.TeamMatchResults) {
	Repository.GormDB.Create(&teamMatchResults)
}

// GetTeamName マスタテーブルからチーム名を取得する
func (Repository *TeamRepository) GetTeamName(teamID string) (teamName string) {
	Repository.GormDB.Raw("SELECT team_name FROM team_name WHERE team_name_id = ?", teamID).Scan(&teamName)
	return teamName
}

// InsertTeamPlayers 年度ごとの選手一覧をDBに登録する
func (Repository *TeamRepository) InsertTeamPlayers(players []team.TeamPlayers) {
	Repository.GormDB.Create(&players)
}

// GetPlayersByTeamIDAndYear チームIDと年から選手一覧を取得する
func (Repository *TeamRepository) GetPlayersByTeamIDAndYear(teamID string, year string) (players []team.TeamPlayers) {
	Repository.GormDB.Where("year = ?", year).Where("team_id = ?", teamID).Find(&players)
	return players
}

type RegisteredCsv struct {
	CsvType  string
	FilePath string
}

// InsertMembersCsv メンバー一覧CSVファイルを登録します
func (Repository *TeamRepository) InsertMembersCsv(filePath string) {
	registeredCsv := RegisteredCsv{MEMBERS, filePath}
	Repository.GormDB.Create(&registeredCsv)
}

// IsRegisteredMembersCsv メンバー一覧CSVファイルが登録済みか確認する
func (Repository *TeamRepository) IsRegisteredMembersCsv(filePath string) bool {
	var count int64
	Repository.GormDB.Model(&RegisteredCsv{}).Where("csv_type = ?", MEMBERS).Where("file_path = ?", filePath).Count(&count)
	return count > 0
}
