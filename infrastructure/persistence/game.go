package persistence

import (
	"database/sql"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type GamePersistence struct {
	DB *sql.DB
}

func NewGamePersistence(db *sql.DB) repository.GameRepo {
	return &GamePersistence{DB: db}
}

func (gp *GamePersistence) Create(game *model.Game) (*model.Game, error) {
	in, err := gp.DB.Prepare("insert into games (started_at) values (?)")
	if err != nil {
		return &model.Game{}, err
	}

	result, err := in.Exec(game.StartedAt)
	if err != nil {
		return &model.Game{}, err
	}

	gameID, err := result.LastInsertId()
	game.ID = gameID
	if err != nil {
		return &model.Game{}, err
	}
	
	return game, nil
}
