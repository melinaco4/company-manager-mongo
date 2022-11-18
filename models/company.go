package models

import "gopkg.in/mgo.v2/bson"

type Company struct {
	Id                bson.ObjectId `json:"id" bson:"_id"`
	Name              string        `json:"name" bson:"name"`
	Description       string        `json:"description" bson:"description"`
	AmountOfEmployees int           `json:"amount_of_employees" bson:"amount_of_employees"`
	Registered        bool          `json:"registered" bson:"registered"`
	Type              string        `json:"type" bson:"type"`
}
