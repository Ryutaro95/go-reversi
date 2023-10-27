package usecase

import (
	"database/sql"
	"time"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type StartGameUsecase interface {
	StartGame() error
}

type StartGame struct {
	DB       *sql.DB
	GameRepo repository.GameRepo
	TurnRepo repository.TurnRepo
	SquareRepo repository.SquareRepo
}

func NewStartGame(db *sql.DB, gameRepo repository.GameRepo, turnRepo repository.TurnRepo, squareRepo repository.SquareRepo) StartGameUsecase {
	return &StartGame{
		DB:       db,
		GameRepo: gameRepo,
		TurnRepo: turnRepo,
		SquareRepo: squareRepo,
	}
}

func (g *StartGame) StartGame() error {
	now := time.Now()
	// ゲーム開始時に保存
	game := &model.Game{StartedAt: now}
	game, err := g.GameRepo.Create(game)
	if err != nil {
		return err
	}

	// 初期ターンを保存
	turn := model.NewFirstTurn(game.ID, now)
	g.TurnRepo.Create(turn)

	// 初期盤面のマスを保存
	g.SquareRepo.InsertAll(turn)

	return nil
}
