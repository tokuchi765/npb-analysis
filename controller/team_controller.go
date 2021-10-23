package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tokuchi765/npb-analysis/entity/player"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/infrastructure"
	"github.com/tokuchi765/npb-analysis/team"
)

// TeamController チームデータを管理するAPI
type TeamController struct {
	TeamInteractor   team.TeamInteractor
	GradesInteractor grades.GradesInteractor
}

// NewTeamController TeamControllerを生成
func NewTeamController(sqlHandler infrastructure.SQLHandler) *TeamController {
	return &TeamController{
		TeamInteractor: team.TeamInteractor{
			TeamRepository: infrastructure.TeamRepository{SQLHandler: sqlHandler},
		},
		GradesInteractor: grades.GradesInteractor{
			GradesRepository: infrastructure.GradesRepository{
				SQLHandler: sqlHandler,
			},
		}}
}

// GetTeamPitching 引数で受け取った年に紐づくチーム投手成績を取得します。
func (controller *TeamController) GetTeamPitching(c Context) {
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamPitchingMap := controller.TeamInteractor.GetTeamPitching(years)
	c.JSON(http.StatusOK, gin.H{
		"teamPitching": teamPitchingMap,
	})
}

// GetTeamBatting 引数で受け取った年に紐づくチーム打撃成績を取得します。
func (controller *TeamController) GetTeamBatting(c Context) {
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamBattingMap := controller.TeamInteractor.GetTeamBatting(years)
	c.JSON(http.StatusOK, gin.H{
		"teamBatting": teamBattingMap,
	})
}

// GetTeamStats 引数で受け取った年に紐づくチーム成績を取得します。
func (controller *TeamController) GetTeamStats(c Context) {
	fromYear, _ := strconv.Atoi(c.Query("from_year"))
	toYear, _ := strconv.Atoi(c.Query("to_year"))
	years := makeRange(fromYear, toYear)
	teamStats := controller.TeamInteractor.GetTeamStats(years)
	c.JSON(http.StatusOK, gin.H{
		"teanStats": teamStats,
	})
}

// GetCareers チームごとの選手情報一覧を取得
func (controller *TeamController) GetCareers(c Context) {
	teamID := c.Param("teamId")
	year := c.Param("year")

	players := controller.GradesInteractor.GetPlayersByTeamIDAndYear(teamID, year)
	var careers []player.CAREER
	for _, player := range players {
		career := controller.GradesInteractor.GetCareer(player.PlayerID)
		careers = append(careers, career)
	}
	c.JSON(http.StatusOK, gin.H{
		"careers": careers,
	})
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}