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

    $ mongod

    * POSTMAN :

    GET http://localhost:3000/api/v1/trivia

  `Add collection methods POST, PUT, DELETE`

- Using the API to add, delete and update an Object in the dataSet

  /dao/numObject_dao.go add Collection Methods (https://docs.mongodb.com/manual/reference/method/js-collection/)

  * POSTMAN :
    
    POST http://localhost:3000/api/v1/trivia
    body(raw) :
    {
      "text":"Neuilly-Plaisance",
      "number":93360,
      "found":true,
      "type":"trivia"
    }
    =>{
          "id": "5aa6dce509f17a0581553518",
          "text": "Neuilly-Plaisance",
          "number": 93360,
          "found": true,
          "type": "trivia"
      }

    PUT http://localhost:3000/api/v1/trivia
    body(raw) :
    {
        "id": "5aa6dce509f17a0581553518",
        "text": "Neuilly-Plaisance my lovely city",
        "number": 93360,
        "found": true,
        "type": "trivia"
    }
    =>{
          "result": "success"
      }

    GET http://localhost:3000/api/v1/trivia/5aa6dce509f17a0581553518
    =>{
          "id": "5aa6dce509f17a0581553518",
          "text": "Neuilly-Plaisance my lovely city",
          "number": 93360,
          "found": true,
          "type": "trivia"
      }

    DELETE http://localhost:3000/api/v1/trivia
    body(raw) :
    {
        "id": "5aa6dce509f17a0581553518",
        "text": "Neuilly-Plaisance my lovely city",
        "number": 93360,
        "found": true,
        "type": "trivia"
    }
    =>{
        "result": "success"
      }


  `Implementation query and filter by number`

- Implement filter by query on Number

  blocking to query with a number => Add post on slack : 
  
  https://stackoverflow.com/questions/49225624/get-object-by-specific-field/49225678?noredirect=1#comment85456271_49225678

##TODO

- Change dot import to classic import
- Debbug Find by Number Function
- Make swagger.json