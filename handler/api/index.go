package api

import "net/http"

// Handler Api
func HandlerApi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API Page Levxa"))
}
