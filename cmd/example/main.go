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
	codes, err := rd.List()
	check(err)
	for _, code := range *codes {
		fmt.Printf("Key: %s\n", code.Key)
		fmt.Printf("\t=> Seret: %s\n", code.Code)
	}
	// get one key
	item, err := rd.Read("sample")
	check(err)
	fmt.Printf("Sample key: %+v\n", item)
}
