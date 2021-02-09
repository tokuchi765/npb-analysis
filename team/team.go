package team

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	teamData "github.com/tokuchi765/npb-analysis/entity/team"
)

// GetTeamBatting 引数で受け取った年に紐づくチーム打撃成績を取得します。
func GetTeamBatting(years []int, db *sql.DB) (teamBattingMap map[string][]teamData.TeamBatting) {
	teamBattingMap = make(map[string][]teamData.TeamBatting)
	for _, year := range years {
		strYear := strconv.Itoa(year)
		rows, err := db.Query("SELECT * FROM team_batting where year = $1", strYear)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}

		var teamBattins []teamData.TeamBatting
		for rows.Next() {
			var teamBatting teamData.TeamBatting
			rows.Scan(&teamBatting.TeamID, &teamBatting.Year, &teamBatting.BattingAverage, &teamBatting.Games, &teamBatting.PlateAppearance,
				&teamBatting.AtBat, &teamBatting.Score, &teamBatting.Hit, &teamBatting.Double, &teamBatting.Triple, &teamBatting.HomeRun,
				&teamBatting.BaseHit, &teamBatting.RunsBattedIn, &teamBatting.StolenBase, &teamBatting.CaughtStealing, &teamBatting.SacrificeHits,
				&teamBatting.SacrificeFlies, &teamBatting.BaseOnBalls, &teamBatting.IntentionalWalk, &teamBatting.HitByPitches, &teamBatting.StrikeOut,
				&teamBatting.GroundedIntoDoublePlay, &teamBatting.SluggingPercentage, &teamBatting.OnBasePercentage,
			)
			teamBattins = append(teamBattins, teamBatting)
		}

		teamBattingMap[strYear] = teamBattins
	}

	return teamBattingMap
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func GetTeamStats(years []int, db *sql.DB) (teamStatsMap map[string][]teamData.TeamLeagueStats) {
	teamStatsMap = make(map[string][]teamData.TeamLeagueStats)
	for _, year := range years {
		strYear := strconv.Itoa(year)
		rows, err := db.Query("SELECT * FROM team_season_stats where year = $1", strYear)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}

		var teamStatses []teamData.TeamLeagueStats
		for rows.Next() {
			var teamStats teamData.TeamLeagueStats
			rows.Scan(&teamStats.TeamID, &teamStats.Year, &teamStats.Games, &teamStats.Win, &teamStats.Lose, &teamStats.Draw,
				&teamStats.WinningRate, &teamStats.ExchangeWin, &teamStats.ExchangeLose, &teamStats.ExchangeDraw,
				&teamStats.HomeWin, &teamStats.HomeLose, &teamStats.HomeDraw,
				&teamStats.LoadWin, &teamStats.LoadLose, &teamStats.LoadDraw)

			teamStatses = append(teamStatses, teamStats)
		}

		teamStatsMap[strYear] = teamStatses
	}
	return teamStatsMap
}

// InsertSeasonLeagueStats チームごとのシーズン成績をDBに登録する
func InsertSeasonLeagueStats(csvPath string, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO team_season_stats(team_id, year, games, win, lose, draw, winning_rate, exchange_win, exchange_lose, exchange_draw, home_win, home_lose, home_draw, load_win, load_lose, load_draw) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	years := makeRange(2005, 2020)
	for _, year := range years {
		cTeamLeagueStats, _ := readTeamLeagueStats(csvPath, "c", strconv.Itoa(year))
		pTeamLeagueStats, _ := readTeamLeagueStats(csvPath, "p", strconv.Itoa(year))

		// DBに登録する
		insertTeamLeagueStats(append(cTeamLeagueStats, pTeamLeagueStats...), stmt)
	}
}

// InsertSeasonMatchResults 各チームの対戦成績をDBに登録する
func InsertSeasonMatchResults(csvPath string, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO team_match_results(team_id, year, competitive_team_id, vs_type, win, lose, draw) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	years := makeRange(2005, 2020)
	for _, year := range years {
		_, cTeamMatchResults := readTeamLeagueStats(csvPath, "c", strconv.Itoa(year))
		_, pTeamMatchResults := readTeamLeagueStats(csvPath, "p", strconv.Itoa(year))

		if year != 2020 {
			cTeamExchangeMatchResults := readTeamExchangeStats(csvPath, "c", strconv.Itoa(year))
			cTeamMatchResults = append(cTeamMatchResults, cTeamExchangeMatchResults...)
			pTeamExchangeMatchResults := readTeamExchangeStats(csvPath, "p", strconv.Itoa(year))
			pTeamMatchResults = append(pTeamMatchResults, pTeamExchangeMatchResults...)
		}

		// DBに登録する
		insertMatchResults(append(cTeamMatchResults, pTeamMatchResults...), stmt)
	}
}

// InsertTeamPitchings チーム投手成績をDBに登録する
func InsertTeamPitchings(csvPath string, league string, db *sql.DB) {
	years := makeRange(2005, 2020)
	for _, year := range years {
		teamPitching := readTeamPitching(csvPath, league, strconv.Itoa(year))
		insertTeamPitching(teamPitching, db)
	}
}

// InsertTeamBattings チーム打撃成績をDBに登録する
func InsertTeamBattings(csvPath string, league string, db *sql.DB) {
	years := makeRange(2005, 2020)
	for _, year := range years {
		teamBatting := readTeamBatting(csvPath, league, strconv.Itoa(year))
		insertTeamBatting(teamBatting, db)
	}
}

func insertMatchResults(teamMatchResults []teamData.TeamMatchResults, stmt *sql.Stmt) {
	for _, result := range teamMatchResults {
		if _, err := stmt.Exec(result.TeamID, result.Year, result.CompetitiveTeamID, result.VsType, result.Win, result.Lose, result.Draw); err != nil {
			fmt.Println(result.TeamID + ":" + result.Year)
			log.Print(err)
		}
	}
}

func insertTeamLeagueStats(teamLeagueStats []teamData.TeamLeagueStats, stmt *sql.Stmt) {
	for _, stats := range teamLeagueStats {
		if _, err := stmt.Exec(stats.TeamID, stats.Year, stats.Games, stats.Win, stats.Lose, stats.Draw, stats.WinningRate, stats.ExchangeWin, stats.ExchangeLose, stats.ExchangeDraw, stats.HomeWin, stats.HomeLose, stats.HomeDraw, stats.LoadWin, stats.LoadLose, stats.LoadDraw); err != nil {
			fmt.Println(stats.TeamID + ":" + stats.Year)
			log.Print(err)
		}
	}
}

func readTeamExchangeStats(csvPath string, league string, year string) (teamExchangeMatchResults []teamData.TeamMatchResults) {
	path := csvPath + "/teams/stats/season/" + league + "/" + year + "_exchange_stats.csv"
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを読み込み、対戦成績のインデックスを取得する
	header, _ := reader.Read()
	indexMap := getIndex(header)

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		teamExchangeMatchResults = append(teamExchangeMatchResults, setTeamMatchResults(line, year, indexMap, "exchange")...)
	}

	return teamExchangeMatchResults
}

func readTeamLeagueStats(csvPath string, league string, year string) (teamLeagueStats []teamData.TeamLeagueStats, teamMatchResults []teamData.TeamMatchResults) {
	path := csvPath + "/teams/stats/season/" + league + "/" + year + "_league_stats.csv"
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
	}
	// 	終わったらファイルを閉じる
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダーを読み込み、対戦成績のインデックスを取得する
	header, _ := reader.Read()
	indexMap := getIndex(header)

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		teamLeagueStats = append(teamLeagueStats, setTeamLeagueStats(line, year, indexMap))
		teamMatchResults = append(teamMatchResults, setTeamMatchResults(line, year, indexMap, "league")...)
	}

	return teamLeagueStats, teamMatchResults
}

func setTeamMatchResults(line []string, year string, indexMap map[string]int, vsType string) (teamMatchResultsList []teamData.TeamMatchResults) {
	teamID := GetTeamID(line[1])

	// 各チーム戦績
	var giantsResults teamData.TeamMatchResults    // 巨人
	var baystarsResults teamData.TeamMatchResults  // 横浜
	var tigersResults teamData.TeamMatchResults    // 阪神
	var carpResults teamData.TeamMatchResults      // 広島
	var dragonsResults teamData.TeamMatchResults   // 中日
	var swallowsResults teamData.TeamMatchResults  // ヤクルト
	var lionsResults teamData.TeamMatchResults     // 西武
	var hawksResults teamData.TeamMatchResults     // ホークス
	var eaglesResults teamData.TeamMatchResults    // イーグルス
	var marinesResults teamData.TeamMatchResults   // マリーンズ
	var fightersResults teamData.TeamMatchResults  // ファイターズ
	var buffaloesResults teamData.TeamMatchResults // バファローズ

	for key, value := range indexMap {
		competitiveTeamID := getTeamIDByLeagueStatsIndexKey(key)
		result, _ := strconv.Atoi(line[value])

		// 巨人の対戦成績
		if competitiveTeamID == "01" {
			setResults(&giantsResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// 横浜の対戦成績
		if competitiveTeamID == "02" {
			setResults(&baystarsResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// 阪神の対戦成績
		if competitiveTeamID == "03" {
			setResults(&tigersResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// 広島の対戦成績
		if competitiveTeamID == "04" {
			setResults(&carpResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// 中日の対戦成績
		if competitiveTeamID == "05" {
			setResults(&dragonsResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// ヤクルトの対戦成績
		if competitiveTeamID == "06" {
			setResults(&swallowsResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// 西武の対戦成績
		if competitiveTeamID == "07" {
			setResults(&lionsResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// ホークスの対戦成績
		if competitiveTeamID == "08" {
			setResults(&hawksResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// イーグルスの対戦成績
		if competitiveTeamID == "09" {
			setResults(&eaglesResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// マリーンズの対戦成績
		if competitiveTeamID == "10" {
			setResults(&marinesResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// ファイターズの対戦成績
		if competitiveTeamID == "11" {
			setResults(&fightersResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
		// バファローズの対戦成績
		if competitiveTeamID == "12" {
			setResults(&buffaloesResults, teamID, year, competitiveTeamID, key, result, vsType)
		}
	}

	// 各チーム戦績をリストに挿入
	results := []teamData.TeamMatchResults{giantsResults, baystarsResults, tigersResults, carpResults, dragonsResults, swallowsResults,
		lionsResults, hawksResults, eaglesResults, marinesResults, fightersResults, buffaloesResults}
	for _, result := range results {
		if result.TeamID != "" {
			teamMatchResultsList = append(teamMatchResultsList, result)
		}
	}

	return teamMatchResultsList
}

func setResults(teamMatchResults *teamData.TeamMatchResults, teamID string, year string, competitiveTeamID string, index string, result int, league string) {
	teamMatchResults.TeamID = teamID
	teamMatchResults.Year = year
	teamMatchResults.CompetitiveTeamID = competitiveTeamID
	teamMatchResults.VsType = league
	if strings.Index(index, "win") >= 0 {
		teamMatchResults.Win = result
	} else if strings.Index(index, "lose") >= 0 {
		teamMatchResults.Lose = result
	} else if strings.Index(index, "draw") >= 0 {
		teamMatchResults.Draw = result
	}
}

func getTeamIDByLeagueStatsIndexKey(indexKey string) (teamID string) {
	idDatas := map[string][]string{
		"01": {"Giants_win", "Giants_lose", "Giants_draw"},
		"02": {"Baystars_win", "Baystars_lose", "Baystars_draw"},
		"03": {"Tigers_win", "Tigers_lose", "Tigers_draw"},
		"04": {"Carp_win", "Carp_lose", "Carp_draw"},
		"05": {"Dragons_win", "Dragons_lose", "Dragons_draw"},
		"06": {"Swallows_win", "Swallows_lose", "Swallows_draw"},
		"07": {"Lions_win", "Lions_lose", "Lions_draw"},
		"08": {"Hawks_win", "Hawks_lose", "Hawks_draw"},
		"09": {"Eagles_win", "Eagles_lose", "Eagles_draw"},
		"10": {"Marines_win", "Marines_lose", "Marines_draw"},
		"11": {"Fighters_win", "Fighters_lose", "Fighters_draw"},
		"12": {"Buffaloes_win", "Buffaloes_lose", "Buffaloes_draw"},
	}
	for key, idData := range idDatas {
		for _, id := range idData {
			if id == indexKey {
				return key
			}
		}
	}
	return ""
}

func setTeamLeagueStats(line []string, year string, indexMap map[string]int) (teamPitching teamData.TeamLeagueStats) {
	teamPitching.TeamID = GetTeamID(line[1])
	teamPitching.Year = year
	teamPitching.Games, _ = strconv.Atoi(line[2])
	teamPitching.Win, _ = strconv.Atoi(line[3])
	teamPitching.Lose, _ = strconv.Atoi(line[4])
	teamPitching.Draw, _ = strconv.Atoi(line[5])
	teamPitching.WinningRate, _ = strconv.ParseFloat(line[6], 64)
	teamPitching.ExchangeWin, _ = strconv.Atoi(line[indexMap["exchange_win"]])
	teamPitching.ExchangeLose, _ = strconv.Atoi(line[indexMap["exchange_lose"]])
	teamPitching.ExchangeDraw, _ = strconv.Atoi(line[indexMap["exchange_draw"]])
	teamPitching.HomeWin, _ = strconv.Atoi(line[indexMap["home_win"]])
	teamPitching.HomeLose, _ = strconv.Atoi(line[indexMap["home_lose"]])
	teamPitching.HomeDraw, _ = strconv.Atoi(line[indexMap["home_draw"]])
	teamPitching.LoadWin, _ = strconv.Atoi(line[indexMap["load_win"]])
	teamPitching.LoadLose, _ = strconv.Atoi(line[indexMap["load_lose"]])
	teamPitching.LoadDraw, _ = strconv.Atoi(line[indexMap["load_draw"]])
	return teamPitching
}

func getIndex(lines []string) (indexMap map[string]int) {
	indexMap = make(map[string]int)
	headerNameMap := []map[string][]string{
		{"exchange_win": {"交流戦(勝)"}}, {"exchange_lose": {"交流戦(負)"}}, {"exchange_draw": {"交流戦(引)"}},
		{"home_win": {"ホ｜ム(勝)"}}, {"home_lose": {"ホ｜ム(負)"}}, {"home_draw": {"ホ｜ム(引)"}},
		{"load_win": {"ロ｜ド(勝)"}}, {"load_lose": {"ロ｜ド(負)"}}, {"load_draw": {"ロ｜ド(引)"}},
		{"Tigers_win": {"対神(勝)"}}, {"Tigers_lose": {"対神(負)"}}, {"Tigers_draw": {"対神(引)"}},
		{"Dragons_win": {"対中(勝)"}}, {"Dragons_lose": {"対中(負)"}}, {"Dragons_draw": {"対中(引)"}},
		{"Baystars_win": {"対横(勝)", "対デ(勝)"}}, {"Baystars_lose": {"対横(負)", "対デ(負)"}}, {"Baystars_draw": {"対横(引)", "対デ(引)"}},
		{"Swallows_win": {"対ヤ(勝)"}}, {"Swallows_lose": {"対ヤ(負)"}}, {"Swallows_draw": {"対ヤ(引)"}},
		{"Giants_win": {"対巨(勝)"}}, {"Giants_lose": {"対巨(負)"}}, {"Giants_draw": {"対巨(引)"}},
		{"Carp_win": {"対広(勝)"}}, {"Carp_lose": {"対広(負)"}}, {"Carp_draw": {"対広(引)"}},
		{"Lions_win": {"対西(勝)"}}, {"Lions_lose": {"対西(負)"}}, {"Lions_draw": {"対西(引)"}},
		{"Hawks_win": {"対ソ(勝)"}}, {"Hawks_lose": {"対ソ(負)"}}, {"Hawks_draw": {"対ソ(引)"}},
		{"Eagles_win": {"対楽(勝)"}}, {"Eagles_lose": {"対楽(負)"}}, {"Eagles_draw": {"対楽(引)"}},
		{"Marines_win": {"対ロ(勝)"}}, {"Marines_lose": {"対ロ(負)"}}, {"Marines_draw": {"対ロ(引)"}},
		{"Fighters_win": {"対日(勝)"}}, {"Fighters_lose": {"対日(負)"}}, {"Fighters_draw": {"対日(引)"}},
		{"Buffaloes_win": {"対オ(勝)"}}, {"Buffaloes_lose": {"対オ(負)"}}, {"Buffaloes_draw": {"対オ(引)"}},
	}
	for index, line := range lines {
		for _, headerNames := range headerNameMap {
			for key, headerName := range headerNames {
				if headerName[0] == line {
					indexMap[key] = index
				}
			}
		}
	}

	return indexMap
}

func insertTeamPitching(teamPitching []teamData.TeamPitching, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO team_pitching(team_id, year, earned_run_average, games, win, lose, save, hold, hold_point, complete_game, shutout, no_walks, winning_rate, batter, innings_pitched, hit, home_run, base_on_balls, intentional_walk, hit_by_ptches, strike_out, wild_pitches, balk, runs_allowed, earned_run) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, pitching := range teamPitching {
		if _, err := stmt.Exec(pitching.TeamID, pitching.Year, pitching.EarnedRunAverage, pitching.Games, pitching.Win, pitching.Lose, pitching.Save, pitching.Hold, pitching.HoldPoint, pitching.CompleteGame, pitching.Shutout, pitching.NoWalks, pitching.WinningRate, pitching.Batter, pitching.InningsPitched, pitching.Hit, pitching.HomeRun, pitching.BaseOnBalls, pitching.IntentionalWalk, pitching.HitByPitches, pitching.StrikeOut, pitching.WildPitches, pitching.Balk, pitching.RunsAllowed, pitching.EarnedRun); err != nil {
			fmt.Println(pitching.TeamID + ":" + pitching.Year)
			log.Print(err)
		}
	}
}

func readTeamPitching(csvPath string, league string, year string) (teamPitching []teamData.TeamPitching) {
	path := csvPath + "/teams/stats/team_pitching/" + league + "/" + year + ".csv"
	file, err := os.Open(path)

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
		teamPitching = append(teamPitching, setTeamPitching(line, year))
	}

	return teamPitching
}

func setTeamPitching(line []string, year string) (teamPitching teamData.TeamPitching) {
	teamPitching.TeamID = GetTeamID(line[1])
	teamPitching.Year = year
	teamPitching.EarnedRunAverage, _ = strconv.ParseFloat(line[2], 64)
	teamPitching.Games, _ = strconv.Atoi(line[3])
	teamPitching.Win, _ = strconv.Atoi(line[4])
	teamPitching.Lose, _ = strconv.Atoi(line[5])
	teamPitching.Save, _ = strconv.Atoi(line[6])
	teamPitching.Hold, _ = strconv.Atoi(line[7])
	teamPitching.HoldPoint, _ = strconv.Atoi(line[8])
	teamPitching.CompleteGame, _ = strconv.Atoi(line[9])
	teamPitching.Shutout, _ = strconv.Atoi(line[10])
	teamPitching.NoWalks, _ = strconv.Atoi(line[11])
	teamPitching.WinningRate, _ = strconv.ParseFloat(line[12], 64)
	teamPitching.Batter, _ = strconv.Atoi(line[13])
	teamPitching.InningsPitched, _ = strconv.Atoi(line[14])
	teamPitching.Hit, _ = strconv.Atoi(line[15])
	teamPitching.HomeRun, _ = strconv.Atoi(line[16])
	teamPitching.BaseOnBalls, _ = strconv.Atoi(line[17])
	teamPitching.IntentionalWalk, _ = strconv.Atoi(line[18])
	teamPitching.HitByPitches, _ = strconv.Atoi(line[19])
	teamPitching.StrikeOut, _ = strconv.Atoi(line[20])
	teamPitching.WildPitches, _ = strconv.Atoi(line[21])
	teamPitching.Balk, _ = strconv.Atoi(line[22])
	teamPitching.RunsAllowed, _ = strconv.Atoi(line[23])
	teamPitching.EarnedRun, _ = strconv.Atoi(line[24])
	return teamPitching
}

func insertTeamBatting(teamBatting []teamData.TeamBatting, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO team_batting(team_id, year, batting_average, games, plate_appearance, at_bat, score, hit, double, triple, home_run, base_hit, runs_batted_in, stolen_base, caught_stealing, sacrifice_hits, sacrifice_flies, base_on_balls, intentional_walk, hit_by_pitches, strike_out, grounded_into_double_play, slugging_percentage, on_base_percentage) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, batting := range teamBatting {
		if _, err := stmt.Exec(batting.TeamID, batting.Year, batting.BattingAverage, batting.Games, batting.PlateAppearance, batting.AtBat, batting.Score, batting.Hit, batting.Double, batting.Triple, batting.HomeRun, batting.BaseHit, batting.RunsBattedIn, batting.StolenBase, batting.CaughtStealing, batting.SacrificeHits, batting.SacrificeFlies, batting.BaseOnBalls, batting.IntentionalWalk, batting.HitByPitches, batting.StrikeOut, batting.GroundedIntoDoublePlay, batting.SluggingPercentage, batting.OnBasePercentage); err != nil {
			fmt.Println(batting.TeamID + ":" + batting.Year)
			log.Print(err)
		}
	}
}

func readTeamBatting(csvPath string, league string, year string) (teamBatting []teamData.TeamBatting) {
	path := csvPath + "/teams/stats/team_batting/" + league + "/" + year + ".csv"
	file, err := os.Open(path)

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
		teamBatting = append(teamBatting, setTeamBatting(line, year))
	}

	return teamBatting
}

func setTeamBatting(line []string, year string) (teamBatting teamData.TeamBatting) {
	teamBatting.TeamID = GetTeamID(line[1])
	teamBatting.Year = year
	teamBatting.BattingAverage, _ = strconv.ParseFloat(line[2], 64)
	teamBatting.Games, _ = strconv.Atoi(line[3])
	teamBatting.PlateAppearance, _ = strconv.Atoi(line[4])
	teamBatting.AtBat, _ = strconv.Atoi(line[5])
	teamBatting.Score, _ = strconv.Atoi(line[6])
	teamBatting.Hit, _ = strconv.Atoi(line[7])
	teamBatting.Double, _ = strconv.Atoi(line[8])
	teamBatting.Triple, _ = strconv.Atoi(line[9])
	teamBatting.HomeRun, _ = strconv.Atoi(line[10])
	teamBatting.BaseHit, _ = strconv.Atoi(line[11])
	teamBatting.RunsBattedIn, _ = strconv.Atoi(line[12])
	teamBatting.StolenBase, _ = strconv.Atoi(line[13])
	teamBatting.CaughtStealing, _ = strconv.Atoi(line[14])
	teamBatting.SacrificeHits, _ = strconv.Atoi(line[15])
	teamBatting.SacrificeFlies, _ = strconv.Atoi(line[16])
	teamBatting.BaseOnBalls, _ = strconv.Atoi(line[17])
	teamBatting.IntentionalWalk, _ = strconv.Atoi(line[18])
	teamBatting.HitByPitches, _ = strconv.Atoi(line[19])
	teamBatting.StrikeOut, _ = strconv.Atoi(line[20])
	teamBatting.GroundedIntoDoublePlay, _ = strconv.Atoi(line[21])
	teamBatting.SluggingPercentage, _ = strconv.ParseFloat(line[22], 64)
	teamBatting.OnBasePercentage, _ = strconv.ParseFloat(line[23], 64)
	return teamBatting
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
