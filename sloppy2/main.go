package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/articles", handler)
	r.HandleFunc("/articles/comments", handler)
	r.HandleFunc("/articles/likes", handler)
	//r.HandleFunc("/articles/{id}", handler)
	r.HandleFunc("/products/{id}/stuff", handler)

	http.Handle("/", NewGorilla(r))
	http.ListenAndServe(":8080", nil)
}
