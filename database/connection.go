package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetConnection() (*sql.DB, error) {
	once.Do(func() {
		db, err = sql.Open("sqlite3", "./database.db")
		if err != nil {
			fmt.Println("Error creating database: ", err)
		}

		statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS currency (symbol TEXT, price INTEGER, updated INTEGER)")
		statement.Exec()
	})

	return db, err
}
