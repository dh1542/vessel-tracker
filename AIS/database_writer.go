package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func main() {
	db := connectToDB()
	setUpDBTables(db)

	fmt.Println("Database connected successfully")

}

func connectToDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func setUpDBTables(db *sql.DB) {

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS ships (
		id SERIAL PRIMARY KEY,
		mmsi BIGINT UNIQUE,
		ship_name VARCHAR(255),
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created or already exists")

}
