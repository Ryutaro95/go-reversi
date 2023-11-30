package repository

import "github.com/Ryutaro95/go-reversi/domain/model"

type SquareRepo interface {
	InsertAll(turn *model.Turn) error
	FetchSquaresByTurnID(turnID int64) ([]*model.Square, error)
}
