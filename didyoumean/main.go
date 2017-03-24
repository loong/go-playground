package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ld(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}
	}
	return d[len(s)][len(t)]
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/articles/", handler)
	r.HandleFunc("/articles/comments", handler)
	r.HandleFunc("/articles/likes", handler)
	//r.HandleFunc("/articles/{id}", handler)
	//r.HandleFunc("/products/{id}/stuff", handler)

	http.Handle("/", NewGorilla(r))
	http.ListenAndServe(":8080", nil)
}
