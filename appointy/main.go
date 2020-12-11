package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arjunChouksey/appointy-task/appointy/helper"
	"github.com/arjunChouksey/appointy-task/appointy/models"
	"github.com/arjunChouksey/appointy-task/appointy/new_helper"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collectionUser = helper.ConnectDB()
var collectionContact = new_helper.ConnectDB()

func getUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collectionUser.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our book model.
	result, err := collectionUser.InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&contact)

	// insert our book model.
	result, err := collectionContact.InsertOne(context.TODO(), contact)

	if err != nil {
		new_helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// var client *mongo.Client

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	//r.HandleFunc("/contacts?user=<user id>&infection_timestamp=<timestamp>", listContact).Methods("GET")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
