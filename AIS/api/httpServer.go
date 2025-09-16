package api

import (
	"aisstream/db/generated"
	"aisstream/db/params"
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type homeHandler struct{}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func ServeHTTPServer(dbContext context.Context) {
	router := mux.NewRouter()

	// Define endpoints
	router.HandleFunc("/map/{minLatitude}/{maxLatitude}/{minLongitude}/{maxLongitude}", getShipsForPosition).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getShipsForPosition(w http.ResponseWriter, r *http.Request) {
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

	getPostionDataArgs := params.BuildGetPositionDataParams(minLatitude, minLongitude, maxLatitude, maxLongitude)

	data, err := generated.GetPositionDataP

}
