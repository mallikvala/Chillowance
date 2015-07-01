package main

import (  
    // Standard library packages
    "net/http"

    //mgo
    "gopkg.in/mgo.v2"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/dicknaniel/Chillowance/api/restart/controllers"
)

func main() {  
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    cc := controllers.NewChildController(getSession())

    // Get a user resource
    r.GET("/child/:id", cc.GetChild)

    r.POST("/child", cc.CreateChild)

    r.DELETE("/child/:id", cc.RemoveChild)

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {  
    // Connect to our local mongo
    s, err := mgo.Dial("192.168.59.103:27017")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}