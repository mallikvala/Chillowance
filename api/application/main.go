package main

import (
	"log"
	"net/http"
	"flag"
	"github.com/dmotylev/nutrition"
	"gopkg.in/mgo.v2"

	"github.com/dicknaniel/Chillowance/api/application/routes" // for Router
	"github.com/dicknaniel/Chillowance/api/application/util"   // for Common stuff
)

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
		log.Printf("MongoDB connection failed, with address '%s'.", Configuration.MongoUrl)
	}
	defer mongoSession.Close()
	
	appContext := AppContext{}
	appContext.mongo.session = mongoSession

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
