package main

import (
	"fmt"
	"os"
	"shvet/handlers"
)

func main() {
	var err error
	args := os.Args

	for i, val := range args {
		switch val {
		case "darken":
			err = handlers.HandleDarken(args[i+1], args[i+2])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "lighten":
			err = handlers.HandleLighten(args[i+1], args[i+2])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "gruvbox":
			err = handlers.HandleGruvbox(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "gruvbox2":
			if (len(args) == i + 2) {
				err = handlers.HandleGruvbox2("50", args[i+1])
			} else {
				err = handlers.HandleGruvbox2(args[i+1], args[i+2])
			}
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "nord":
			err = handlers.HandleNord(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "nord2":
			if (len(args) == i + 2) {
				err = handlers.HandleNord2("50", args[i+1])
			} else {
				err = handlers.HandleNord2(args[i+1], args[i+2])
			}
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "tokyonight":
			err = handlers.HandleTokyonight(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "tokyonight2":
			if (len(args) == i + 2) {
				err = handlers.HandleTokyonight2("50", args[i+1])
			} else {
				err = handlers.HandleTokyonight2(args[i+1], args[i+2])
			}
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return

		case "avgvalues":
			err = handlers.HandleAvgValues(args[i+1])
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			return
		}
	}
}
