package models

import (
	"aisstream/db/generated"
	"aisstream/util"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
)

func BuildUpsertPositionEntryParams(shipName string, report aisstream.PositionReport) generated.UpsertPositionEntryParams {
	return generated.UpsertPositionEntryParams{
		Mmsi:                      int64(report.UserID),
		ShipName:                  shipName,
		Latitude:                  report.Latitude,
		Longitude:                 report.Longitude,
		Cog:                       int32(report.Cog),
		Sog:                       int32(report.Sog),
		TrueHeading:               report.TrueHeading,
		NavigationalStatus:        report.NavigationalStatus,
		PositionAccuracy:          report.PositionAccuracy,
		CommunicationState:        int64(report.CommunicationState),
		RateOfTurn:                report.RateOfTurn,
		SpecialManoeuvreIndicator: report.SpecialManoeuvreIndicator,
		RepeatIndicator:           report.RepeatIndicator,
		MessageID:                 report.MessageID,
		Valid:                     true,
		TimeUtc:                   util.TimeFromInt32(report.Timestamp),
	}
}

func BuildGetPositionDataParams(minLatitude, minLongitude, maxLongitude, maxLatitude float64) generated.GetPositionDataParams {
	return generated.GetPositionDataParams{
		Latitude:    minLatitude,
		Latitude_2:  maxLatitude,
		Longitude:   minLongitude,
		Longitude_2: maxLongitude,
	}
}
