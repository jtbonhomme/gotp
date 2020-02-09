package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
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
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Key", "Time-based OTP"})

		for i, code := range *codes {
			t.AppendRow([]interface{}{i + 1, code.Key, code.Code})
		}
		t.SetStyle(table.StyleColoredBright)
		t.Render()
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
