package main

import (
	"encoding/json"
	"net/http"
	"os"
	"stockcoder/server/db"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func dataHandler(w http.ResponseWriter, r *http.Request) {
	chartData, _ := db.GetData(r.FormValue("cd"))
	json.NewEncoder(w).Encode(chartData)
}

func volumeHandler(w http.ResponseWriter, r *http.Request) {
	volume, _ := db.GetVolumeData(r.FormValue("cd"))
	json.NewEncoder(w).Encode(volume)
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	date, _ := db.GetDateData(r.FormValue("cd"))
	json.NewEncoder(w).Encode(date)
}

func main() {
	// allowedOrigins := handlers.AllowedOrigins([]string{"https://stockcoders.appspot.com"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})

	godotenv.Load()
	router := mux.NewRouter()
	router.HandleFunc("/data/", dataHandler)
	router.HandleFunc("/volume/", volumeHandler)
	router.HandleFunc("/date/", dateHandler)
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
}