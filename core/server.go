package core

import (
	"fmt"
	"net/http"
)

// Run Server Function
func Run() {
	muxHandle()

	fmt.Printf("Server Running On PORT 9000\n")
	err := http.ListenAndServe(":9000", muxCall)
	if err != nil {
		fmt.Println("[!] warning ", err.Error())
	}
}
