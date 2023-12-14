package app

import (
	"database/sql"
	"learn_golang_api/app/helpers"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:admin12345@localhost:5432/learn_go_api")
	helpers.PanicOnError(err)

	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
