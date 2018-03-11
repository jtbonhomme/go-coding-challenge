package main

//TODO replace dot import
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/config"
	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/dao"
)

// Save the configuration of mongodatabase (localhost and which db use) in Config array
var config = Config{}
var dao = NumObjectDAO{}

// GET list of all objects
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes and define the differents endpoints
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", AllMoviesEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
