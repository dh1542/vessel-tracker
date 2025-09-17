package api

import (
	"aisstream/db/generated"
	"aisstream/db/params"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type homeHandler struct{}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func ServeHTTPServer(ctx context.Context, db *generated.Queries) {
	router := mux.NewRouter()

	// Define endpoints
	router.HandleFunc("/api/{minLatitude}/{maxLatitude}/{minLongitude}/{maxLongitude}", GetShipsForPositionHandler(ctx, db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetShipsForPositionHandler(ctx context.Context, db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)

		minLatitude, err := strconv.ParseFloat(parameter["minLatitude"], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		maxLatitude, err := strconv.ParseFloat(parameter["maxLatitude"], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		minLongitude, err := strconv.ParseFloat(parameter["minLongitude"], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		maxLongitude, err := strconv.ParseFloat(parameter["maxLongitude"], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res, err := fetchShipsForPosition(ctx, db, minLatitude, maxLatitude, minLongitude, maxLongitude)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func fetchShipsForPosition(ctx context.Context, db *generated.Queries, minLatitude, minLongitude, maxLatitude, maxLongitude float64) ([]generated.PositionReport, error) {
	log.Println("Fetching ships for", minLatitude, minLongitude, maxLongitude)
	postionDataArgs := params.BuildGetPositionDataParams(minLatitude, minLongitude, maxLatitude, maxLongitude)

	data, err := db.GetPositionData(ctx, postionDataArgs)
	if err != nil {
		return nil, err
	}
	return data, nil
}
