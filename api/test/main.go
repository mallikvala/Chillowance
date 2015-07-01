package main

import (
        "log"
        "math/rand"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type AppContext struct {
        mongo      MongoContext
}

type MongoContext struct {
        session *mgo.Session
}

type Child struct {
        id      bson.ObjectId   `json:"id" bson:"_id" `
        name    string          `json:"name" bson:"name"`
        points  int             `json:"points" bson:"points"`
}

const (
        childrenCollection   = "CHILDREN"
)

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// generates a random string of fixed size
func srand(size int) string {
        buf := make([]byte, size)
        for i := 0; i < size; i++ {
                buf[i] = alpha[rand.Intn(len(alpha))]
        }
        return string(buf)
}

// GetNewId returns random 5 character string.
// Random characters are [a-zA-Z0-9] excluding those with visible similarity = l,1,O,0.
// This gives 550 731 776 unique id's.
func GetNewId() string {
        return srand(5)
}

func info(template string, values ...interface{}) {
        log.Printf("[chillowance] "+template+"\n", values...)
}


func (m *MongoContext) insertChild(child *Child) (id string, err error) {
        mongoSession := m.session.Clone()
        defer mongoSession.Close()

        c := mongoSession.DB("test").C(childrenCollection)

        child.id = bson.NewObjectId()
        
        log.Printf("id %s", id)
        log.Printf("childId %s", child.id)
        log.Printf("name %s", child.name)

        err = c.Insert(child)
        if mgo.IsDup(err) {
                // retry insert with new id
                m.insertChild(child)
        }
        return
}

func main() {
        mongoSession, _ := mgo.Dial("192.168.59.103:27017")
        defer mongoSession.Close()
        
        appContext := AppContext{}
        appContext.mongo.session = mongoSession

        child := Child{}
        child.name = "Nick"
        appContext.mongo.insertChild(&child)


        // err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	       //         &Person{"Cla", "+55 53 8402 8510"})
        // if err != nil {
        //         log.Fatal(err)
        // }

        // result := Person{}
        // err = c.Find(bson.M{"name": "Ale"}).One(&result)
        // if err != nil {
        //         log.Fatal(err)
        // }

        // fmt.Println("Phone:", result.Phone)
}