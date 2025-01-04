package database

import (
	"log"
	"database/sql"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

)

var db *sql.DB

func GetDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := mysql.Config{
        User:   os.Getenv("USERNAME"),
        Passwd: os.Getenv("PW"),
        Net:    os.Getenv("NET"),
        Addr:   os.Getenv("ADDR"),
		DBName: os.Getenv("DBNAME"),
    }

    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	return db, nil
}
