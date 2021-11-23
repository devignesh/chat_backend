package api

import (
	"chat_backend/schema"
	"chat_backend/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB
var err error

//database connection using gorm.io/gorm
func DbConn() {
	dsn := "host=localhost port=5432 user=postgres dbname=chatbackend sslmode=disable password=postgres"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=chatbackend sslmode=disable password=postgres")
	if err != nil {
		log.Println("failed to connect database", err)
	}

	log.Println("db connected successfully")

	//Automigrate use to migrate the go tables
	err := db.AutoMigrate(&schema.Chatroom{}, &schema.Users{}, &schema.Messages{})
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

//create a new chatroom, insert the new record in the chatroom table
func CreateChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//initialize the chatroom req types
	var req types.ChatroomReq
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	var chatroom schema.Chatroom
	//assign the roomname
	chatroom.RoomName = req.RoomName
	if err = db.Create(&chatroom).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(&chatroom)
	//initialize the empty chatroomuser list for join table
	var chatroomusers []schema.ChatroomUsers

	//load the all userid and append the chatroomid and userid into chatroomusers table
	for _, v := range req.UserIDs {
		chatroomusers = append(chatroomusers, schema.ChatroomUsers{ChatroomID: chatroom.ID, UsersID: v})
	}

	fmt.Println(chatroomusers)

	//create the new record
	if err = db.Create(&chatroomusers).Error; err != nil {
		log.Fatal(err)
	}

	var chatrooms schema.Chatroom
	//select the chatroom records based on the chatroomid and it preloads the users information from chatroomusrs
	db.Preload("Users").First(&chatrooms, chatroom.ID)

	json.NewEncoder(w).Encode(&chatrooms)

}

//get the list of chatrooms in the table
func GetChatrooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chatrooms []schema.Chatroom

	//find all the records from table with preload the users records
	if err = db.Preload("Users").Find(&chatrooms).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("{}", chatrooms)

	json.NewEncoder(w).Encode(&chatrooms)

}

//get the single record of chatroom based on the id
func GetChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var chatroom schema.Chatroom

	//fetch the matching record based on the request id with preloads the usrs information
	if err = db.Preload("Users").First(&chatroom, params["id"]).Error; err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&chatroom)
	if err != nil {
		log.Fatal(err)
	}

}

//update the chatroom record based on the id
func UpdateChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var chatroom schema.Chatroom

	err := json.NewDecoder(r.Body).Decode(&chatroom)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//update the record based on the id and Take return a record that match given conditions
	if err = db.Debug().Model(&schema.Chatroom{}).Where("id = ?", params["id"]).Take(&schema.Chatroom{}).UpdateColumns(
		map[string]interface{}{
			"room_name":  chatroom.RoomName,
			"updated_at": time.Now(),
		},
	).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated successfully")

	//after updation it gives the response along with users record
	if err = db.Preload("Users").First(&chatroom, params["id"]).Error; err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&chatroom)

}

//delete chatroom function
func DeleteChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var chatroom schema.Chatroom

	if err = db.First(&chatroom, params["id"]).Error; err != nil {
		log.Fatal(err)

	}

	db.Select(clause.Associations).Delete(&chatroom)
	json.NewEncoder(w).Encode(&chatroom)

}
