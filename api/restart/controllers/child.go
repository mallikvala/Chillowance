package controllers

import (  
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/dicknaniel/Chillowance/api/restart/models"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type (  
    // UserController represents the controller for operating on the User resource
    ChildController struct{
        session *mgo.Session
    }
)

func NewChildController(s *mgo.Session) *ChildController {  
    return &ChildController{s}
}

// GetUser retrieves an individual user resource
func (cc ChildController) GetChild(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Stub user
    c := models.Child{}

    // Fetch user
    if err := cc.session.DB("test").C("CHILDREN").FindId(oid).One(&c); err != nil {
        w.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    cj, _ := json.Marshal(c)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", cj)
}

// CreateUser creates a new user resource
func (cc ChildController) CreateChild(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Stub an user to be populated from the body
    c := models.Child{}

    // Populate the user data
    json.NewDecoder(r.Body).Decode(&c)

    // Add an Id
    c.Id = bson.NewObjectId()

    // Write the user to mongo
    cc.session.DB("test").C("CHILDREN").Insert(c)

    // Marshal provided interface into JSON structure
    cj, _ := json.Marshal(c)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", cj)
}

// RemoveUser removes an existing user resource
func (cc ChildController) RemoveChild(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // TODO: only write status for now
    w.WriteHeader(200)
}