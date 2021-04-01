package main

import (
	"fmt"
	"go-monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	cur_user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", cur_user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
