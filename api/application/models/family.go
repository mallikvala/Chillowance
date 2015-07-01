package models

type Family struct {
	FamilyId   string	`bson:"_id" json:"familyId"`
	Members    []Person    `json:"members"`
}

