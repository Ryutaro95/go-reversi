package database

import (
	"database/sql"
	"fmt"
	"log"
)

func NewDB() *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"reversi",
		"password",
		"localhost",
		"3306",
		"reversi",
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL init: %s", err)
	}

	return conn
}
