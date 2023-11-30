package handler

import (
	"log"
	"net/http"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/usecase"
	"github.com/gin-gonic/gin"
)

type GetLatestTurn struct {
	Usecase usecase.GetLatestTurnUsecase
}

type getLatestTurnResponse struct {
	Board      [][]*model.Disc `json:"board"`
	TurnCount  int             `json:"turn_count"`
	NextDisc   int             `json:"next_disc"`
	WinnerDisc int             `json:"winner_disc"`
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
		NextDisc:   int(*output.NextDisc),
		WinnerDisc: int(output.WinnerDisc),
	}

	log.Println(responseBody)

	c.JSON(http.StatusOK, responseBody)
}
