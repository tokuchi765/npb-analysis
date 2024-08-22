package repository

import (
	"github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/entity/team"
)

// SyastemRepository システム設定を登録するインターフェース
type SyastemRepository interface {
	GetSystemSetting(setting string) (value string)
	SetSystemSetting(setting string, value string)
}

// GradesRepository チーム成績データアクセスを管理するインターフェース
type GradesRepository interface {
	GetPitchings(playerID string) (pitchings []player.PitcherGrades)
	GetBattings(playerID string) (battings []player.BatterGrades)
	GetPlayers(playerID string) (career player.Players)
	InsertPlayers(careers []player.Players)
	InsertPicherGrades(picher player.PitcherGrades)
	InsertBatterGrades(batterGrades player.BatterGrades)
	SearchCareerByName(name string) (careers []player.Players)
}

// TeamRepository チーム成績データアクセスを管理するインターフェース
type TeamRepository interface {
	InsertTeamPitchings(teamPitching team.TeamPitching)
	GetTeamPitchings(years []int) (teamPitchingMap map[string][]team.TeamPitching)
	GetTeamPitchingByTeamIDAndYear(year string, teamID string) (teamPitching team.TeamPitching)
	GetTeamPitchingMax() (maxStrikeOutRate float64, maxRunsAllowed int)
	GetTeamPitchingMin() (minStrikeOutRate float64, minRunsAllowed int)
	InsertTeamBattings(teamBatting team.TeamBatting)
	GetTeamBattings(years []int) (teamBattingMap map[string][]team.TeamBatting)
	GetTeamBattingByTeamIDAndYear(teamID string, year string) (teamBatting team.TeamBatting)
	GetTeamBattingMax() (maxHomeRun int, maxSluggingPercentage float64, maxOnBasePercentage float64)
	GetTeamBattingMin() (minHomeRun int, minSluggingPercentage float64, minOnBasePercentage float64)
	GetTeamStats(years []int) (teamStatsMap map[string][]team.TeamSeasonStats)
	InsertPythagoreanExpectation(teamBattings []team.TeamBatting, teamPitchings []team.TeamPitching)
	InsertTeamSeasonStats(teamSeasonStats []team.TeamSeasonStats)
	InsertMatchResults(teamMatchResults []team.TeamMatchResults)
	GetTeamName(teamID string) (teamName string)
	GetPlayersByTeamIDAndYear(teamID string, year string) (players []team.TeamPlayers)
	InsertTeamPlayers(members []team.TeamPlayers)
	InsertMembersCsv(fileName string)
	IsRegisteredMembersCsv(fileName string) bool
}
