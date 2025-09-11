package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	db := connect()
	setUpDBTables(db)
	emptyDBTables(db)
	return db
}

func connect() *sql.DB {
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Failed to parse DB_PORT env: %s", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to DB")

	return db
}

func setUpDBTables(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS ships (
		mmsi BIGINT PRIMARY KEY,
		ship_name VARCHAR(255)
	);

	CREATE TABLE IF NOT EXISTS position_reports (
		id SERIAL PRIMARY KEY,
		mmsi BIGINT REFERENCES ships(mmsi) ON DELETE CASCADE,
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		cog INTEGER,
		sog INTEGER,
		true_heading INTEGER,
		navigational_status INTEGER,
		position_accuracy BOOLEAN,
		communication_state BIGINT,
		rate_of_turn INTEGER,
		special_manoeuvre_indicator INTEGER,
		repeat_indicator INTEGER,
		message_id INTEGER,
		valid BOOLEAN,
		time_utc TIMESTAMP
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed on create table: %s ", err)
	}
	log.Println("Successfully created table or table already existed")
}

func emptyDBTables(db *sql.DB) {
	query := `TRUNCATE TABLE position_reports RESTART IDENTITY CASCADE;`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete table data for position_reports")
}
