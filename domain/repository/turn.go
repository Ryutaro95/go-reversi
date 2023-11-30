package repository

import "github.com/Ryutaro95/go-reversi/domain/model"

type TurnRepo interface {
	Create(turn *model.Turn) (*model.Turn, error)
	FindForGameIdAndTurnCount(gameId int64, turnCount string) (*model.Turn, error)
}
