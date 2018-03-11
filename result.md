#codding challenge Intercloud : 

##Build Restful API in Go and MongoDB

##Objectives :

 * Import and Use db.JSON in Mongo
 * Deploy working API using Go language
 * Document endpoints from : /api/v1/trivia
 * API should populate database and apply filter on it

##Walkthrough : 

 `Initial commit`

  * Import and Use db.JSON in Mongo
  
  $ mongoimport --db goChallenge --collection numbersapi --drop --file /Users/gbeaumont/Documents/Projects/go/src/github.com/sandaleRaclette/03-coddingChallengeIntercloud/db/db.json --jsonArray

  database => goChallenge

  collection => numbersapi

  * Deploy working API using Go language

  => following tutorial/model : https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94

  - Using the API to listen on port 3000 (main.go):

  $ go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux

  toml : Parse the configuration file (MongoDB server & credentials)
  mux : Request router and dispatcher for matching incoming requests to their respective handler
  mgo : MongoDB driver

  `getting all array in db.JSON after import`

  - Using the API to read the dataset imported in MongoDB
    
    /config/config.toml describe the database

    /dao/numObject_dao.go acces to the db.JSON array in mongoDB by connecting (mgo.Dial(m.Server)) and finding all object define in /models/numObject.go

  `getting all array in db.JSON after import`


    