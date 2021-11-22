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

func DbConn() {
	dsn := "host=localhost port=5432 user=postgres dbname=chatbackend sslmode=disable password=postgres"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=chatbackend sslmode=disable password=postgres")
	if err != nil {
		log.Println("failed to connect database", err)
	}

	log.Println("db connected successfully")

	err := db.AutoMigrate(&schema.Chatroom{}, &schema.Users{}, &schema.Messages{})
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

func CreateChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req types.ChatroomReq
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	var chatroom schema.Chatroom
	chatroom.RoomName = req.RoomName
	if err = db.Create(&chatroom).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(&chatroom)
	var chatroomusers []schema.ChatroomUsers

	for _, v := range req.UserIDs {
		chatroomusers = append(chatroomusers, schema.ChatroomUsers{ChatroomID: chatroom.ID, UsersID: v})
	}

	fmt.Println(chatroomusers)

	if err = db.Create(&chatroomusers).Error; err != nil {
		log.Fatal(err)
	}

	var chatrooms schema.Chatroom
	db.Preload("Users").First(&chatrooms, chatroom.ID)

	json.NewEncoder(w).Encode(&chatrooms)

}

func GetChatrooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chatrooms []schema.Chatroom

	if err = db.Preload("Users").Find(&chatrooms).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("{}", chatrooms)

	json.NewEncoder(w).Encode(&chatrooms)

}

func GetChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// fmt.Println("test vale", params)
	var chatroom schema.Chatroom

	if err = db.Preload("Users").First(&chatroom, params["id"]).Error; err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&chatroom)
	if err != nil {
		log.Fatal(err)
	}

}
func UpdateChatroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var chatroom schema.Chatroom

	err := json.NewDecoder(r.Body).Decode(&chatroom)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	if err = db.Debug().Model(&schema.Chatroom{}).Where("id = ?", params["id"]).Take(&schema.Chatroom{}).UpdateColumns(
		map[string]interface{}{
			"room_name":  chatroom.RoomName,
			"updated_at": time.Now(),
		},
	).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated successfully")

	if err = db.Preload("Users").First(&chatroom, params["id"]).Error; err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&chatroom)

}

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
