package main

import (
	"net/http"

	"github.com/tihuang/go-firebase-chat/app/server"
)

func main() {
	chatServer := server.NewChatServer()
	http.ListenAndServe("localhost:8000", chatServer)
}
