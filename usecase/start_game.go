package usecase

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type StartGameUsecase interface {
	StartGame() error
}

type StartGame struct {
	DB         *sql.DB
	GameRepo   repository.GameRepo
	TurnRepo   repository.TurnRepo
	SquareRepo repository.SquareRepo
}

func NewStartGame(gameRepo repository.GameRepo, turnRepo repository.TurnRepo, squareRepo repository.SquareRepo) StartGameUsecase {
	return &StartGame{
		GameRepo:   gameRepo,
		TurnRepo:   turnRepo,
		SquareRepo: squareRepo,
	}
}

func (g *StartGame) StartGame() error {
	now := time.Now()
	game := &model.Game{StartedAt: now}
	game, err := g.GameRepo.Create(game)
	if err != nil {
		return fmt.Errorf("StartGame() fail: %w", err)
	}

	turn := model.NewFirstTurn(game.ID, now)
	g.TurnRepo.Create(turn)

	if err := g.SquareRepo.InsertAll(turn); err != nil {
		return fmt.Errorf("StartGame() fail: %w", err)
	}

	return nil
}
