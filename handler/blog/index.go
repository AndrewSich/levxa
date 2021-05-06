package blog

import "net/http"

// Handler Blog
func HandlerBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to BLOG Page Levxa"))
}
