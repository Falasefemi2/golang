// package db

// import (
// 	"database/sql"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDB() {
// 	DB, err := sql.Open("sqlite3", "api.db")

// 	if err != nil {
// 		panic("Could not connect to database")
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	createTables()
// }

// func createTables() {
// 	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
//         name TEXT NOT NULL,
// 		description TEXT NOT NULL,
// 		location TEXT NOT NULL,
// 		dateTime DATETIME NOT NULL,
// 		user_id INTEGER
// 	)
// 	`

// 	_, err := DB.Exec(createEventsTable)

// 	if err != nil {
// 		panic("Could not create events table")
// 	}
// }

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err := createTables(); err != nil {
		return fmt.Errorf("could not create tables: %w", err)
	}

	return nil
}

func createTables() error {
	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER
        )
    `

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		return fmt.Errorf("could not create events table: %w", err)
	}

	return nil
}
