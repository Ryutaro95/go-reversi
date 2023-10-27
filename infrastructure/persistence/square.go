package persistence

import (
	"database/sql"
	"strings"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/model/repository"
)

type SquarePersistence struct {
	DB *sql.DB
}

func NewSquarePersistence(db *sql.DB) repository.SquareRepo {
	return &SquarePersistence{DB: db}
}

func (sp *SquarePersistence) InsertAll(turn *model.Turn) error {
	squaresInsertSql := "insert into squares (turn_id, x, y, disc) values "
	squaresCount := 0
	// 盤面数を計算
	for _, line := range turn.Board {
		squaresCount += len(line)
	}
	querys := []string{}
	for i := 0; i < squaresCount; i++ {
		querys = append(querys, "(?, ?, ?, ?)")
	}
	query := strings.Join(querys, ", ")
	squaresInsertSql += query

	in, err := sp.DB.Prepare(squaresInsertSql)
	if err != nil {
		return err
	}

	squaresInsertValues := []any{}
	for y, line := range turn.Board {
		for x, disc := range line {
			squaresInsertValues = append(squaresInsertValues, turn.ID)
			squaresInsertValues = append(squaresInsertValues, x)
			squaresInsertValues = append(squaresInsertValues, y)
			squaresInsertValues = append(squaresInsertValues, disc)
		}
	}
	_, err = in.Exec(squaresInsertValues...)
	if err != nil {
		return err
	}

	return nil
}
