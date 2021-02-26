package grades

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	data "github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/team"
)

// GetPitching 個人投手成績一覧を取得する
func GetPitching(playerID string, db *sql.DB) (pitchings []data.PICHERGRADES) {
	rows, err := db.Query("SELECT * FROM picher_grades WHERE player_id = $1", playerID)

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
			&pitching.EarnedRun, &pitching.EarnedRunAverage)

		pitchings = append(pitchings, pitching)
	}

	return pitchings
}

// GetBatting 個人打撃成績一覧を取得する
func GetBatting(playerID string, db *sql.DB) (battings []data.BATTERGRADES) {
	rows, err := db.Query("SELECT * FROM batter_grades WHERE player_id = $1", playerID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var playerID string
		var batting data.BATTERGRADES
		rows.Scan(&playerID, &batting.Year, &batting.TeamID, &batting.Team, &batting.Games,
			&batting.PlateAppearance, &batting.AtBat, &batting.Score, &batting.Hit,
			&batting.Double, &batting.Triple, &batting.HomeRun, &batting.BaseHit,
			&batting.RunsBattedIn, &batting.StolenBase, &batting.CaughtStealing, &batting.SacrificeHits,
			&batting.SacrificeFlies, &batting.BaseOnBalls, &batting.HitByPitches, &batting.StrikeOut,
			&batting.GroundedIntoDoublePlay, &batting.BattingAverage, &batting.SluggingPercentage, &batting.OnBasePercentage)

		battings = append(battings, batting)
	}

	return battings
}

func GetCareer(playerID string, db *sql.DB) (career data.CAREER) {
	rows, err := db.Query("SELECT * FROM players WHERE player_id = $1", playerID)

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

func GetPlayersByTeamIDandYear(teamID string, year string, db *sql.DB) (players []data.PLAYER) {
	rows, err := db.Query("SELECT * FROM team_players WHERE year = $1 AND team_id = $2", year, teamID)

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
func InsertTeamPlayers(initial string, players [][]string, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO team_players(year,team_id,team_name,player_id,player_name) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	teamID := team.GetTeamID(initial)
	teamName := getTeamName(teamID, db)
	for _, player := range players {
		playerID := extractionPlayerID(player[0])
		if _, err := stmt.Exec("2020", teamID, teamName, playerID, player[1]); err != nil {
			fmt.Println(teamID + ":" + playerID)
			log.Print(err)
		}
	}
}

func getTeamName(teamID string, db *sql.DB) (teamName string) {
	rows, err := db.Query("SELECT team_name FROM team_name WHERE team_name_id = $1", teamID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&teamName)
	}

	return teamName
}

// GetPlayers 引数で受け取った x_players.csv ファイルを読み取って、配列にして返す
func GetPlayers(url string) (players [][]string) {
	// バイト列を読み込む
	file, err := os.Open(url)
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを読み飛ばす
	_, _ = reader.Read()

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		players = append(players, []string{line[1], line[2]})
	}

	return players
}

// ReadCareers 引数で受け取った選手リストをもとに、経歴をまとめたデータクラスのリストを返す
func ReadCareers(playersPath string, initial string, players [][]string) (careerList []data.CAREER) {
	for _, player := range players {
		id := extractionPlayerID(player[0])
		url := playersPath + initial + "/careers/" + id + "_" + player[1] + "_career.csv"
		if exists(url) {
			career := readCareer(url)
			careerList = append(careerList, career)
		}
	}
	return careerList
}

func extractionPlayerID(url string) string {
	return strings.Replace(strings.Replace(url, "/bis/players/", "", 1), ".html", "", 1)
}

// ExtractionCareers 引数で受け取ったCAREERリストから重複選手を除外する
func ExtractionCareers(careers *[]data.CAREER, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM players")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var selectCareer data.CAREER
		rows.Scan(&selectCareer.PlayerID, &selectCareer.Name, &selectCareer.Position, &selectCareer.PitchingAndBatting, &selectCareer.Height, &selectCareer.Weight, &selectCareer.Birthday, &selectCareer.Draft, &selectCareer.Career)
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
func InsertCareers(careers []data.CAREER, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO players(player_id, name, position, pitching_and_batting, height, weight, birthday, draft, career) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, career := range careers {
		if _, err := stmt.Exec(career.PlayerID, career.Name, career.Position, career.PitchingAndBatting, career.Height, career.Weight, career.Birthday, career.Draft, career.Career); err != nil {
			fmt.Println(career.PlayerID + ":" + career.Name)
			log.Print(err)
		}
	}
}

// ReadGradesMap 引数のplayersに設定されている選手成績を読み込み、Mapにして返す
func ReadGradesMap(playersPath string, initial string, players [][]string) (picherMap map[string][]data.PICHERGRADES, batterMap map[string][]data.BATTERGRADES) {
	picherMap = make(map[string][]data.PICHERGRADES)
	batterMap = make(map[string][]data.BATTERGRADES)
	for _, player := range players {
		id := strings.Replace(strings.Replace(player[0], "/bis/players/", "", 1), ".html", "", 1)
		path := playersPath + initial + "/grades/" + id + "_" + player[1] + "_grades.csv"

		if exists(path) {
			picherGrades, batterGrades := readGrades(path)
			if picherGrades != nil {
				picherMap[id] = picherGrades
			} else {
				batterMap[id] = batterGrades
			}
		}
	}
	return picherMap, batterMap
}

// ExtractionPicherGrades 引数で受け取ったPICHERGRADESリストから重複選手を除外する
func ExtractionPicherGrades(picherMap *map[string][]data.PICHERGRADES, teamID string, db *sql.DB) {
	rows, err := db.Query("SELECT DISTINCT player_id FROM picher_grades where team_id = $1", teamID)

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
func InsertPicherGrades(picherMap map[string][]data.PICHERGRADES, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO picher_grades(player_id, year, team_id, team, piched, win, lose, save, hold, hold_point, complete_game, shutout, no_walks, winning_rate, batter, innings_pitched, hit, home_run, base_on_balls, hit_by_ptches, strike_out, wild_pitches, balk, runs_allowed, earned_run, earned_run_average) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)")
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
func ExtractionBatterGrades(batterMap *map[string][]data.BATTERGRADES, teamID string, db *sql.DB) {
	rows, err := db.Query("SELECT DISTINCT player_id FROM batter_grades where team_id = $1", teamID)

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
func InsertBatterGrades(batterMap map[string][]data.BATTERGRADES, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO batter_grades(player_id, year, team_id, team, games, plate_appearance, at_bat, score, hit, double, triple, home_run, base_hit, runs_batted_in, stolen_base, caught_stealing, sacrifice_hits, sacrifice_flies, base_on_balls, hit_by_pitches, strike_out, grounded_into_double_play, batting_average, slugging_percentage, on_base_percentage) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	for key, value := range batterMap {
		for _, batter := range value {
			if _, err := stmt.Exec(key, batter.Year, batter.TeamID, batter.Team, batter.Games, batter.PlateAppearance, batter.AtBat, batter.Score, batter.Hit, batter.Double, batter.Triple, batter.HomeRun, batter.BaseHit, batter.RunsBattedIn, batter.StolenBase, batter.CaughtStealing, batter.SacrificeHits, batter.SacrificeFlies, batter.BaseOnBalls, batter.HitByPitches, batter.StrikeOut, batter.GroundedIntoDoublePlay, batter.BattingAverage, batter.SluggingPercentage, batter.OnBasePercentage); err != nil {
				fmt.Println(key + ":" + batter.Year)
				log.Print(err)
			}
		}
	}
}

func readGrades(path string) (picherGradesList []data.PICHERGRADES, batterGradesList []data.BATTERGRADES) {
	// バイト列を読み込む
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを取得
	header, err := reader.Read()
	if err != nil {
		log.Print(err)
	}

	for {
		line, err := reader.Read()

		if err != nil {
			break
		}

		if header[3] == "登板" {
			picherGradesList = append(picherGradesList, setPicherGrades(line))
		} else {
			batterGradesList = append(batterGradesList, setBatterGrades(line))
		}
	}
	return picherGradesList, batterGradesList
}

func setPicherGrades(line []string) (grades data.PICHERGRADES) {
	grades.Year = line[1]
	grades.TeamID = team.GetTeamID(line[2])
	grades.Team = line[2]
	grades.Piched, _ = strconv.ParseFloat(line[3], 64)
	grades.Win, _ = strconv.ParseFloat(line[4], 64)
	grades.Lose, _ = strconv.ParseFloat(line[5], 64)
	grades.Save, _ = strconv.ParseFloat(line[6], 64)
	grades.Hold, _ = strconv.ParseFloat(line[7], 64)
	grades.HoldPoint, _ = strconv.ParseFloat(line[8], 64)
	grades.CompleteGame, _ = strconv.ParseFloat(line[9], 64)
	grades.Shutout, _ = strconv.ParseFloat(line[10], 64)
	grades.NoWalks, _ = strconv.ParseFloat(line[11], 64)
	grades.WinningRate, _ = strconv.ParseFloat(line[12], 64)
	grades.Batter, _ = strconv.ParseFloat(line[13], 64)
	grades.InningsPitched, _ = strconv.ParseFloat(line[14], 64)
	grades.Hit, _ = strconv.ParseFloat(line[15], 64)
	grades.HomeRun, _ = strconv.ParseFloat(line[16], 64)
	grades.BaseOnBalls, _ = strconv.ParseFloat(line[17], 64)
	grades.HitByPitches, _ = strconv.ParseFloat(line[18], 64)
	grades.StrikeOut, _ = strconv.ParseFloat(line[19], 64)
	grades.WildPitches, _ = strconv.ParseFloat(line[20], 64)
	grades.Balk, _ = strconv.ParseFloat(line[21], 64)
	grades.RunsAllowed, _ = strconv.ParseFloat(line[22], 64)
	grades.EarnedRun, _ = strconv.ParseFloat(line[23], 64)
	grades.EarnedRunAverage, _ = strconv.ParseFloat(line[24], 64)

	return grades
}

func setBatterGrades(line []string) (grades data.BATTERGRADES) {
	grades.Year = line[1]
	grades.TeamID = team.GetTeamID(line[2])
	grades.Team = line[2]
	grades.Games, _ = strconv.Atoi(line[3])
	grades.PlateAppearance, _ = strconv.Atoi(line[4])
	grades.AtBat, _ = strconv.Atoi(line[5])
	grades.Score, _ = strconv.Atoi(line[6])
	grades.Hit, _ = strconv.Atoi(line[7])
	grades.Double, _ = strconv.Atoi(line[8])
	grades.Triple, _ = strconv.Atoi(line[9])
	grades.HomeRun, _ = strconv.Atoi(line[10])
	grades.BaseHit, _ = strconv.Atoi(line[11])
	grades.RunsBattedIn, _ = strconv.Atoi(line[12])
	grades.StolenBase, _ = strconv.Atoi(line[13])
	grades.CaughtStealing, _ = strconv.Atoi(line[14])
	grades.SacrificeHits, _ = strconv.Atoi(line[15])
	grades.SacrificeFlies, _ = strconv.Atoi(line[16])
	grades.BaseOnBalls, _ = strconv.Atoi(line[17])
	grades.HitByPitches, _ = strconv.Atoi(line[18])
	grades.StrikeOut, _ = strconv.Atoi(line[19])
	grades.GroundedIntoDoublePlay, _ = strconv.Atoi(line[20])
	grades.BattingAverage, _ = strconv.ParseFloat(line[21], 64)
	grades.SluggingPercentage, _ = strconv.ParseFloat(line[22], 64)
	grades.OnBasePercentage, _ = strconv.ParseFloat(line[23], 64)

	return grades
}

func readCareer(path string) (career data.CAREER) {
	// バイト列を読み込む
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)
	var lines []string

	// ヘッダーを取得
	_, err = reader.Read()
	if err != nil {
		log.Print(err)
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		lines = append(lines, line[2])
	}

	return setCareer(lines)
}

func setCareer(line []string) (career data.CAREER) {
	career.PlayerID = line[7]
	career.Name = line[8]
	career.Position = line[0]
	career.PitchingAndBatting = line[1]
	career.Height = line[2]
	career.Weight = line[6]
	career.Birthday = line[3]
	career.Career = line[4]
	career.Draft = line[5]

	return career
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
