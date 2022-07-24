package main

import (
	"fmt"
	"os"
	"shvet/operations"
)

func main() {
	args := os.Args

	for i, val := range args {
		if val == "darken" {
			err := operations.Darken(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return
		}
		if val == "lighten" {
			err := operations.Lighten(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return
		}
	}
}
