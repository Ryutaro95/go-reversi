package database

import (
	"database/sql"
	"fmt"
)

func NewDB() (*sql.DB, error) {
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
		return nil, err
	}

	return conn, nil
}
