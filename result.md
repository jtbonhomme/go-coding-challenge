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
  
  mongoimport --db goChallenge --collection numbersapi --drop --file /Users/gbeaumont/Documents/Projects/go/src/github.com/sandaleRaclette/03-coddingChallengeIntercloud/db/db.json --jsonArray

  database => goChallenge

  collection => numbersapi

  * Deploy working API using Go language

  => following tutorial/model : https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94

  