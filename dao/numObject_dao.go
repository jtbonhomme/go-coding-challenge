package dao

import (
	"log"

	. "github.com/sandaleRaclette/03-coddingChallengeIntercloud/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NumObjectDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// The collection use in MongoDB database define in config.toml
const (
	COLLECTION = "numbersapi"
)

// Establish a connection to database
func (m *NumObjectDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of object
func (m *NumObjectDAO) FindAll() ([]Numobject, error) {
	var numObject []Numobject
	err := db.C(COLLECTION).Find(bson.M{}).All(&numObject)
	return numObject, err
}
