package main

import (
	"chat_backend/api"
	"chat_backend/handler"

	"fmt"
)

func main() {

	fmt.Println("Chat Backend server")

	//db connection
	api.DbConn()

	//api handler routes
	handler.Routes()
}
