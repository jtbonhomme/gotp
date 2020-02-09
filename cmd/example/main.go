package main

import (
	"fmt"

	"github.com/jtbonhomme/gotp/backend/random"
)

// panics if error is not nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create a new storage backend for test
	rd := random.New("test")
	// add a new fake key { "myServiceURI", "KZAUYVKFGM======"}
	err := rd.Store("myServiceURI", "KZAUYVKFGM======")
	check(err)
	// fetch random fake codes
	keys, err := rd.List()
	check(err)
	fmt.Printf("keys: %#v\n", keys)
}
