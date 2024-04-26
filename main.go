package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commnads!\n")
	repl.Start(os.Stdin, os.Stdout)
}
