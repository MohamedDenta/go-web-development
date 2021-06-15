package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You have requested the book :%s on page %s\n", title, page)
	})
	// .Host("www.moha.com")
	// r.Schemes("http")
	http.ListenAndServe(":8080", r)
}
