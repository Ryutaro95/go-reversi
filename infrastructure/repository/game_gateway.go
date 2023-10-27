package repository

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/domain/model"
)

type gameRepository struct{}

func NewGameRepository() *gameRepository {
	return &gameRepository{}
}

func (gr *gameRepository) Insert(db *sql.DB, game model.Game) (int, error) {
	in, err := db.Prepare("insert into games (started_at) values (?)")
	if err != nil {
		return 0, err
	}

	result, err := in.Exec(game.StartedAt)
	if err != nil {
		return 0, err
	}

	gameId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(gameId), nil
}
