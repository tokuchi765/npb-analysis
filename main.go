package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/controller"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/infrastructure"
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

	sqlHandler := new(infrastructure.SQLHandler)
	sqlHandler.Conn = db
	teamInteractor := team.TeamInteractor{
		TeamRepository: infrastructure.TeamRepository{SQLHandler: *sqlHandler},
	}

	gradesInteractor := grades.GradesInteractor{
		GradesRepository: infrastructure.GradesRepository{SQLHandler: *sqlHandler},
		TeamRepository:   infrastructure.TeamRepository{SQLHandler: *sqlHandler},
	}

	syastemRepository := infrastructure.SyastemRepository{SQLHandler: *sqlHandler}

	// プレイヤーの成績をDBに登録する
	createdGades, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_player_grades"))
	if !createdGades {
		// リーグ文字列の配列
		leagues := []string{"b", "c", "d", "db", "e", "f", "g", "h", "l", "m", "s", "t"}

		for _, league := range leagues {
			setPlayerGrades(league, gradesInteractor, db)
		}

		syastemRepository.SetSystemSetting("created_player_grades", "true")
	}

	// チーム成績をDBに登録する
	createdTeamStats, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_team_stats"))
	if !createdTeamStats {
		setTeamStats(db, teamInteractor)
		syastemRepository.SetSystemSetting("created_team_stats", "true")
	}

	// 算出が必要なDB項目を登録する
	createdAddValue, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_add_value"))
	if !createdAddValue {
		setTeamStatsAddValue(teamInteractor)
		syastemRepository.SetSystemSetting("created_add_value", "true")
	}

	// webサーバーを起動
	router := setupRouter()
	router.Run(":8081")
}

func setTeamStatsAddValue(teamInteractor team.TeamInteractor) {
	years := makeRange(2005, 2020)
	teamBattings := teamInteractor.GetTeamBatting(years)
	teamPitching := teamInteractor.GetTeamPitching(years)
	teamInteractor.InsertPythagoreanExpectation(years, teamBattings, teamPitching)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS対策
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	sqlHandler := *infrastructure.NewSQLHandler()
	playerController := controller.NewPlayerController(sqlHandler)
	teamController := controller.NewTeamController(sqlHandler)

	// チーム打撃成績を取得
	router.GET("/team/pitching", func(c *gin.Context) { teamController.GetTeamPitching(c) })

	// チーム打撃成績を取得
	router.GET("/team/batting", func(c *gin.Context) { teamController.GetTeamBatting(c) })

	// チーム成績を取得
	router.GET("/team/stats", func(c *gin.Context) { teamController.GetTeamStats(c) })

	// チームごとの選手情報一覧を取得
	router.GET("/team/careers/:teamId/:year", func(c *gin.Context) { teamController.GetCareers(c) })

	// 選手情報取得
	router.GET("/player/:playerId", func(c *gin.Context) { playerController.GetPlayer(c) })

	// 画面表示
	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	return router

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func setTeamStats(db *sql.DB, teamInteractor team.TeamInteractor) {

	current, _ := os.Getwd()

	csvPath := current + "/" + "csv"

	// チーム打撃成績をDBに登録する
	teamInteractor.InsertTeamBattings(csvPath, "central", db)
	teamInteractor.InsertTeamBattings(csvPath, "pacific", db)

	// チーム投手成績をDBに登録する
	teamInteractor.InsertTeamPitchings(csvPath, "central", db)
	teamInteractor.InsertTeamPitchings(csvPath, "pacific", db)

	// チームシーズン成績をDBに登録する
	teamInteractor.InsertSeasonLeagueStats(csvPath)
	teamInteractor.InsertSeasonMatchResults(csvPath)

}

func setPlayerGrades(initial string, gradesInteractor grades.GradesInteractor, db *sql.DB) {

	current, _ := os.Getwd()

	players := grades.GetPlayers(current + "/csv/teams/" + initial + "_players.csv")

	gradesInteractor.InsertTeamPlayers(initial, players)

	playersPath := current + "/csv/players/"
	careers := grades.ReadCareers(playersPath, initial, players)

	gradesInteractor.ExtractionCareers(&careers)

	gradesInteractor.InsertCareers(careers)

	picherMap, batterMap := grades.ReadGradesMap(playersPath, initial, players)

	gradesInteractor.ExtractionPicherGrades(&picherMap, team.GetTeamID(initial))

	gradesInteractor.InsertPicherGrades(picherMap)

	gradesInteractor.ExtractionBatterGrades(&batterMap, team.GetTeamID(initial))

	gradesInteractor.InsertBatterGrades(batterMap, current)

}
