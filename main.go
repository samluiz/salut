package main

import (
	"fmt"
	"os"
	"os/user"
	"salut/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the Salut console, %s!\n", user.Username)
	fmt.Printf("You may start to type the commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}