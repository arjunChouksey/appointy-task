package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        *Name              `json:"name,omitempty" bson:"name,omitempty"`
	DOB         *DOB               `json:"dob,omitempty" bson:"dob,omitempty"`
	PhoneNumber string             `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}

type DOB struct {
	date  int    `json:"date,omitempty" bson:"date,omitempty"`
	month string `json:"month,omitempty" bson:"month,omitempty"`
	year  int    `json:"year,omitempty" bson:"year,omitempty"`
}

type Name struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type Contact struct {
	UserIdOne primitive.ObjectID `json:"_id1,omitempty" bson:"_id1,omitempty"`
	UserIdTwo primitive.ObjectID `json:"_id2,omitempty" bson:"_id2,omitempty"`
}
