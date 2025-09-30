package main

import (
	"aisstream/api"
	"aisstream/api/socket"
	"aisstream/db"

	"context"
	"log"
	"os"

	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	postgresDB := db.InitDB()
	if err := postgresDB.CreatePositionReportTableIfNotExist(ctx); err != nil {
		log.Fatal(err)
	}
	if err := postgresDB.CreateImagesTableIfNotExist(ctx); err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("AIS_STREAM_URL")
	subscription := aisstream.SubscriptionMessage{
		APIKey: os.Getenv("AIS_STREAM_API_KEY"),
		BoundingBoxes: [][][]float64{
			{{54.71668856895074, 9.148864746093752}, {55.85064987433714, 12.293701171875002}},
		},
	}

	go api.ServeHTTPServer(ctx, postgresDB)

	socket.ConnectAndSubscribe(ctx, postgresDB, url, subscription)

}
