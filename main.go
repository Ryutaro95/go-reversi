package main

import (
	"log"

	"github.com/Ryutaro95/go-reversi/infrastructure/database"
	routers "github.com/Ryutaro95/go-reversi/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	router := gin.Default()

	db, err := database.NewDB()

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// 同時使用可能なコネクション数
	// これにより同時実行できるコネクションの数が増える
	// 値が大きすぎるとDBに負荷がかかってしまうので値は適切に決める
	// db.SetMaxOpenConns(5)
	//
	// コネクションプールに保持する最大アイドルコネクション数
	// アイドルコネクション: クエリを実行していないコネクション
	// SetMaxIdleConns()
	if err != nil {
		return err
	}
	// routers.SetRouting(db, router)
	routers.SetRouting(db, router)
	router.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。

	return nil
}
