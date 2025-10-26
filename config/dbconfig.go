package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDBConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "user=m0_user dbname=m0_database sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to db with err: %v", err)
	}
	return db
}
