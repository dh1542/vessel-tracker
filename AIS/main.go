package main

import (
	"aisstream/db"
	"encoding/json"
	"fmt"
	"log"
	"os"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresDataBase := db.InitDB()
	fmt.Println(postgresDataBase)

	url := os.Getenv("AIS_STREAM_URL")

	// websocket
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer ws.Close()

	subscriptionMessage := aisstream.SubscriptionMessage{
		APIKey:        os.Getenv("AIS_STREAM_API_KEY"),
		BoundingBoxes: [][][]float64{{{54.901184, 10.883331}, {55.463490, 11.057739}}}, // worl
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
		// case aisstream.Standard_Class_B_Position_Report:
		// 	var classBPositionReport aisstream.StandardClassBPositionReport
		// 	classBPositionReport = *packet.Message.StandardClassBPositionReport
		// 	fmt.Printf("MMSI: %d Ship Name: %s Latitude: %f Longitude: %f\n",
		// 		classBPositionReport.UserID, shipName, classBPositionReport.Latitude, classBPositionReport.Longitude)
		// }

	}

}
