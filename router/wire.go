//go:build wireinject
// +build wireinject

package router

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
