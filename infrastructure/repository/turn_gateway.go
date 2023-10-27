package repository

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/domain/model"
)

type turnRepository struct{}

func NewTurnRepository() *turnRepository {
	return &turnRepository{}
}

func (tr *turnRepository) Insert(db *sql.DB, turn model.Turn) (int, error) {
	in, err := db.Prepare("insert into turns (game_id, turn_count, next_disc, end_at) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := in.Exec(turn.GameId, turn.TurnCount, turn.NextDisc, turn.EndAt)
	if err != nil {
		return 0, err
	}

	turnId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(turnId), nil
}
