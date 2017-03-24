package main

import (
	"fmt"
	"regexp"

	"github.com/gorilla/mux"
)

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
