package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "challenge.db")
	if err != nil {
		panic("failed to connect database")
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

func GetConnectionGORM() *gorm.DB {
	dsn := "host=localhost user=challenge password=challenge dbname=challenge port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type Database interface{}
