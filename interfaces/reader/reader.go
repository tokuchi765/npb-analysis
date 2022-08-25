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
}

type GradesReader interface {
	GetPlayers(csvPath string, initial string, year string) (players [][]string)
	ReadCareer(csvPath string, initial string, playerID string, playerName string) (career player.CAREER, exsist bool)
	ReadGrades(csvPath string, initial string, playerID string, playerName string) (picherGradesList []player.PICHERGRADES, batterGradesList []player.BATTERGRADES, exsist bool)
}
