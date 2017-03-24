package main

import (
	"net/http"
	"strings"
)

const WildcardMarker = "{}"

type Sloppy struct {
	http.Handler
	routes []*route
}

func New(handler http.Handler, uris []string) Sloppy {
	res := Sloppy{
		Handler: handler,
		routes:  []*route{},
	}

	for _, route := range uris {
		res.addRoute(route)
	}

	return res
}

type route struct {
	uri       string
	wildcards []int
}

func (s *Sloppy) addRoute(uri string) {
	// find wildcards
	wildcards := []int{}
	splits := strings.Split(uri, "/")

	for i, part := range splits {
		if part == WildcardMarker {
			wildcards = append(wildcards, i)
		}
	}

	s.routes = append(s.routes, &route{
		uri:       uri,
		wildcards: wildcards,
	})
}

func (s *Sloppy) suggest(uri string) (string, bool) {
	min := 100
	suggested := ""
	for _, route := range s.routes {
		dist := getEditDist(route, uri)
		if dist < min {
			min = dist
			suggested = route.uri
		}
	}

	return suggested, true
}

func getEditDist(route *route, uri string) int {
	// apply wildcards
	splits := strings.Split(uri, "/")

	for i := range route.wildcards {
		splits[i] = WildcardMarker
	}

	uri = strings.Join(splits, "/")

	// compute edit distance
	return LevenshteinDist(route.uri, uri)
}

type interceptResponseWriter struct {
	http.ResponseWriter
	Sloppy

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

	suggested, ok := w.suggest(w.uri)
	if ok {
		return w.ResponseWriter.Write([]byte(`Did you mean ` + suggested))
	} else {
		return w.ResponseWriter.Write(buf)
	}

	return w.ResponseWriter.Write(buf)
}

// ServeHTTP intercepts requests before passing it on to the actual
// handler function
func (s Sloppy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	interceptor := &interceptResponseWriter{
		ResponseWriter: w,
		Sloppy:         s,
		uri:            req.RequestURI,
	}
	s.Handler.ServeHTTP(interceptor, req)
}
