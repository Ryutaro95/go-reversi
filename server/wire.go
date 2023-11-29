//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/handler"
	"github.com/Ryutaro95/go-reversi/infrastructure/persistence"
	"github.com/Ryutaro95/go-reversi/usecase"
	"github.com/google/wire"
)

func InitStartGame(db *sql.DB) *handler.StartGame {
	wire.Build(
		handler.NewStartGame,
		usecase.NewStartGame,
		persistence.NewGamePersistence,
		persistence.NewTurnPersistence,
		persistence.NewSquarePersistence,
	)
	return &handler.StartGame{}
}

func InitLatestTurn(db *sql.DB) *handler.GetLatestTurn {
	wire.Build(
		persistence.NewGamePersistence,
		persistence.NewTurnPersistence,
		persistence.NewSquarePersistence,
		persistence.NewMovePersistence,
		usecase.NewGetLatestTurn,
		handler.NewGetLatestTurn,
	)
	return &handler.GetLatestTurn{}
}
