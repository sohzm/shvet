package main

import (
	"fmt"
	"github.com/sz47/shvet/arguments"
	"github.com/sz47/shvet/handlers"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := []string{"go run main.go", "gruvbox", "best.jpg"}
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
}
