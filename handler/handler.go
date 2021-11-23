package handler

import (
	"chat_backend/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	//new mux router
	router := mux.NewRouter()

	//users routes
	router.HandleFunc("/users", api.CreateUser).Methods("POST")
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", api.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", api.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")

	//chatroom routes
	router.HandleFunc("/chatrooms", api.CreateChatroom).Methods("POST")
	router.HandleFunc("/chatrooms", api.GetChatrooms).Methods("GET")
	router.HandleFunc("/chatroom/{id}", api.GetChatroom).Methods("GET")
	router.HandleFunc("/chatroom/{id}", api.UpdateChatroom).Methods("PUT")
	router.HandleFunc("/chatroom/{id}", api.DeleteChatroom).Methods("DELETE")

	//message routes
	router.HandleFunc("/messages", api.CreateMessage).Methods("POST")
	router.HandleFunc("/messages", api.GetMessages).Methods("GET")
	router.HandleFunc("/message/{id}", api.GetMessage).Methods("GET")
	router.HandleFunc("/message/{id}", api.UpdateMessage).Methods("PUT")
	router.HandleFunc("/message/{id}", api.DeleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
