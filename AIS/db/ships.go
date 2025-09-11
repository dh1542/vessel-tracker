package db

import (
	"aisstream/models"
	"database/sql"
	"errors"
	"fmt"
)

func getShip(MMSI int64, db *sql.DB) (*models.Ship, error) {
	query := `SELECT mmsi, ship_name FROM ships WHERE MMSI = $1`

	var ship models.Ship
	row := db.QueryRow(query, MMSI)

	err := row.Scan(&ship.MMSI, &ship.ShipName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying ship: %w", err)
	}
	return &ship, nil
}

func CreateShip(MMSI int64, shipName string, db *sql.DB) error {
	query := `INSERT INTO ships (mmsi, ship_name) VALUES ($1, $2)`
	_, err := db.Exec(query, MMSI, shipName)
	if err != nil {
		return err
	}
	return nil
}
