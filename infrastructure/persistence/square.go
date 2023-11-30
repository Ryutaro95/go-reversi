package persistence

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/Ryutaro95/go-reversi/domain/repository"
)

type SquarePersistence struct {
	DB *sql.DB
}

func NewSquarePersistence(db *sql.DB) repository.SquareRepo {
	return &SquarePersistence{DB: db}
}

func (sp *SquarePersistence) FetchSquaresByTurnID(turnId int64) ([]*model.Square, error) {
	var squares []*model.Square
	rows, err := sp.DB.Query("select id, turn_id, x, y, disc from squares where turn_id = ?", turnId)
	if err != nil {
		return []*model.Square{}, fmt.Errorf("FetchSquaresByTurnID() -> sp.DB.Query() fail: %w", err)
	}

	for rows.Next() {
		square := &model.Square{}
		if err := rows.Scan(&square.ID, &square.TurnID, &square.X, &square.Y, &square.Disc); err != nil {
			return []*model.Square{}, fmt.Errorf("FetchSquaresByTurnID() fail: %w", err)
		}
		squares = append(squares, square)
	}
	if err = rows.Err(); err != nil {
		return []*model.Square{}, fmt.Errorf("FetchSquaresByTurnID() fail: %w", err)
	}

	return squares, nil
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
