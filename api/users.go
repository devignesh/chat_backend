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

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user schema.Users

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	if err = db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(&user)
	json.NewEncoder(w).Encode(&user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user []schema.Users
	err := db.Find(&user)
	if err == nil {
		log.Fatal(err)
	}
	fmt.Println("{}", user)

	json.NewEncoder(w).Encode(&user)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Print("user test", params)
	var user schema.Users
	err := db.First(&user, params["id"])

	if err == nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println("test1", params)

	var user schema.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("test1fkjdnckjdsnc")
	fmt.Println(user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	db.Debug().Model(&schema.Users{}).Where("id = ?", params["id"]).Take(&schema.Users{}).UpdateColumns(
		map[string]interface{}{
			"name":       user.Name,
			"updated_at": time.Now(),
		},
	)

	if err == nil {
		fmt.Println("updated successfully")
	}

	db.First(&user, params["id"])

	// db.Save(&chatroom)

	json.NewEncoder(w).Encode(&user)

	fmt.Println("sdcsdcscsdcsc")
	fmt.Println("vicky", user)

}

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
