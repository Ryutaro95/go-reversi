package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/infrastructure/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/games", StartNewGame)

	router.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。
}

func StartNewGame(c *gin.Context) {
	// 現在時刻を取得
	now := time.Now()
	// ゲームデータをstructに保存
	game := model.Game{StartedAt: now}
	// mysqlのコネクションを取得
	db, err := sql.Open("mysql", "reversi:password@tcp(localhost:3306)/reversi")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer db.Close()

	// ゲーム開始時にgamesテーブルに開始日時の記録としてstartedAtをinsertする
	gameRepo := repository.NewGameRepository()
	gameId, err := gameRepo.Insert(db, game)
	if err != nil {
		panic(err)
	}
	game.Id = gameId

	// ゲーム開始時に1ターン目を保存
	turn := model.NewFirstTurn(game.Id, now)
	turnRepository := repository.NewTurnRepository()
	turnId, err := turnRepository.Insert(db, turn)
	if err != nil {
		panic(err)
	}

	squareRepository := repository.NewSquareRepository()
	squareRepository.InsertAll(db, turn, turnId)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
