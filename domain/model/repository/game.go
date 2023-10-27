package repository

import (
	"github.com/Ryutaro95/go-reversi/domain/model"
)

type GameRepo interface {
	Create(game *model.Game) (*model.Game, error)
}
