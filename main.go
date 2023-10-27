package main

import (
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
	if err != nil {
		return err
	}
	// routers.SetRouting(db, router)
	routers.SetRouting(db, router)
	router.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。

	return nil
}
