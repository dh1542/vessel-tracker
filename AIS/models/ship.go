package models

type Ship struct {
	MMSI     int64  `json:"mmsi" db:"mmsi"`
	ShipName string `json:"ShipName" db:"ship_name"`
}
