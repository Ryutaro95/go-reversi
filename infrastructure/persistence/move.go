package persistence

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type MovePersistence struct {
	DB *sql.DB
}

func NewMovePersistence(db *sql.DB) repository.MoveRepo {
	return &MovePersistence{DB: db}
}

func (mp *MovePersistence) FindByTurnId(turnId int64) (*model.Move, error) {
	move := &model.Move{}
	if err := mp.DB.QueryRow(
		"select id, turn_id, disc, x, y from moves where turn_id = ?",
		turnId,
	).Scan(&move.ID, &move.TurnID, &move.Disc, &move.X, &move.Y); err != nil {
		return &model.Move{}, nil
	}
	return move, nil
}
