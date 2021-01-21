package base

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func ConnectMysql() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

