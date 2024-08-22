package infrastructure

import (
	"strings"

	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/entity/player"
)

// GradesRepository チーム成績データアクセスを管理するリポジトリ
type GradesRepository struct {
	SQLHandler
}

// GetPitchings 個人投手成績一覧を取得する
func (Repository *GradesRepository) GetPitchings(playerID string) (pitchings []player.PitcherGrades) {
	Repository.GormDB.Where("player_id = ?", playerID).Find(&pitchings)
	return pitchings
}

// GetBattings 個人打撃成績一覧を取得する
func (Repository *GradesRepository) GetBattings(playerID string) (battings []player.BatterGrades) {
	Repository.GormDB.Where("player_id = ?", playerID).Find(&battings)
	return battings
}

// GetPlayers 選手情報を取得する
func (Repository *GradesRepository) GetPlayers(playerID string) (players player.Players) {
	Repository.GormDB.Where("player_id = ?", playerID).Find(&players)
	return players
}

// SearchCareerByName 選手名から選手データを検索する
func (Repository *GradesRepository) SearchCareerByName(name string) (careers []player.Players) {
	Repository.GormDB.Where("search_name LIKE ?", "%"+name+"%").Find(&careers)
	return careers
}

func extractionPlayerID(url string) string {
	return strings.Replace(strings.Replace(url, "/bis/players/", "", 1), ".html", "", 1)
}

// InsertPlayers 引数で受け取った選手情報をDBへ登録する
func (Repository *GradesRepository) InsertPlayers(players []player.Players) {
	Repository.GormDB.Create(&players)
}

// InsertPicherGrades 引数で受け取った投手成績データをテーブルに追加する
func (Repository *GradesRepository) InsertPicherGrades(picherGrades player.PitcherGrades) {
	Repository.GormDB.Create(&picherGrades)
}

// InsertBatterGrades 引数で受け取ったBATTERGRADESをDBに登録する
func (Repository *GradesRepository) InsertBatterGrades(batterGrades player.BatterGrades) {
	Repository.GormDB.Create(&batterGrades)
}
