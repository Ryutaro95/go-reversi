package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Ryutaro95/go-reversi/domain/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/games", StartNewGame)

	router.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。
}


func StartNewGame(c *gin.Context) {
	// 現在時刻を取得
	now := time.Now()
	// ゲームデータをstructに保存
	game := model.Game{StartedAt: now}
	// mysqlのコネクションを取得
	db, err := sql.Open("mysql", "reversi:password@tcp(localhost:3306)/reversi")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer db.Close()

	// ゲーム開始時にgamesテーブルに開始日時の記録としてstartedAtをinsertする
	in, err := db.Prepare("insert into games (started_at) values (?)")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	result, err := in.Exec(game.StartedAt)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	game.Id = int(insertId)

	// 1ターン目を保存
	turn := model.NewFirstTurn(game.Id, now)
	in, err = db.Prepare("insert into turns (game_id, turn_count, next_disc, end_at) values (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	result, err = in.Exec(turn.GameId, turn.TurnCount, turn.NextDisc, turn.EndAt)
	turnId, err := result.LastInsertId()

	// squaresに開始時点の盤面を保存
	// squaresの1レコードは1マスのため、1盤面保存するためには64レコード保存する必要がある
	squaresInsertSql := "insert into squares (turn_id, x, y, disc) values "
	squaresCount := 0
	for _, line := range turn.Board {
		squaresCount += len(line)
	}
	querys := []string{}
	for i := 0; i < squaresCount; i++ {
		querys = append(querys, "(?, ?, ?, ?)")
	}
	query := strings.Join(querys, ", ")
	squaresInsertSql += query
	// そのため上の処理で以下のようなSQLを作り上げている
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
	in, err = db.Prepare(squaresInsertSql)
	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}
	// move　= 手
	// プレイヤーが石を盤面に打ったらその情報をmovesテーブル保存するが、今回はゲームが始まっただけで
	// 誰もターンを消費して手を打っているわけではないので、movesテーブルのinsertは発生しない
	// ゲーム開始後はmovesテーブルへのinsertが必要になる
	//
	// 以下の構造体はそのmoveをinsertするときに欲しくなると思う
	// ただ、必要な情報はturnRecord.idだけなので、この構造体がどうしても欲しい物なのかは分からない
	// turnRecord := TurnRecord{
	// 	id:        int(turnId),
	// 	gameId:    turn.gameId,
	// 	turnCount: turn.turnCount,
	// 	nextDisc:  turn.nextDisc,
	// 	endAt:     turn.endAt,
	// }

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
