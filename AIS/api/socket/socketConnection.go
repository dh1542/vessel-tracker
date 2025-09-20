package socket

import (
	"aisstream/db/generated"
	"context"
	"encoding/json"
	"log"
	"time"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
)

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
				break
			}
			HandleSocketMessage(ctx, postgresDB, msg)
		}
		time.Sleep(5 * time.Second)
		log.Println("Reconnecting WebSocket...")
	}
}
