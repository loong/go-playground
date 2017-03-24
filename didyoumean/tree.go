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
	//p := strings.TrimPrefix(path, "/")
	//return strings.Split(p, "/")

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

	for _, p := range parts {
		fmt.Print(p, "-")
	}
	fmt.Println()

	curr := t.root
	for _, p := range parts {
		next, ok := curr.children[p]
		if !ok {
			next = &node{p, make(map[string]*node)}
			curr.children[p] = next
		}

		curr = next
	}
	t.root.Print("")
}

func (t *PathTree) Suggest(path string) (string, bool) {
	parts := splitPath(path)

	if len(t.root.children) == 0 {
		log.Fatal("TODO: No paths added yet")
	}

	suggested := ""

	curr := t.root
	for _, p := range parts {
		next_sug := ""
		next, ok := curr.children[p]
		if !ok {
			min := 100
			for k, v := range curr.children {

				// This is good for the following
				// case:
				//
				// Let there be /likes and
				// /comments. If we request /c, this
				// function will suggest /likes due to
				// the fact that comments edit
				// distance is larger due to the
				// length. Hence we cut the lenght to
				// a similar size first.
				comp := k
				if len(k) > len(p) {
					comp = k[:len(p)]
				}

				dist := ld(comp, p)
				if dist < min {
					min = dist
					next = v
					next_sug = k
				}
			}
		}
		if next == nil {
			return "", false
		}
		if next_sug != "" {
			suggested += "/" + next_sug
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
