package persistence

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type TurnPersistence struct {
	DB *sql.DB
}

func NewTurnPersistence(db *sql.DB) repository.TurnRepo {
	return &TurnPersistence{DB: db}
}

func (tp *TurnPersistence) Create(turn *model.Turn) (*model.Turn, error) {
	in, err := tp.DB.Prepare("insert into turns (game_id, turn_count, next_disc, end_at) values (?, ?, ?, ?)")
	if err != nil {
		return &model.Turn{}, err
	}

	result, err := in.Exec(turn.GameID, turn.TurnCount, turn.NextDisc, turn.EndAt)
	if err != nil {
		return &model.Turn{}, err
	}

	turnID, err := result.LastInsertId()
	if err != nil {
		return &model.Turn{}, err
	}
	turn.ID = turnID
	return turn, nil
}
