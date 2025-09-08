package main

import (
	"encoding/json"
	"fmt"
	"log"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
)

func main() {
	url := "wss://stream.aisstream.io/v0/stream"

	// websocket
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer ws.Close()

	subscriptionMessage := aisstream.SubscriptionMessage{
		APIKey:        "",
		BoundingBoxes: [][][]float64{{{-90.0, -180.0}, {90.0, 180.0}}}, // bounding box for the entire world
	}

	subMsgBytes, _ := json.Marshal(subscriptionMessage)
	if err := ws.WriteMessage(websocket.TextMessage, subMsgBytes); err != nil {
		log.Fatalln(err)
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}
		var packet aisstream.AisStreamMessage

		err = json.Unmarshal(p, &packet)
		if err != nil {
			log.Fatalln(err)
		}

		var shipName string
		// field may or may not be populated
		if packetShipName, ok := packet.MetaData["ShipName"]; ok {
			shipName = packetShipName.(string)
		}

		switch packet.MessageType {
		case aisstream.POSITION_REPORT:
			var positionReport aisstream.PositionReport
			positionReport = *packet.Message.PositionReport
			fmt.Printf("MMSI: %d Ship Name: %s Latitude: %f Longitude: %f\n",
				positionReport.UserID, shipName, positionReport.Latitude, positionReport.Longitude)
		}

	}

}
