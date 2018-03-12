package main

//TODO replace dot import
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/config"
	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/dao"
	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/models"
)

// Save the configuration of mongodatabase (localhost and which db use) in Config array
var config = Config{}
var dao = NumObjectDAO{}

// GET list of all objects
func AllObjectsEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// GET an Object by its ID
func FindObjectEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	numObject, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Object ID")
		return
	}
	respondWithJson(w, http.StatusOK, numObject)
}

// POST a new object
func AddObjectEndPoint(w http.ResponseWriter, r *http.Request) {
	//A defer statement defers the execution of a function until the surrounding function returns.
	defer r.Body.Close()
	var numObject Numobject
	if err := json.NewDecoder(r.Body).Decode(&numObject); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	numObject.ID = bson.NewObjectId()
	if err := dao.Insert(numObject); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, numObject)
}

// Delete an object
func DeleteObjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var numObject Numobject
	if err := json.NewDecoder(r.Body).Decode(&numObject); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(numObject); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Update an object
func UpdateObjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var numObject Numobject
	if err := json.NewDecoder(r.Body).Decode(&numObject); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(numObject); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//	Respond Methods
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
	r.HandleFunc("/api/v1/trivia", AllObjectsEndPoint).Methods("GET")
	r.HandleFunc("/api/v1/trivia", AddObjectEndPoint).Methods("POST")
	r.HandleFunc("/api/v1/trivia", DeleteObjectEndPoint).Methods("DELETE")
	r.HandleFunc("/api/v1/trivia", UpdateObjectEndPoint).Methods("PUT")
	r.HandleFunc("/api/v1/trivia/{id}", FindObjectEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
