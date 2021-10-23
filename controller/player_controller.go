package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tokuchi765/npb-analysis/grades"
	"github.com/tokuchi765/npb-analysis/infrastructure"
)

// PlayerController チームデータを管理するAPI
type PlayerController struct {
	GradesInteractor grades.GradesInteractor
}

// NewPlayerController PlayerControllerを生成
func NewPlayerController(sqlHandler infrastructure.SQLHandler) *PlayerController {
	return &PlayerController{
		GradesInteractor: grades.GradesInteractor{
			GradesRepository: infrastructure.GradesRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetPlayer 選手情報取得を取得します。
func (controller *PlayerController) GetPlayer(c Context) {
	playerID := c.Param("playerId")
	c.JSON(http.StatusOK, gin.H{
		"career":   controller.GradesInteractor.GetCareer(playerID),
		"batting":  controller.GradesInteractor.GetBatting(playerID),
		"pitching": controller.GradesInteractor.GetPitching(playerID),
	})
}
