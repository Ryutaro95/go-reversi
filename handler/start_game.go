package handler

import (
	"net/http"

	"github.com/Ryutaro95/go-reversi/usecase"
	"github.com/gin-gonic/gin"
)

type StartGame struct {
	Usecase usecase.StartGameUsecase
}

func NewStartGame(s usecase.StartGameUsecase) *StartGame {
	return &StartGame{Usecase: s}
}

func (sg *StartGame) ServeHTTP(c *gin.Context) {
	if err := sg.Usecase.StartGame(); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
