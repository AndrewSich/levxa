package main

import (
	"log"
	"net/http"
)

func main() {

	// mux routing
	http.HandleFunc("/", handlerIndex)

	// run server
	log.Printf("Server Running on Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}

// Handler Index Page
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home Page Levxa"))
}
