package reader

import (
	"github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/entity/team"
)

// TeamReader チーム成績CSVの読み込みを管理する
type TeamReader interface {
	ReadTeamLeagueStats(csvPath string, league string, year string) (teamLeagueStats []team.TeamLeagueStats, teamMatchResults []team.TeamMatchResults)
	ReadManager(csvPath string, teamID string, year string) (manager string)
	ReadTeamExchangeStats(csvPath string, league string, year string) (teamExchangeMatchResults []team.TeamMatchResults)
	ReadTeamPitching(csvPath string, league string, year string) (teamPitching []team.TeamPitching)
	ReadTeamBatting(csvPath string, league string, year string) (teamBatting []team.TeamBatting)
	ReadTeamPlayers(csvPath string, initial string, teamName string) (players map[string][]team.Member)
}

// GradesReader 選手情報CSVの読み込みを管理する
type GradesReader interface {
	ReadCareers(csvPath string) (careers []player.CAREER)
	ReadBatterGrades(csvPath string) (batterGrades map[string][]player.BATTERGRADES)
	ReadPitcherGrades(csvPath string) (pitcherGrades map[string][]player.PICHERGRADES)
}
