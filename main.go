package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/controller"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/infrastructure"
	"github.com/tokuchi765/npb-analysis/infrastructure/csv"
	"github.com/tokuchi765/npb-analysis/team"
)

func main() {

	migration()

	sqlHandler := infrastructure.NewSQLHandler()
	teamInteractor := team.TeamInteractor{
		TeamRepository: &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
		TeamReader:     &csv.TeamReader{},
	}

	gradesInteractor := grades.GradesInteractor{
		GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
		GradesReader:     &csv.GradesReader{},
		TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
	}

	syastemRepository := infrastructure.SyastemRepository{SQLHandler: *sqlHandler}

	current, _ := os.Getwd()
	csvPath := current + "/" + "csv"

	// 選手キャリア情報をDBに登録する
	createdPlayerCareers, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_player_careers"))
	if !createdPlayerCareers {
		gradesInteractor.InsertCareers(csvPath)
		syastemRepository.SetSystemSetting("created_player_careers", "true")
	}

	// 選手打撃情報をDBに登録する
	createdPlayerBattings, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_player_battings"))
	if !createdPlayerBattings {
		gradesInteractor.InsertBatterGrades(current)
		syastemRepository.SetSystemSetting("created_player_battings", "true")
	}

	// 選手投手情報をDBに登録する
	createdPlayerPitchings, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_player_pitchings"))
	if !createdPlayerPitchings {
		gradesInteractor.InsertPicherGrades(csvPath)
		syastemRepository.SetSystemSetting("created_player_pitchings", "true")
	}

	// 選手一覧をDBに登録する
	createdTeamPlayers, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_team_players"))
	if !createdTeamPlayers {
		// リーグ文字列の配列
		leagues := []string{"b", "bs", "c", "d", "db", "e", "f", "g", "h", "l", "m", "s", "t"}

		for _, league := range leagues {
			teamInteractor.InsertTeamPlayers(csvPath, league)
		}

		syastemRepository.SetSystemSetting("created_team_players", "true")
	}

	years := makeRange(2005, 2023)

	// チーム成績をDBに登録する
	createdTeamStats, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_team_stats"))
	if !createdTeamStats {
		setTeamStats(teamInteractor, years)
		syastemRepository.SetSystemSetting("created_team_stats", "true")
	}

	// 算出が必要なDB項目を登録する
	createdAddValue, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_add_value"))
	if !createdAddValue {
		setTeamStatsAddValue(teamInteractor, years)
		syastemRepository.SetSystemSetting("created_add_value", "true")
	}

	// webサーバーを起動
	router := setupRouter()
	router.Run(":8081")
}

func migration() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://npb-analysis:postgres@localhost:5555/npb-analysis?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func setTeamStatsAddValue(teamInteractor team.TeamInteractor, years []int) {
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

	// チーム投手成績を取得
	router.GET("/team/pitching", func(c *gin.Context) { teamController.GetTeamPitching(c) })
	router.GET("/team/pitching/:teamId/:year", func(c *gin.Context) { teamController.GetTeamPitchingByTeamIDAndYear(c) })
	router.GET("/team/pitching/max", func(c *gin.Context) { teamController.GetTeamPitchingMax(c) })
	router.GET("/team/pitching/min", func(c *gin.Context) { teamController.GetTeamPitchingMin(c) })

	// チーム打撃成績を取得
	router.GET("/team/batting", func(c *gin.Context) { teamController.GetTeamBatting(c) })
	router.GET("/team/batting/:teamId/:year", func(c *gin.Context) { teamController.GetTeamBattingByTeamIDAndYear(c) })
	router.GET("/team/batting/max", func(c *gin.Context) { teamController.GetTeamBattingMax(c) })
	router.GET("/team/batting/min", func(c *gin.Context) { teamController.GetTeamBattingMin(c) })

	// チーム成績を取得
	router.GET("/team/stats", func(c *gin.Context) { teamController.GetTeamStats(c) })

	// チームごとの選手情報一覧を取得
	router.GET("/team/careers/:teamId/:year", func(c *gin.Context) { teamController.GetCareers(c) })

	// 選手情報取得
	router.GET("/player/:playerId", func(c *gin.Context) { playerController.GetPlayer(c) })
	router.GET("/player/search", func(c *gin.Context) { playerController.SearchPlayer(c) })
	router.GET("/player/search/batter-grades", func(c *gin.Context) { playerController.SearchBatterGrades(c) })
	router.GET("/player/search/pitcher-grades", func(c *gin.Context) { playerController.SearchPitcherGrades(c) })

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

func setTeamStats(teamInteractor team.TeamInteractor, years []int) {

	current, _ := os.Getwd()

	csvPath := current + "/" + "csv"

	// チーム打撃成績をDBに登録する
	teamInteractor.InsertTeamBattings(csvPath, "central", years)
	teamInteractor.InsertTeamBattings(csvPath, "pacific", years)

	// チーム投手成績をDBに登録する
	teamInteractor.InsertTeamPitchings(csvPath, "central", years)
	teamInteractor.InsertTeamPitchings(csvPath, "pacific", years)

	// チームシーズン成績をDBに登録する
	teamInteractor.InsertSeasonLeagueStats(csvPath, years)
	teamInteractor.InsertSeasonMatchResults(csvPath, years)

}
