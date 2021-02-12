package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/team"
)

func getDB() (db *sql.DB) {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5555 password=postgres user=test dbname=test sslmode=disable")

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func main() {

	db := getDB()

	defer db.Close()

	// プレイヤーの成績をDBに登録する
	createdGades, _ := strconv.ParseBool(getSystemSetting("created_player_grades", db))
	if !createdGades {
		// リーグ文字列の配列
		leagues := []string{"b", "c", "d", "db", "e", "f", "g", "h", "l", "m", "s", "t"}

		for _, league := range leagues {
			setPlayerGrades(league, db)
		}

		setSystemSetting("created_player_grades", "true", db)
	}

	// チーム成績をDBに登録する
	createdTeamStats, _ := strconv.ParseBool(getSystemSetting("created_team_stats", db))
	if !createdTeamStats {
		setTeamStats(db)
		setSystemSetting("created_team_stats", "true", db)
	}

	// 算出が必要なDB項目を登録する
	createdAddValue, _ := strconv.ParseBool(getSystemSetting("created_add_value", db))
	if !createdAddValue {
		setTeamStatsAddValue(db)
		setSystemSetting("created_add_value", "true", db)
	}

	// webサーバーを起動
	router := setupRouter()
	router.Run(":8081")
}

func setTeamStatsAddValue(db *sql.DB) {
	years := makeRange(2005, 2020)
	teamBattings := team.GetTeamBatting(years, db)
	teamPitching := team.GetTeamPitching(years, db)
	team.InsertPythagoreanExpectation(years, teamBattings, teamPitching, db)
}

func getSystemSetting(setting string, db *sql.DB) (value string) {
	rows, err := db.Query("SELECT * FROM system_setting where setting = $1", setting)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var setting string
		rows.Scan(&setting, &value)
	}

	return value
}

func setSystemSetting(setting string, value string, db *sql.DB) {
	rows, err := db.Query("UPDATE system_setting SET value = $1 WHERE setting = $2", value, setting)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS対策
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	// チーム打撃成績を取得
	router.GET("/team/pitching", getTeamPitching)

	// チーム打撃成績を取得
	router.GET("/team/batting", getTeamBatting)

	// チーム成績を取得
	router.GET("/team/stats", getTeamStats)

	// 画面表示
	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	return router

}

func getTeamPitching(c *gin.Context) {
	db := getDB()
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamPitchingMap := team.GetTeamPitching(years, db)
	c.JSON(http.StatusOK, gin.H{
		"teamPitching": teamPitchingMap,
	})
}

func getTeamBatting(c *gin.Context) {
	db := getDB()
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamBattingMap := team.GetTeamBatting(years, db)
	c.JSON(http.StatusOK, gin.H{
		"teamBatting": teamBattingMap,
	})
}

func getTeamStats(c *gin.Context) {
	db := getDB()
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamStats := team.GetTeamStats(years, db)
	c.JSON(http.StatusOK, gin.H{
		"teanStats": teamStats,
	})
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func setTeamStats(db *sql.DB) {

	current, _ := os.Getwd()

	csvPath := current + "/" + "csv"

	// チーム打撃成績をDBに登録する
	team.InsertTeamBattings(csvPath, "central", db)
	team.InsertTeamBattings(csvPath, "pacific", db)

	// チーム投手成績をDBに登録する
	team.InsertTeamPitchings(csvPath, "central", db)
	team.InsertTeamPitchings(csvPath, "pacific", db)

	// チームシーズン成績をDBに登録する
	team.InsertSeasonLeagueStats(csvPath, db)
	team.InsertSeasonMatchResults(csvPath, db)

}

func setPlayerGrades(initial string, db *sql.DB) {

	current, _ := os.Getwd()

	players := grades.GetPlayers(current + "/csv/teams/" + initial + "_players.csv")

	playersPath := current + "/csv/players/"
	careers := grades.ReadCareers(playersPath, initial, players)

	grades.ExtractionCareers(&careers, db)

	grades.InsertCareers(careers, db)

	picherMap, batterMap := grades.ReadGradesMap(playersPath, initial, players)

	grades.ExtractionPicherGrades(&picherMap, team.GetTeamID(initial), db)

	grades.InsertPicherGrades(picherMap, db)

	grades.ExtractionBatterGrades(&batterMap, team.GetTeamID(initial), db)

	grades.InsertBatterGrades(batterMap, db)

}
