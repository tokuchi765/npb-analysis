package infrastructure

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	data "github.com/tokuchi765/npb-analysis/entity/player"
)

// GradesRepository チーム成績データアクセスを管理するリポジトリ
type GradesRepository struct {
	SQLHandler
}

// GetPitchings 個人投手成績一覧を取得する
func (Repository *GradesRepository) GetPitchings(playerID string) (pitchings []data.PICHERGRADES) {
	rows, err := Repository.Conn.Query("SELECT * FROM picher_grades WHERE player_id = $1", playerID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var playerID string
		var pitching data.PICHERGRADES
		rows.Scan(&playerID, &pitching.Year, &pitching.TeamID, &pitching.Team,
			&pitching.Piched, &pitching.Win, &pitching.Lose, &pitching.Save,
			&pitching.Hold, &pitching.HoldPoint, &pitching.CompleteGame, &pitching.Shutout,
			&pitching.NoWalks, &pitching.WinningRate, &pitching.Batter, &pitching.InningsPitched,
			&pitching.Hit, &pitching.HomeRun, &pitching.BaseOnBalls, &pitching.HitByPitches,
			&pitching.StrikeOut, &pitching.WildPitches, &pitching.Balk, &pitching.RunsAllowed,
			&pitching.EarnedRun, &pitching.EarnedRunAverage, &pitching.BABIP)

		pitchings = append(pitchings, pitching)
	}

	return pitchings
}

// GetBattings 個人打撃成績一覧を取得する
func (Repository *GradesRepository) GetBattings(playerID string) (battings []data.BATTERGRADES) {
	rows, err := Repository.Conn.Query("SELECT * FROM batter_grades WHERE player_id = $1", playerID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var playerID string
		var batting data.BATTERGRADES
		rows.Scan(&playerID, &batting.Year, &batting.TeamID, &batting.Team, &batting.Games,
			&batting.PlateAppearance, &batting.AtBat, &batting.Score, &batting.Hit, &batting.Single,
			&batting.Double, &batting.Triple, &batting.HomeRun, &batting.BaseHit,
			&batting.RunsBattedIn, &batting.StolenBase, &batting.CaughtStealing, &batting.SacrificeHits,
			&batting.SacrificeFlies, &batting.BaseOnBalls, &batting.HitByPitches, &batting.StrikeOut,
			&batting.GroundedIntoDoublePlay, &batting.BattingAverage, &batting.SluggingPercentage, &batting.OnBasePercentage,
			&batting.Woba, &batting.RC, &batting.BABIP)

		battings = append(battings, batting)
	}

	return battings
}

// GetCareer 選手情報を取得する
func (Repository *GradesRepository) GetCareer(playerID string) (career data.CAREER) {
	rows, err := Repository.Conn.Query("SELECT * FROM players WHERE player_id = $1", playerID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&career.PlayerID, &career.Name, &career.Position, &career.PitchingAndBatting,
			&career.Height, &career.Weight, &career.Birthday, &career.Draft, &career.Career)
	}

	return career
}

// GetPlayersByTeamIDAndYear チームIDと年から選手一覧を取得する
func (Repository *GradesRepository) GetPlayersByTeamIDAndYear(teamID string, year string) (players []data.PLAYER) {
	rows, err := Repository.Conn.Query("SELECT * FROM team_players WHERE year = $1 AND team_id = $2", year, teamID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var player data.PLAYER
		rows.Scan(&player.Year, &player.TeamID, &player.Team, &player.PlayerID, &player.Name)
		players = append(players, player)
	}

	return players
}

// InsertTeamPlayers 年度ごとの選手一覧をDBに登録する
func (Repository *GradesRepository) InsertTeamPlayers(teamID string, teamName string, players [][]string, year string) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO team_players(year,team_id,team_name,player_id,player_name) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, player := range players {
		playerID := extractionPlayerID(player[0])
		if _, err := stmt.Exec(year, teamID, teamName, playerID, player[1]); err != nil {
			fmt.Println(teamID + ":" + playerID)
			log.Print(err)
		}
	}
}

func extractionPlayerID(url string) string {
	return strings.Replace(strings.Replace(url, "/bis/players/", "", 1), ".html", "", 1)
}

// ExtractionCareers 引数で受け取ったCAREERリストから重複選手を除外する
func (Repository *GradesRepository) ExtractionCareers(careers *[]data.CAREER) {
	rows, err := Repository.Conn.Query("SELECT * FROM players")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var selectCareer data.CAREER
		rows.Scan(&selectCareer.PlayerID, &selectCareer.Name, &selectCareer.Position,
			&selectCareer.PitchingAndBatting, &selectCareer.Height, &selectCareer.Weight,
			&selectCareer.Birthday, &selectCareer.Draft, &selectCareer.Career)
		for index, career := range *careers {
			if career.PlayerID == selectCareer.PlayerID {
				*careers = unset(*careers, index)
			}
		}
	}
}

func unset(s []data.CAREER, i int) []data.CAREER {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// InsertCareers 引数で受け取った CAREER をDBへ登録する
func (Repository *GradesRepository) InsertCareers(careers []data.CAREER) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO players(player_id, name, position, pitching_and_batting, height, weight, birthday, draft, career) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, career := range careers {
		if _, err := stmt.Exec(career.PlayerID, career.Name, career.Position,
			career.PitchingAndBatting, career.Height, career.Weight,
			career.Birthday, career.Draft, career.Career); err != nil {
			fmt.Println(career.PlayerID + ":" + career.Name)
			log.Print(err)
		}
	}
}

// ExtractionPicherGrades 引数で受け取ったPICHERGRADESリストから重複選手を除外する
func (Repository *GradesRepository) ExtractionPicherGrades(picherMap *map[string][]data.PICHERGRADES, teamID string) {
	rows, err := Repository.Conn.Query("SELECT DISTINCT player_id FROM picher_grades where team_id = $1", teamID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var playerID string
		rows.Scan(&playerID)

		for key := range *picherMap {
			if key == playerID {
				delete(*picherMap, key)
			}
		}
	}
}

// InsertPicherGrades 引数で受け取ったPICHERGRADESリストから重複選手を除外する
func (Repository *GradesRepository) InsertPicherGrades(picherMap map[string][]data.PICHERGRADES) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO picher_grades(player_id, year, team_id, team, piched, win, lose, save, hold, hold_point, complete_game, shutout, no_walks, winning_rate, batter, innings_pitched, hit, home_run, base_on_balls, hit_by_ptches, strike_out, wild_pitches, balk, runs_allowed, earned_run, earned_run_average) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	for key, value := range picherMap {
		for _, picher := range value {
			if _, err := stmt.Exec(key, picher.Year, picher.TeamID, picher.Team, picher.Piched, picher.Win, picher.Lose, picher.Save, picher.Hold, picher.HoldPoint, picher.CompleteGame, picher.Shutout, picher.NoWalks, picher.WinningRate, picher.Batter, picher.InningsPitched, picher.Hit, picher.HomeRun, picher.BaseOnBalls, picher.HitByPitches, picher.StrikeOut, picher.WildPitches, picher.Balk, picher.RunsAllowed, picher.EarnedRun, picher.EarnedRunAverage); err != nil {
				fmt.Println(key + ":" + picher.Year)
				log.Print(err)
			}
		}
	}
}

// ExtractionBatterGrades 引数で受け取ったBATTERGRADESリストから重複選手を除外する
func (Repository *GradesRepository) ExtractionBatterGrades(batterMap *map[string][]data.BATTERGRADES, teamID string) {
	rows, err := Repository.Conn.Query("SELECT DISTINCT player_id FROM batter_grades where team_id = $1", teamID)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var playerID string
		rows.Scan(&playerID)

		for key := range *batterMap {
			if key == playerID {
				delete(*batterMap, key)
			}
		}
	}

	rows.Close()
}

// InsertBatterGrades 引数で受け取ったBATTERGRADESをDBに登録する
func (Repository *GradesRepository) InsertBatterGrades(playerID string, batterGrades data.BATTERGRADES) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO batter_grades(player_id, year, team_id, team, games, plate_appearance, at_bat, score, hit, single, double, triple, home_run, base_hit, runs_batted_in, stolen_base, caught_stealing, sacrifice_hits, sacrifice_flies, base_on_balls, hit_by_pitches, strike_out, grounded_into_double_play, batting_average, slugging_percentage, on_base_percentage, w_oba, rc) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	// 加重出塁率の計算に必要なconfigファイルを読み込む
	if _, err := stmt.Exec(playerID, batterGrades.Year, batterGrades.TeamID, batterGrades.Team, batterGrades.Games, batterGrades.PlateAppearance, batterGrades.AtBat, batterGrades.Score, batterGrades.Hit, batterGrades.Single, batterGrades.Double, batterGrades.Triple, batterGrades.HomeRun, batterGrades.BaseHit, batterGrades.RunsBattedIn, batterGrades.StolenBase, batterGrades.CaughtStealing, batterGrades.SacrificeHits, batterGrades.SacrificeFlies, batterGrades.BaseOnBalls, batterGrades.HitByPitches, batterGrades.StrikeOut, batterGrades.GroundedIntoDoublePlay, batterGrades.BattingAverage, batterGrades.SluggingPercentage, batterGrades.OnBasePercentage, batterGrades.Woba, batterGrades.RC); err != nil {
		fmt.Println(playerID + ":" + batterGrades.Year)
		log.Print(err)
	}
}
