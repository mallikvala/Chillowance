package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

type AppContext struct {
	mongo      MongoContext
}

type MongoContext struct {
	session *mgo.Session
}

var Configuration struct {
	MongoUrl                   string
	MongoDbName                string	
}


func init() {
	configurationFile := flag.String("conf", "chillowance.cfg", "Full path to configuration file")

	err := nutrition.Env("CHILLOWANCE_").File(*configurationFile).Feed(&Configuration)
	if err != nil {
		log.Fatalf("[chillowance] Unable to read properties:%v\n", err)
	}
}

func main() {

	mongoSession, err := mgo.Dial(Configuration.MongoUrl)
	if err != nil {
		Error.Printf("MongoDB connection failed, with address '%s'.", Configuration.MongoUrl)
	}
	defer mongoSession.Close()
	
	appContext := AppContext{}
	appContext.mongo.session = mongoSession

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
