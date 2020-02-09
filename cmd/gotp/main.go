package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jtbonhomme/gotp/backend/secure"
)

// panics if error is not nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	user := os.Getenv("USER")
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addKey := addCmd.String("key", "", "key")
	addValue := addCmd.String("value", "", "value")

	delCmd := flag.NewFlagSet("del", flag.ExitOnError)
	delKey := delCmd.String("key", "", "key")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'del' or 'list' subcommands")
		os.Exit(1)
	}

	// create the key ring
	secring := secure.New(user + ":mfa")

	switch os.Args[1] {
	case "list":
		// fetch stored codes
		codes, err := secring.List()
		check(err)
		for _, code := range *codes {
			fmt.Printf("Key: %s\n", code.Key)
			fmt.Printf("\t=> Secret: %s\n\n", code.Code)
		}
	case "add":
		err := addCmd.Parse(os.Args[2:])
		check(err)
		err = secring.Store(*addKey, *addValue)
		check(err)
	case "del":
		err := delCmd.Parse(os.Args[2:])
		check(err)
		err = secring.Remove(*delKey)
		check(err)
	default:
		fmt.Println("expected 'add', 'del' or 'list' subcommands")
		os.Exit(1)
	}
}
