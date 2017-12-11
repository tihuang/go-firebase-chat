package server

import (
	"net/http"

	"github.com/tihuang/go-firebase-chat/internal/health"
	goji "goji.io"
	"goji.io/pat"
)

// Server returns the entire application server
type chatServer struct {
	*goji.Mux
}

// NewChatServer instantiates a new chat server
func NewChatServer() http.Handler {
	root := goji.NewMux()
	debug := goji.SubMux()
	root.Handle(pat.New("/health/*"), debug)

	debug.HandleFunc(pat.Get("/ping"), health.Ping)
	return root
}
