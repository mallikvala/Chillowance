package main

type Child struct {
	childId   string	`bson:"_id" json:"storyId"`
	Name      string    `json:"name"`
	Points	  int       `json:"points"`
}

type Children []Child
