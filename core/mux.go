package core

import (
	"fmt"
	"net/http"
	"strings"

	"levxa/handler"
)

// Mux Struct for ServeMux
type Mux struct {
	main, blog, api *http.ServeMux
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// checking if root domain
	if r.Host == "localhost:9000" {
		fmt.Println("[+] success ", r.Host)
		mux.main.ServeHTTP(w, r)
		return
	}

	// checking if subdomain valid
	domainParts := strings.Split(r.Host, ".")
	if domainParts[0] == "blog" {
		fmt.Println("[+] success ", r.Host)
		mux.blog.ServeHTTP(w, r)
	} else if domainParts[0] == "api" {
		fmt.Println("[+] success ", r.Host)
		mux.api.ServeHTTP(w, r)
	} else {
		fmt.Println("[!] warning ", r.Host, " subdomain not found")
		http.Error(w, "Pages Not Found", 404)
	}
}

var muxCall = &Mux{
	main: http.NewServeMux(),
	blog: http.NewServeMux(),
	api:  http.NewServeMux(),
}

func muxHandle() {
	muxCall.main.HandleFunc("/", handler.HandlerMain)
	muxCall.blog.HandleFunc("/", handler.HandlerBlog)
	muxCall.api.HandleFunc("/", handler.HandlerApi)
}
