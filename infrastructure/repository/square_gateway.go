package repository

import (
	"database/sql"
	"strings"

	"github.com/Ryutaro95/go-reversi/domain/model"
)

type squareRepository struct{}

func NewSquareRepository() *squareRepository {
	return &squareRepository{}
}


// このようなSQLを作成してInsertする
// INSERT INTO KEY_VALUE
// 	(KEY_NO, STRING_VALUE, NUMBER_VALUE)
// VALUES
// 	(1, 'VALUE1', 100),
// 	(2, 'VALUE2', 200),
// 	(3, 'VALUE3', 300),
// 	(4, 'VALUE4', 400),
// 	(5, 'VALUE5', 500),
// 	(6, 'VALUE6', 600),
// 	(7, 'VALUE7', 700),
// 	(8, 'VALUE8', 800),
// 	(9, 'VALUE9', 900),
// 	(10, 'VALUE10', 1000);
func (sr *squareRepository) InsertAll(db *sql.DB, turn model.Turn, turnId int) error {
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

	in, err := db.Prepare(squaresInsertSql)
	if err != nil {
		return err
	}

	squaresInsertValues := []any{}
	for y, line := range turn.Board {
		for x, disc := range line {
			squaresInsertValues = append(squaresInsertValues, turnId)
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
