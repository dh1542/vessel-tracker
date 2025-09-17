package models

import (
	"database/sql"

	"aisstream/db/generated"
	"aisstream/util"
	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
)

func BuildUpsertPositionEntryParams(shipName string, report aisstream.PositionReport) generated.UpsertPositionEntryParams {
	return generated.UpsertPositionEntryParams{
		Mmsi:                      int64(report.UserID),
		ShipName:                  sql.NullString{String: shipName, Valid: true},
		Latitude:                  sql.NullFloat64{Float64: report.Latitude, Valid: true},
		Longitude:                 sql.NullFloat64{Float64: report.Longitude, Valid: true},
		Cog:                       sql.NullInt32{Int32: int32(report.Cog), Valid: true},
		Sog:                       sql.NullInt32{Int32: int32(report.Sog), Valid: true},
		TrueHeading:               sql.NullInt32{Int32: int32(report.TrueHeading), Valid: true},
		NavigationalStatus:        sql.NullInt32{Int32: int32(report.NavigationalStatus), Valid: true},
		PositionAccuracy:          sql.NullBool{Bool: report.PositionAccuracy, Valid: true},
		CommunicationState:        sql.NullInt64{Int64: int64(report.CommunicationState), Valid: true},
		RateOfTurn:                sql.NullInt32{Int32: int32(report.RateOfTurn), Valid: true},
		SpecialManoeuvreIndicator: sql.NullInt32{Int32: int32(report.SpecialManoeuvreIndicator), Valid: true},
		RepeatIndicator:           sql.NullInt32{Int32: int32(report.RepeatIndicator), Valid: true},
		MessageID:                 sql.NullInt32{Int32: int32(report.MessageID), Valid: true},
		Valid:                     sql.NullBool{Bool: true, Valid: true},
		TimeUtc:                   sql.NullTime{Time: util.TimeFromInt32(report.Timestamp), Valid: true},
	}
}

func BuildGetPositionDataParams(minLatitude, minLongitude, maxLongitude, maxLatitude float64) generated.GetPositionDataParams {
	return generated.GetPositionDataParams{
		Latitude:    sql.NullFloat64{Float64: minLatitude, Valid: true},
		Latitude_2:  sql.NullFloat64{Float64: maxLatitude, Valid: true},
		Longitude:   sql.NullFloat64{Float64: minLongitude, Valid: true},
		Longitude_2: sql.NullFloat64{Float64: maxLongitude, Valid: true},
	}
}
