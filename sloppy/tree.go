package main

import (
	"fmt"
	"log"
	"strings"
)

type PathTree struct {
	root *node
}

func NewPathTree() *PathTree {
	return &PathTree{
		root: &node{
			children: make(map[string]*node),
		},
	}
}

func splitPath(path string) []string {
	parts := []string{}
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			parts = append(parts, p)
		}
	}

	return parts
}

func (t *PathTree) AddPath(path string) {
	parts := splitPath(path)

	curr := t.root
	for _, p := range parts {
		next, ok := curr.children[p]
		if !ok {
			next = &node{p, make(map[string]*node)}
			curr.children[p] = next
		}

		curr = next
	}
}

func suggestNext(part string, n *node) (string, *node) {
	suggestion := part

	// if this part of the path is correct, move on to the next
	next, ok := n.children[part]
	if ok {
		return suggestion, next
	}

	min := 100
	for k, v := range n.children {
		// This is good for the following case:
		//
		// Let there be /likes and /comments. If we request
		// /c, this function will suggest /likes due to the
		// fact that comments edit distance is larger due to
		// the length. Hence we cut the lenght to a similar
		// size first.
		comp := k
		if len(k) > len(part) {
			comp = k[:len(part)]
		}

		dist := LevenshteinDist(comp, part)
		if dist < min {
			min = dist
			next = v
			suggestion = k
		}
	}

	return suggestion, next
}

func (t *PathTree) Suggest(path string) (string, bool) {
	parts := splitPath(path)

	if len(t.root.children) == 0 {
		log.Fatal("TODO: No paths added yet")
	}

	suggested := ""

	curr := t.root
	for _, p := range parts {
		nextSuggestion, next := suggestNext(p, curr)

		if next == nil {
			return "", false
		}

		if nextSuggestion != "" {
			suggested += "/" + nextSuggestion
		}
		curr = next
	}

	return suggested, true
}

type node struct {
	value    string
	children map[string]*node
}

func (n *node) Print(prefix string) {
	fmt.Println(prefix, n.value+"|")

	for _, v := range n.children {
		v.Print(prefix + "   ")
	}
}
