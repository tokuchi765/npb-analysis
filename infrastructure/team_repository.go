package infrastructure

import (
	"fmt"
	"log"
	"strconv"

	teamData "github.com/tokuchi765/npb-analysis/entity/team"
)

// TeamRepository チーム成績データアクセスを管理するリポジトリ
type TeamRepository struct {
	SQLHandler
}

// InsertTeamPitchings チーム投手成績をDBに登録する
func (Repository *TeamRepository) InsertTeamPitchings(pitching teamData.TeamPitching) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO team_pitching(team_id, year, earned_run_average, games, win, lose, save, hold, hold_point, complete_game, shutout, no_walks, winning_rate, batter, innings_pitched, hit, home_run, base_on_balls, intentional_walk, hit_by_ptches, strike_out, wild_pitches, balk, runs_allowed, earned_run, babip, strike_out_rate) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(pitching.TeamID, pitching.Year, pitching.EarnedRunAverage, pitching.Games, pitching.Win, pitching.Lose, pitching.Save, pitching.Hold, pitching.HoldPoint, pitching.CompleteGame, pitching.Shutout, pitching.NoWalks, pitching.WinningRate, pitching.Batter, pitching.InningsPitched, pitching.Hit, pitching.HomeRun, pitching.BaseOnBalls, pitching.IntentionalWalk, pitching.HitByPitches, pitching.StrikeOut, pitching.WildPitches, pitching.Balk, pitching.RunsAllowed, pitching.EarnedRun, pitching.BABIP, pitching.StrikeOutRate); err != nil {
		fmt.Println(pitching.TeamID + ":" + pitching.Year)
		log.Print(err)
	}
}

// GetTeamPitchings 引数で受け取った年に紐づくチーム投手成績を取得します。
func (Repository *TeamRepository) GetTeamPitchings(years []int) (teamPitchingMap map[string][]teamData.TeamPitching) {
	teamPitchingMap = make(map[string][]teamData.TeamPitching)

	for _, year := range years {
		strYear := strconv.Itoa(year)
		rows, err := Repository.SQLHandler.Conn.Query("SELECT * FROM team_pitching where year = $1", strYear)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}

		var teamPitchings []teamData.TeamPitching
		for rows.Next() {
			var teamPitching teamData.TeamPitching
			rows.Scan(&teamPitching.TeamID, &teamPitching.Year, &teamPitching.EarnedRunAverage,
				&teamPitching.Games, &teamPitching.Win, &teamPitching.Lose,
				&teamPitching.Save, &teamPitching.Hold, &teamPitching.HoldPoint,
				&teamPitching.CompleteGame, &teamPitching.Shutout, &teamPitching.NoWalks,
				&teamPitching.WinningRate, &teamPitching.Batter, &teamPitching.InningsPitched,
				&teamPitching.Hit, &teamPitching.HomeRun, &teamPitching.BaseOnBalls,
				&teamPitching.IntentionalWalk, &teamPitching.HitByPitches, &teamPitching.StrikeOut,
				&teamPitching.WildPitches, &teamPitching.Balk, &teamPitching.RunsAllowed, &teamPitching.EarnedRun, &teamPitching.BABIP, &teamPitching.StrikeOutRate)
			teamPitchings = append(teamPitchings, teamPitching)
		}
		teamPitchingMap[strYear] = teamPitchings
	}
	return teamPitchingMap
}

// InsertTeamBattings チーム打撃成績をDBに登録する
func (Repository *TeamRepository) InsertTeamBattings(batting teamData.TeamBatting) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO team_batting(team_id, year, batting_average, games, plate_appearance, at_bat, score, hit, double, triple, home_run, base_hit, runs_batted_in, stolen_base, caught_stealing, sacrifice_hits, sacrifice_flies, base_on_balls, intentional_walk, hit_by_pitches, strike_out, grounded_into_double_play, slugging_percentage, on_base_percentage, babip) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(batting.TeamID, batting.Year, batting.BattingAverage, batting.Games, batting.PlateAppearance, batting.AtBat, batting.Score, batting.Hit, batting.Double, batting.Triple, batting.HomeRun, batting.BaseHit, batting.RunsBattedIn, batting.StolenBase, batting.CaughtStealing, batting.SacrificeHits, batting.SacrificeFlies, batting.BaseOnBalls, batting.IntentionalWalk, batting.HitByPitches, batting.StrikeOut, batting.GroundedIntoDoublePlay, batting.SluggingPercentage, batting.OnBasePercentage, batting.BABIP); err != nil {
		fmt.Println(batting.TeamID + ":" + batting.Year)
		log.Print(err)
	}
}

// GetTeamBattings 引数で受け取った年に紐づくチーム打撃成績を取得します。
func (Repository *TeamRepository) GetTeamBattings(years []int) (teamBattingMap map[string][]teamData.TeamBatting) {
	teamBattingMap = make(map[string][]teamData.TeamBatting)
	for _, year := range years {
		strYear := strconv.Itoa(year)
		rows, err := Repository.Conn.Query("SELECT * FROM team_batting where year = $1", strYear)

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
				&teamBatting.GroundedIntoDoublePlay, &teamBatting.SluggingPercentage, &teamBatting.OnBasePercentage, &teamBatting.BABIP,
			)
			teamBattins = append(teamBattins, teamBatting)
		}

		teamBattingMap[strYear] = teamBattins
	}

	return teamBattingMap
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func (Repository *TeamRepository) GetTeamStats(years []int) (teamStatsMap map[string][]teamData.TeamLeagueStats) {
	teamStatsMap = make(map[string][]teamData.TeamLeagueStats)
	for _, year := range years {
		strYear := strconv.Itoa(year)
		rows, err := Repository.SQLHandler.Conn.Query("SELECT * FROM team_season_stats where year = $1", strYear)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}

		var teamStatses []teamData.TeamLeagueStats
		for rows.Next() {
			var teamStats teamData.TeamLeagueStats
			rows.Scan(&teamStats.TeamID, &teamStats.Year, &teamStats.Manager, &teamStats.Games, &teamStats.Win, &teamStats.Lose, &teamStats.Draw,
				&teamStats.WinningRate, &teamStats.ExchangeWin, &teamStats.ExchangeLose, &teamStats.ExchangeDraw,
				&teamStats.HomeWin, &teamStats.HomeLose, &teamStats.HomeDraw,
				&teamStats.LoadWin, &teamStats.LoadLose, &teamStats.LoadDraw,
				&teamStats.PythagoreanExpectation)

			teamStatses = append(teamStatses, teamStats)
		}

		teamStatsMap[strYear] = teamStatses
	}
	return teamStatsMap
}

// InsertPythagoreanExpectation ピタゴラス勝率をDBに登録します。
func (Repository *TeamRepository) InsertPythagoreanExpectation(teamBattings []teamData.TeamBatting, teamPitchings []teamData.TeamPitching) {
	stmt, err := Repository.Conn.Prepare("UPDATE team_season_stats SET pythagorean_expectation = $1 WHERE team_id = $2 AND year = $3")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	for _, teamBatting := range teamBattings {
		for _, teamPitching := range teamPitchings {
			if teamBatting.TeamID == teamPitching.TeamID {
				pythagoreanExpectation := calcPythagoreanExpectation(teamBatting.Score, teamPitching.RunsAllowed)
				if _, err := stmt.Exec(pythagoreanExpectation, teamPitching.TeamID, teamPitching.Year); err != nil {
					fmt.Println(teamPitching.TeamID + ":" + teamPitching.Year)
					log.Print(err)
				}
			}
		}
	}
}

func calcPythagoreanExpectation(score int, runsAllowed int) float64 {
	fScore := float64(score)
	fRunsAllowed := float64(runsAllowed)
	return (fScore * fScore) / ((fScore * fScore) + (fRunsAllowed * fRunsAllowed))
}

// InsertTeamLeagueStats チームごとのシーズン成績をDBに登録する
func (Repository *TeamRepository) InsertTeamLeagueStats(teamLeagueStats []teamData.TeamLeagueStats) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO team_season_stats(team_id, year, manager, games, win, lose, draw, winning_rate, exchange_win, exchange_lose, exchange_draw, home_win, home_lose, home_draw, load_win, load_lose, load_draw) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	for _, stats := range teamLeagueStats {
		if _, err := stmt.Exec(stats.TeamID, stats.Year, stats.Manager, stats.Games, stats.Win, stats.Lose, stats.Draw, stats.WinningRate, stats.ExchangeWin, stats.ExchangeLose, stats.ExchangeDraw, stats.HomeWin, stats.HomeLose, stats.HomeDraw, stats.LoadWin, stats.LoadLose, stats.LoadDraw); err != nil {
			fmt.Println(stats.TeamID + ":" + stats.Year)
			log.Print(err)
		}
	}
}

// InsertMatchResults 各チームの対戦成績をDBに登録する
func (Repository *TeamRepository) InsertMatchResults(teamMatchResults []teamData.TeamMatchResults) {
	stmt, err := Repository.Conn.Prepare("INSERT INTO team_match_results(team_id, year, competitive_team_id, vs_type, win, lose, draw) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	for _, result := range teamMatchResults {
		if _, err := stmt.Exec(result.TeamID, result.Year, result.CompetitiveTeamID, result.VsType, result.Win, result.Lose, result.Draw); err != nil {
			fmt.Println(result.TeamID + ":" + result.Year)
			log.Print(err)
		}
	}
}

// GetTeamName マスタテーブルからチーム名を取得する
func (Repository *TeamRepository) GetTeamName(teamID string) (teamName string) {
	rows, err := Repository.Conn.Query("SELECT team_name FROM team_name WHERE team_name_id = $1", teamID)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&teamName)
	}

	return teamName
}
