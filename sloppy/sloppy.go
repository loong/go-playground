package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type Sloppy struct {
	http.Handler
	routes *PathTree
}

var gorillaExp = regexp.MustCompile("{.+}")

func NewGorilla(handler *mux.Router) Sloppy {
	routes := []string{}
	handler.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		r := gorillaExp.ReplaceAllString(t, "{}")
		fmt.Println(r)
		routes = append(routes, r)
		return nil
	})

	return New(handler, routes)
}

func New(handler http.Handler, routes []string) Sloppy {
	tree := NewPathTree()
	for _, route := range routes {
		tree.AddPath(route)
	}

	return Sloppy{
		Handler: handler,
		routes:  tree,
	}
}

type interceptResponseWriter struct {
	http.ResponseWriter
	*PathTree
	uri    string
	status int
}

func (w *interceptResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *interceptResponseWriter) Write(buf []byte) (int, error) {
	// Proceed as usual
	if w.status != 404 {
		return w.ResponseWriter.Write(buf)
	}

	suggested, ok := w.PathTree.Suggest(w.uri)
	if ok {
		return w.ResponseWriter.Write([]byte(`Did you mean ` + suggested))
	} else {
		return w.ResponseWriter.Write(buf)
	}
}

// ServeHTTP intercepts requests before passing it on to the actual
// handler function
func (s Sloppy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	interceptor := &interceptResponseWriter{
		ResponseWriter: w,
		PathTree:       s.routes,
		uri:            req.RequestURI,
	}
	s.Handler.ServeHTTP(interceptor, req)
}
