package repository

import "github.com/Ryutaro95/go-reversi/domain/model"

type TurnRepo interface {
	Create(turn *model.Turn) (*model.Turn, error)
}
