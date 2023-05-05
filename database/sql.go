package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetupSql() (*sql.DB, error) {
	err := godotenv.Load(".env")
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))

	db, err := sql.Open("mysql", config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
