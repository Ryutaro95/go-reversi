// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package router

import (
	"database/sql"
	"github.com/Ryutaro95/go-reversi/handler"
	"github.com/Ryutaro95/go-reversi/infrastructure/persistence"
	"github.com/Ryutaro95/go-reversi/usecase"
)

// Injectors from wire.go:

func InitStartGame(db *sql.DB) *handler.StartGame {
	gameRepo := persistence.NewGamePersistence(db)
	turnRepo := persistence.NewTurnPersistence(db)
	squareRepo := persistence.NewSquarePersistence(db)
	startGameUsecase := usecase.NewStartGame(db, gameRepo, turnRepo, squareRepo)
	startGame := handler.NewStartGame(startGameUsecase)
	return startGame
}

func InitLatestTurn(db *sql.DB) *handler.GetLatestTurn {
	gameRepo := persistence.NewGamePersistence(db)
	turnRepo := persistence.NewTurnPersistence(db)
	squareRepo := persistence.NewSquarePersistence(db)
	moveRepo := persistence.NewMovePersistence(db)
	getLatestTurnUsecase := usecase.NewGetLatestTurn(db, gameRepo, turnRepo, squareRepo, moveRepo)
	getLatestTurn := handler.NewGetLatestTurn(getLatestTurnUsecase)
	return getLatestTurn
}
