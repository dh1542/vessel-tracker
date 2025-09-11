package models

import "time"

type positionReport struct {
	ID                        int64     `json:"id" db:"id"`
	MMSI                      int64     `json:"mmsi" db:"mmsi"`
	Latitude                  float64   `json:"latitude" db:"latitude"`
	Longitude                 float64   `json:"longitude" db:"longitude"`
	COG                       int       `json:"cog" db:"cog"`
	SOG                       int       `json:"sog" db:"sog"`
	TrueHeading               int       `json:"true_heading" db:"true_heading"`
	NavigationalStatus        int       `json:"navigational_status" db:"navigational_status"`
	PositionAccuracy          bool      `json:"position_accuracy" db:"position_accuracy"`
	CommunicationState        int64     `json:"communication_state" db:"communication_state"`
	RateOfTurn                int       `json:"rate_of_turn" db:"rate_of_turn"`
	SpecialManoeuvreIndicator int       `json:"special_manoeuvre_indicator" db:"special_manoeuvre_indicator"`
	RepeatIndicator           int       `json:"repeat_indicator" db:"repeat_indicator"`
	MessageID                 int       `json:"message_id" db:"message_id"`
	Valid                     bool      `json:"valid" db:"valid"`
	TimeUTC                   time.Time `json:"time_utc" db:"time_utc"`
}
