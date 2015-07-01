package models
import "time"

type Person struct {
	PersonId    int	`bson:"_id" json:"personId"`
	Role 		string `json:"role"`
	Type 		string `json:"type"`
	DOB			time.Time `json:"dob"`
	Name 		string `json:"name"`
	EmailAddress string `json:"emailAddress"`
}
