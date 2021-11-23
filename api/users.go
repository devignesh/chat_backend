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

//function for creating new users
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user schema.Users

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//db.create() insert the data into tables
	if err = db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(&user)
	json.NewEncoder(w).Encode(&user)

}

//Get the all list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user []schema.Users
	//db.first() do the select all the data from table
	err := db.Find(&user)
	if err == nil {
		log.Fatal(err)
	}
	fmt.Println("{}", user)

	json.NewEncoder(w).Encode(&user)

}

//get the individual record based on the id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Print("user test", params)
	var user schema.Users
	//select the record from which id is matched with the requested id
	err := db.First(&user, params["id"])

	if err == nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&user)
}

//update the user table information based on id
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println("test1", params)

	var user schema.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//update cndition based on id and Take return a record that match given conditions
	if err = db.Debug().Model(&schema.Users{}).Where("id = ?", params["id"]).Take(&schema.Users{}).UpdateColumns(
		map[string]interface{}{
			"name":       user.Name,
			"updated_at": time.Now(),
		},
	).Error; err != nil {
		log.Fatal(err)
	}

	if err == nil {
		fmt.Println("updated successfully")
	}

	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)

}

//Delete the user based on the id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Print("user test", params)

	var user schema.Users
	if err = db.First(&user, params["id"]).Error; err != nil {
		log.Fatal(err)
	}
	err := db.Delete(&user)
	if err == nil {
		log.Println("successfully deleted")
	}

	json.NewEncoder(w).Encode(&user)
}
