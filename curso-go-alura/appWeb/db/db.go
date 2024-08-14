package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	conect := "user=postgres dbname=alura_loja password=0418 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err.Error())
	}
	return db
}
