package infrastructure

import (
	"reflect"
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

// SearchBatterGrades 引数で受け取った条件で野手成績を検索する
func (Repository *GradesRepository) SearchBatterGrades(condition player.SearchBatterGradesCondition) (results []player.SearchBatterGradesResult) {

	if reflect.DeepEqual(condition, reflect.Zero(reflect.TypeOf(condition)).Interface()) {
		return results
	}

	query := Repository.GormDB.Table("batter_grades AS bg").Select("bg.player_id,ps.name,bg.year,bg.team,bg.games,bg.plate_appearance,bg.at_bat,bg.on_base_percentage,bg.batting_average,bg.slugging_percentage,bg.w_oba,bg.rc,bg.babip,bg.base_on_balls,bg.hit_by_pitches,bg.strike_out,bg.strike_out_rate,bg.score,bg.hit,bg.home_run,bg.runs_batted_in,bg.single,bg.double,bg.triple,bg.base_hit,bg.stolen_base,bg.caught_stealing,bg.sacrifice_hits,bg.sacrifice_flies,bg.grounded_into_double_play").Joins("inner join players AS ps ON ps.player_id = bg.player_id")

	if condition.Total {
		query.Where("EXISTS (SELECT 1 FROM batter_grades sub_bg WHERE sub_bg.year != 'nan' AND bg.player_id = sub_bg.player_id GROUP BY sub_bg.player_id HAVING COUNT(*) >= ?)", condition.TotalYear)
		query.Where("bg.year = 'nan'")
	} else if condition.FromYear > 0 && condition.ToYear > 0 {
		query.Where("CAST(REPLACE(bg.year, 'nan', '0') AS INTEGER) BETWEEN ? AND ?", condition.FromYear, condition.ToYear)
	} else {
		query.Where("bg.year != 'nan'")
	}

	if condition.TeamID != "" {
		query.Where("bg.team_id = ?", condition.TeamID)
	}

	if condition.PlateAppearance > 0 {
		query.Where("bg.plate_appearance "+condition.PlateAppearanceThresholdType.String()+" ?", condition.PlateAppearance)
	}

	if condition.OnBasePercentage > 0 {
		query.Where("bg.on_base_percentage "+condition.OnBasePercentageThresholdType.String()+" ?", condition.OnBasePercentage)
	}

	if condition.BattingAverage > 0 {
		query.Where("bg.batting_average "+condition.BattingAverageThresholdType.String()+" ?", condition.BattingAverage)
	}

	if condition.SluggingPercentage > 0 {
		query.Where("bg.slugging_percentage "+condition.SluggingPercentageThresholdType.String()+" ?", condition.SluggingPercentage)
	}

	if condition.HomeRun > 0 {
		query.Where("bg.home_run "+condition.HomeRunThresholdType.String()+" ?", condition.HomeRun)
	}

	if condition.StrikeOut > 0 {
		query.Where("bg.strike_out "+condition.StrikeOutThresholdType.String()+" ?", condition.StrikeOut)
	}

	if condition.StrikeOutRate > 0 {
		query.Where("bg.strike_out_rate "+condition.StrikeOutRateThresholdType.String()+" ?", condition.StrikeOutRate)
	}

	if condition.GroundedIntoDoublePlay > 0 {
		query.Where("bg.grounded_into_double_play "+condition.GroundedIntoDoublePlayThresholdType.String()+" ?", condition.GroundedIntoDoublePlay)
	}

	if condition.WOba > 0 {
		query.Where("bg.w_oba "+condition.WObaThresholdType.String()+" ?", condition.WOba)
	}

	if condition.RC > 0 {
		query.Where("bg.rc "+condition.RCThresholdType.String()+" ?", condition.RC)
	}

	if condition.Babip > 0 {
		query.Where("bg.babip "+condition.BabipThresholdType.String()+" ?", condition.Babip)
	}

	query.Find(&results)
	return results
}

// SearchPicherGrades 引数で受け取った条件で投手成績を検索する
func (Repository *GradesRepository) SearchPicherGrades(condition player.SearchPitcherGradesCondition) (results []player.SearchPitcherGradesResult) {

	if reflect.DeepEqual(condition, reflect.Zero(reflect.TypeOf(condition)).Interface()) {
		return results
	}

	query := Repository.GormDB.Table("pitcher_grades AS pg").Select("pg.player_id,ps.name,pg.player_id,pg.year,pg.team_id,pg.team,pg.pitched,pg.win,pg.lose,pg.save,pg.hold,pg.hold_point,pg.complete_game,pg.shutout,pg.no_walks,pg.winning_rate,pg.batter,pg.innings_pitched,pg.hit,pg.home_run,pg.base_on_balls,pg.hit_by_pitches,pg.strike_out,pg.wild_pitches,pg.balk,pg.runs_allowed,pg.earned_run,pg.earned_run_average,pg.babip,pg.strike_out_rate").Joins("inner join players AS ps ON ps.player_id = pg.player_id")

	if condition.Total {
		query.Where("EXISTS (SELECT 1 FROM pitcher_grades sub_pg WHERE sub_pg.year != 'nan' AND pg.player_id = sub_pg.player_id GROUP BY sub_pg.player_id HAVING COUNT(*) >= ?)", condition.TotalYear)
		query.Where("pg.year = 'nan'")
	} else if condition.FromYear > 0 && condition.ToYear > 0 {
		query.Where("CAST(REPLACE(pg.year, 'nan', '0') AS INTEGER) BETWEEN ? AND ?", condition.FromYear, condition.ToYear)
	} else {
		query.Where("pg.year != 'nan'")
	}

	if condition.TeamID != "" {
		query.Where("pg.team_id = ?", condition.TeamID)
	}

	if condition.Pitched > 0 {
		query.Where("pg.pitched "+condition.PitchedThresholdType.String()+" ?", condition.Pitched)
	}

	if condition.InningsPitched > 0 {
		query.Where("pg.innings_pitched "+condition.InningsPitchedThresholdType.String()+" ?", condition.InningsPitched)
	}

	if condition.EarnedRunAverage > 0 {
		query.Where("pg.earned_run_average "+condition.EarnedRunAverageThresholdType.String()+" ?", condition.EarnedRunAverage)
	}

	if condition.Babip > 0 {
		query.Where("pg.babip "+condition.BabipThresholdType.String()+" ?", condition.Babip)
	}

	if condition.StrikeOutRate > 0 {
		query.Where("pg.strike_out_rate "+condition.StrikeOutRateThresholdType.String()+" ?", condition.StrikeOutRate)
	}

	if condition.StrikeOut > 0 {
		query.Where("pg.strike_out "+condition.StrikeOutThresholdType.String()+" ?", condition.StrikeOut)
	}

	if condition.Hit > 0 {
		query.Where("pg.hit "+condition.HitThresholdType.String()+" ?", condition.Hit)
	}

	if condition.BaseOnBalls > 0 {
		query.Where("pg.base_on_balls "+condition.BaseOnBallsThresholdType.String()+" ?", condition.BaseOnBalls)
	}

	if condition.HomeRun > 0 {
		query.Where("pg.home_run "+condition.HomeRunThresholdType.String()+" ?", condition.HomeRun)
	}

	if condition.Win > 0 {
		query.Where("pg.win "+condition.WinThresholdType.String()+" ?", condition.Win)
	}

	if condition.Lose > 0 {
		query.Where("pg.lose "+condition.LoseThresholdType.String()+" ?", condition.Lose)
	}

	if condition.Save > 0 {
		query.Where("pg.save "+condition.SaveThresholdType.String()+" ?", condition.Save)
	}

	if condition.Hold > 0 {
		query.Where("pg.hold "+condition.HoldThresholdType.String()+" ?", condition.Hold)
	}

	if condition.HoldPoint > 0 {
		query.Where("pg.hold_point "+condition.HoldPointThresholdType.String()+" ?", condition.HoldPoint)
	}

	if condition.CompleteGame > 0 {
		query.Where("pg.complete_game "+condition.CompleteGameThresholdType.String()+" ?", condition.CompleteGame)
	}

	if condition.Shutout > 0 {
		query.Where("pg.shutout "+condition.ShutoutThresholdType.String()+" ?", condition.Shutout)
	}

	if condition.WinningRate > 0 {
		query.Where("pg.winning_rate "+condition.WinningRateThresholdType.String()+" ?", condition.WinningRate)
	}

	query.Find(&results)
	return results
}
