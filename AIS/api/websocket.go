package api

import (
	"aisstream/db/generated"
	"aisstream/db/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
)

func handleMessage(ctx context.Context, postgresDB *generated.Queries, p []byte) {
	var packet aisstream.AisStreamMessage
	if err := json.Unmarshal(p, &packet); err != nil {
		log.Println("Failed to unmarshal packet:", err)
		return
	}

	var shipName string
	if name, ok := packet.MetaData["ShipName"]; ok {
		shipName = name.(string)
	}

	switch packet.MessageType {
	case aisstream.POSITION_REPORT:
		if packet.Message.PositionReport == nil || !isValidShipName(shipName) {
			return
		}
		positionReport := *packet.Message.PositionReport

		positionReportArgs := models.BuildUpsertPositionEntryParams(shipName, positionReport)
		if err := postgresDB.UpsertPositionEntry(ctx, positionReportArgs); err != nil {
			log.Println("Failed to upsert position entry:", err)
			return
		}

		fmt.Printf("MMSI: %d Ship Name: %s Latitude: %f Longitude: %f\n",
			positionReport.UserID, shipName, positionReport.Latitude, positionReport.Longitude)
	}
}

func ConnectAndSubscribe(ctx context.Context, postgresDB *generated.Queries, url string, subscription aisstream.SubscriptionMessage) {
	for {
		log.Println("Connecting to WebSocket...")
		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Println("WebSocket connection failed:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Println("Connected to WebSocket")

		subMsgBytes, _ := json.Marshal(subscription)
		if err := ws.WriteMessage(websocket.TextMessage, subMsgBytes); err != nil {
			log.Println("Failed to send subscription message:", err)
			ws.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println("WebSocket closed or error:", err)
				ws.Close()
				break // reconnect
			}
			handleMessage(ctx, postgresDB, msg)
		}

		// wait a bit before reconnecting
		time.Sleep(5 * time.Second)
		log.Println("Reconnecting WebSocket...")
	}
}

func isValidShipName(shipName string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return !(len(shipName) == 0 || strings.ContainsFunc(shipName, isNotDigit))
}
