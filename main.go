package main

import (
	"fmt"
	"os"
	"shvet/arguments"
	"shvet/handlers"
)

func main() {
	args := os.Args
	parsedArgs, err := arguments.ParseArguments(args[1:])

	if err != nil {
		fmt.Println("Error:", err.Error())
		fmt.Println()
		fmt.Println("You're using Shvet wrong,")
		fmt.Println("for help see `shvet help`")
	} else {
		handlers.Handle(parsedArgs)
	}
}
