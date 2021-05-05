package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// mux routing
	http.HandleFunc("/", handlerIndex)

	// run server
	log.Printf("Server Running On\n")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Println(err.Error())
	}
}

// Get Port
func GetPort() string {
	var port = os.Getenv("PORT")

	// get default port if port not found
	if port == "" {
		port = "4747"
		log.Println("[!] No PORT env variable detect, default to " + port)
	}

	return ":" + port
}

// Handler Index Page
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home Page Levxa"))
}
