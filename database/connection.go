package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "challenge.db")
	if err != nil {
		panic(err)
	}
	return db
}

func MakeMigrations(db *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS endpoint1 (
	);
	
	CREATE TABLE IF NOT EXISTS endpoint2 (
	);`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
