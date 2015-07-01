package util

import "gopkg.in/mgo.v2"

type AppContext struct {
	Mongo      MongoContext
}

type MongoContext struct {
	Session *mgo.Session
}

var Configuration struct {
	MongoUrl                   string
	MongoDbName                string	
}