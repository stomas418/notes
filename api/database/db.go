package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}
	dns := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}
