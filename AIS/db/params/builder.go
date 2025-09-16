package params

import (
	"database/sql"

	"aisstream/db/generated"
	"aisstream/util"
	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
)

func BuildCreateShipParams(mmsi int64, shipName string) generated.CreateShipParams {
	return generated.CreateShipParams{
		Mmsi: mmsi,
		ShipName: sql.NullString{
			String: shipName,
			Valid:  shipName != "",
		},
	}
}

func BuildUpsertPositionEntryParams(report aisstream.PositionReport) generated.UpsertPositionEntryParams {
	return generated.UpsertPositionEntryParams{
		Mmsi:                      int64(report.UserID),
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
