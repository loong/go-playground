package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Envelope struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

func handler(text string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		resp := make(map[string]string)
		resp["text"] = text

		env := Envelope{
			Meta: Meta{Code: 200},
			Data: resp,
		}

		js, _ := json.MarshalIndent(env, "", "   ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler("Hello there! Try out /v3/tracking"))
	r.HandleFunc("/v4/trackings", handler("Great! How about this /v4/courier/sall"))
	r.HandleFunc("/v4/couriers", handler("Great! How about this /v4/courier/sall"))
	r.HandleFunc("/v4/couriers/all", handler("How about dynamic urls? /v4/notifications/123/456/ad"))
	r.HandleFunc("/v4/notifications/{slug}/{tracking_number}/add", handler("Hope you liked it!"))

	http.Handle("/", NewGorilla(r))
	http.ListenAndServe(":8080", nil)
}
