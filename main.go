package main

import (
	"log"

	"github.com/Ryutaro95/go-reversi/infrastructure/database"
	"github.com/Ryutaro95/go-reversi/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql := database.NewDB()
	defer func() {
		if err := mysql.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	server := server.NewServer(mysql)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
