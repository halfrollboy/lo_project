package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db, nil
}
