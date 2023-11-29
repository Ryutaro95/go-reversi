package repository

import "github.com/Ryutaro95/go-reversi/domain/model"

type MoveRepo interface {
	FindByTurnId(turnId int64) (*model.Move, error)
}
