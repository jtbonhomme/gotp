package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jtbonhomme/gotp"
	"github.com/jtbonhomme/gotp/backend/secure"
)

const (
	MinArgNumber            int  = 2
	GoogleTimeIntervaleSeed uint = 30
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

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getKey := getCmd.String("key", "", "key")

	if len(os.Args) < MinArgNumber {
		fmt.Println("expected 'add', 'get', 'del' or 'list' subcommands")
		os.Exit(1)
	}

	// create the GoTP code manager based from secured key ring
	codeMgr := gotp.New(secure.New(user + ":mfa"))

	switch os.Args[1] {
	case "list":
		codes, err := codeMgr.List()
		check(err)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Key"})
		for i, code := range codes {
			t.AppendRow([]interface{}{i + 1, code})
		}
		t.SetStyle(table.StyleColoredBright)
		t.Render()
	case "add":
		err := addCmd.Parse(os.Args[2:])
		check(err)
		err = codeMgr.Store(*addKey, *addValue)
		check(err)
	case "del":
		err := delCmd.Parse(os.Args[2:])
		check(err)
		err = codeMgr.Remove(*delKey)
		check(err)
	case "get":
		err := getCmd.Parse(os.Args[2:])
		check(err)
		totp, err := codeMgr.
			WithTimeIntervalSeed(GoogleTimeIntervaleSeed).
			Get(*getKey)
		check(err)
		fmt.Printf("%s", totp)
		err = clipboard.WriteAll(totp)
		check(err)
	default:
		fmt.Println("expected 'add', 'get', 'del' or 'list' subcommands")
		os.Exit(1)
	}
}
