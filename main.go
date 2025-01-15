package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ahmetgurcand/interpreter/repl"
)
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monker programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

// 2.2 Why not a parser generator? page 30