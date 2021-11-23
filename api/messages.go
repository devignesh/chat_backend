package api

import (
	"chat_backend/schema"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//create a new message fucntion
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message schema.Messages
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	if err = db.Create(&message).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(&message)
	json.NewEncoder(w).Encode(&message)

}

//get the list of all records from table
func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message []schema.Messages
	err := db.Find(&message)
	if err == nil {
		log.Fatal(err)
	}

	fmt.Println("{}", message)
	json.NewEncoder(w).Encode(&message)

}

//get the sing record from the table based on the id
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message schema.Messages
	err := db.First(&message, params["id"])
	if err == nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&message)
}

//update the record based on the id
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message schema.Messages
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	if err = db.Debug().Model(&schema.Messages{}).Where("id = ?", params["id"]).Take(&schema.Messages{}).UpdateColumns(
		map[string]interface{}{
			"name":       message.Content,
			"updated_at": time.Now(),
		},
	).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated successfully")

	db.First(&message, params["id"])
	json.NewEncoder(w).Encode(&message)

}

//delete the record
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message schema.Messages
	if err = db.First(&message, params["id"]).Error; err != nil {
		log.Fatal(err)

	}

	err := db.Delete(&message)
	if err == nil {
		log.Println("successfully deleted")
	}

	json.NewEncoder(w).Encode(&message)
}
