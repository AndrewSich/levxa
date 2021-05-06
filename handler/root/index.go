package root

import "net/http"

// Handler Main
func HandlerMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to INDEX Page Levxa"))
}
