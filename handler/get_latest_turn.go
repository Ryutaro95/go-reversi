package handler

import (
	"net/http"

	"github.com/Ryutaro95/go-reversi/usecase"
	"github.com/gin-gonic/gin"
)

type GetLatestTurn struct {
	Usecase usecase.GetLatestTurnUsecase
}

type getLatestTurnResponse struct {
	Board      [][]int `json:"board"`
	TurnCount  int     `json:"turn_count"`
	NextDisc   int     `json:"next_disc"`
	WinnerDisc int     `json:"winner_disc"`
}

func NewGetLatestTurn(g usecase.GetLatestTurnUsecase) *GetLatestTurn {
	return &GetLatestTurn{Usecase: g}
}

func (gt *GetLatestTurn) ServeHTTP(c *gin.Context) {
	turnCount := c.Param("turnCount")
	output := gt.Usecase.GetLatestTurn(turnCount)

	responseBody := &getLatestTurnResponse{
		TurnCount:  output.TurnCount,
		Board:      output.Board,
		NextDisc:   output.NextDisc,
		WinnerDisc: output.WinnerDisc,
	}

	c.JSON(http.StatusOK, responseBody)
}
