package postgres

import (
	"database/sql"
	"log"
)

type Client struct {
	DB *sql.DB
}

func CreateConn() (*Client, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &Client{DB: db}, err
}
