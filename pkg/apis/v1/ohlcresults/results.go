package ohlcresults

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sp98/resultstore/pkg/store"
)

/*
TODO: Send proper error messages
*/
var (
	//DBUrl is the connection url for influx db
	DBUrl = ""
	//DBName is the database name
	DBName = ""
)

func init() {
	DBUrl = os.Getenv("MONGO_DB_URL")
	DBName = os.Getenv("ANALYSIS_RESULT_DB_NAME")
}

//Routes define the OHCL routes
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{interval}", GetOHLCResults)
	router.Post("/{interval}", StoreOHLCResult)
	return router
}

//StoreOHLCResult stores the ohlc result received from the analyser
func StoreOHLCResult(w http.ResponseWriter, r *http.Request) {
	interval := chi.URLParam(r, "interval")
	var result store.Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		log.Printf("error getting the last ohlc result. %+v", err)
		return
	}
	db := store.NewDB(DBUrl, DBName, fmt.Sprintf("ohlc-%s", interval))
	err = db.InsertOHLCResult(&result)
	if err != nil {
		log.Printf("error getting the last ohlc result. %+v", err)
	}
}

//GetOHLCResults gets the last stored ohlc analysis results
func GetOHLCResults(w http.ResponseWriter, r *http.Request) {
	interval := chi.URLParam(r, "interval")
	db := store.NewDB(DBUrl, DBName, fmt.Sprintf("ohlc-%s", interval))
	result, err := db.GetOHLCResult()
	if err != nil {
		log.Printf("error getting the last ohlc result. %+v", err)

	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	render.JSON(w, r, result) // A chi router helper for serializing and returning json
}
