package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"./db"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//チャートデータ(json)を返すメソッド
func dataHandler(w http.ResponseWriter, r *http.Request) {
	chartData, _ := db.GetData(r.FormValue("cd"))
	json.NewEncoder(w).Encode(chartData)
}

//出来高データ(json)を返すメソッド
func volumeHandler(w http.ResponseWriter, r *http.Request) {
	volume, _ := db.GetVolumeData(r.FormValue("cd"))
	json.NewEncoder(w).Encode(volume)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	cd := r.FormValue("cd")
	date := r.FormValue("date")
	start, _ := strconv.Atoi(r.FormValue("start"))
	high, _ := strconv.Atoi(r.FormValue("high"))
	low, _ := strconv.Atoi(r.FormValue("low"))
	close, _ := strconv.Atoi(r.FormValue("close"))
	db.AddPriceData(cd, date, start, high, low, close)
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	cd := r.FormValue("cd")
	date := r.FormValue("date")
	db.RemovePriceData(cd, date)
}

func main() {
	//CORS(Cross-Origin Resource Sharing)設定
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})

	//Dotenvを使用
	godotenv.Load()

	//handlerの設定
	router := mux.NewRouter()
	router.HandleFunc("/data/", dataHandler)
	router.HandleFunc("/volume/", volumeHandler)
	router.HandleFunc("/add/", addHandler)
	router.HandleFunc("/remove/", removeHandler)
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
}
