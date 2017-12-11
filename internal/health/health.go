package health

import (
	"fmt"
	"net/http"
)

// Ping returns "Pong"
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}
