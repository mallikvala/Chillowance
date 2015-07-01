package main

type Chore struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Points	  int       `json:"points"`
}

type Chores []Chore
