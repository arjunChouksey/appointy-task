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
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}

type DOB struct {
	Date  int32  `json:"date,omitempty" bson:"date,omitempty"`
	Month string `json:"month,omitempty" bson:"month,omitempty"`
	Year  int32  `json:"year,omitempty" bson:"year,omitempty"`
}

type Name struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type Contact struct {
	Key       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserIdOne string             `json:"_id1,omitempty" bson:"_id1,omitempty"`
	UserIdTwo string             `json:"_id2,omitempty" bson:"_id2,omitempty"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
