package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	conn := "user=postgres dbname=store password=admin host=localhost port=3001 sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}
	return db
}
