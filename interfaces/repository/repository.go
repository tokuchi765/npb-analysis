package repository

import (
	data "github.com/tokuchi765/npb-analysis/entity/player"
	teamData "github.com/tokuchi765/npb-analysis/entity/team"
)

// SyastemRepository システム設定を登録するインターフェース
type SyastemRepository interface {
	GetSystemSetting(setting string) (value string)
	SetSystemSetting(setting string, value string)
}

// GradesRepository チーム成績データアクセスを管理するインターフェース
type GradesRepository interface {
	GetPitchings(playerID string) (pitchings []data.PICHERGRADES)
	GetBattings(playerID string) (battings []data.BATTERGRADES)
	GetCareer(playerID string) (career data.CAREER)
	GetPlayersByTeamIDAndYear(teamID string, year string) (players []data.PLAYER)
	InsertTeamPlayers(teamID string, teamName string, players [][]string)
	ExtractionCareers(careers *[]data.CAREER)
	InsertCareers(careers []data.CAREER)
	ExtractionPicherGrades(picherMap *map[string][]data.PICHERGRADES, teamID string)
	InsertPicherGrades(picherMap map[string][]data.PICHERGRADES)
	ExtractionBatterGrades(batterMap *map[string][]data.BATTERGRADES, teamID string)
	InsertBatterGrades(playerID string, batterGrades data.BATTERGRADES)
}

// TeamRepository チーム成績データアクセスを管理するインターフェース
type TeamRepository interface {
	InsertTeamPitchings(teamPitching []teamData.TeamPitching)
	GetTeamPitchings(years []int) (teamPitchingMap map[string][]teamData.TeamPitching)
	InsertTeamBattings(teamBatting []teamData.TeamBatting)
	GetTeamBattings(years []int) (teamBattingMap map[string][]teamData.TeamBatting)
	GetTeamStats(years []int) (teamStatsMap map[string][]teamData.TeamLeagueStats)
	InsertPythagoreanExpectation(teamBattings []teamData.TeamBatting, teamPitchings []teamData.TeamPitching)
	InsertTeamLeagueStats(teamLeagueStats []teamData.TeamLeagueStats)
	InsertMatchResults(teamMatchResults []teamData.TeamMatchResults)
	GetTeamName(teamID string) (teamName string)
}
