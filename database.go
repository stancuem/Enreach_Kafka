package main

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
}

func ConnectDatabase(url string) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &Database{conn: conn}, nil
}

func (db *Database) Save(user User) error {
	_, err := db.conn.Exec(context.Background(), "INSERT INTO users (first_name, last_name, job, age) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Job, user.Age)
	return err
}
