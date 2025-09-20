package socket

import (
	"aisstream/db/generated"
	"aisstream/db/models"
	"aisstream/util"
	"context"
	"encoding/json"
	"fmt"
	"log"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
)

func HandleSocketMessage(ctx context.Context, postgresDB *generated.Queries, p []byte) {
	var packet aisstream.AisStreamMessage
	if err := json.Unmarshal(p, &packet); err != nil {
		log.Println("Failed to unmarshal packet:", err)
		return
	}

	var shipName string
	if name, ok := packet.MetaData["ShipName"]; ok {
		shipName = name.(string)
	}

	if !util.IsValidShipName(shipName) {
		return
	}

	switch packet.MessageType {
	case aisstream.POSITION_REPORT:
		handlePositionMessage(packet, shipName, ctx, postgresDB)
	case aisstream.SHIP_STATIC_DATA:
		handleStaticMessage(packet, shipName, ctx, postgresDB)
	}
}

func handlePositionMessage(packet aisstream.AisStreamMessage, shipName string, ctx context.Context, postgresDB *generated.Queries) {
	//log.Println("Received a Position Report")
	if packet.Message.PositionReport == nil {
		return
	}
	positionReport := *packet.Message.PositionReport

	positionReportArgs := models.BuildUpsertPositionEntryParams(shipName, positionReport)
	err := postgresDB.UpsertPositionEntry(ctx, positionReportArgs)
	if err != nil {
		log.Println("Failed to upsert position entry:", err)
		return
	}
}

func handleStaticMessage(packet aisstream.AisStreamMessage, shipName string, ctx context.Context, postgresDB *generated.Queries) {
	log.Println("Received a Static Data")
	staticReport := *packet.Message.ShipStaticData
	fmt.Println(staticReport.Destination)
	//fmt.Println(staticReport)
}
