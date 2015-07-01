package models

import "gopkg.in/mgo.v2/bson"

type (  
    // User represents the structure of our resource
    Child struct {
    	Id 		bson.ObjectId	`json:"id" bson:"_id"`
        Name	string			`json:"name" bson:"name"`
        Points	int				`json:"points" bson:"points"`
    }
)