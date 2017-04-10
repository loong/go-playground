package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RandomStringsFromArray picks n elements form an array. Caution: This
// function changes the order of the input array.
func RandomStringsFromArray(list []string, n int) ([]string, error) {
	if n >= len(list) {
		return nil, errors.New("Cannot pick more elements than available")
	}

	for i := 0; i < n; i++ {
		randpos := rand.Intn(len(list))
		list[i], list[randpos] = list[randpos], list[i]
	}

	return list[0:n], nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arr := []string{"0", "1", "2", "3", "4", "5"}

	for i := 0; i < 10; i++ {
		elems, _ := RandomStringsFromArray(arr[:], 3)
		fmt.Println(arr, elems)
	}
}
