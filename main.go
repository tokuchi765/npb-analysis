package main

import (
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tokuchi765/npb-analysis/controller"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/infrastructure"
	"github.com/tokuchi765/npb-analysis/team"
)

func main() {
	sqlHandler := infrastructure.NewSQLHandler()
	teamInteractor := team.TeamInteractor{
		TeamRepository: &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
	}

	gradesInteractor := grades.GradesInteractor{
		GradesRepository: &infrastructure.GradesRepository{SQLHandler: *sqlHandler},
		TeamRepository:   &infrastructure.TeamRepository{SQLHandler: *sqlHandler},
	}

	syastemRepository := infrastructure.SyastemRepository{SQLHandler: *sqlHandler}

	// プレイヤーの成績をDBに登録する
	createdGades, _ := strconv.ParseBool(syastemRepository.GetSystemSetting("created_player_grades"))
	if !createdGades {
		// リーグ文字列の配列
		leagues := []string{"b", "c", "d", "db", "e", "f", "g", "h", "l", "m", "s", "t"}

		for _, league := range leagues {
			setPlayerGrades(league, gradesInteractor)
		}

		syastemRepository.SetSystemSetting("created_player_grades", "true")
	}

	years := makeRange(2005, 2021)

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

func setPlayerGrades(initial string, gradesInteractor grades.GradesInteractor) {

	current, _ := os.Getwd()

	csvPath := current + "/csv"

	// 2020~2021の選手一覧を取得する
	years := []string{"2020", "2021"}
	for _, year := range years {
		players := gradesInteractor.GetPlayers(csvPath, initial, year)

		gradesInteractor.InsertTeamPlayers(initial, players, year)

		careers := gradesInteractor.ReadCareers(csvPath, initial, players)

		gradesInteractor.ExtractionCareers(&careers)

		gradesInteractor.InsertCareers(careers)

		picherMap, batterMap := gradesInteractor.ReadGradesMap(csvPath, initial, players)

		gradesInteractor.ExtractionPicherGrades(&picherMap, gradesInteractor.TeamUtil.GetTeamID(initial))

		gradesInteractor.InsertPicherGrades(picherMap)

		gradesInteractor.ExtractionBatterGrades(&batterMap, gradesInteractor.TeamUtil.GetTeamID(initial))

		gradesInteractor.InsertBatterGrades(batterMap, current)
	}
}
