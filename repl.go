package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)
	textSlice := strings.Split(trimmedText, " ")
	for _, s := range textSlice {
		if s != "" && s != " " {
			result = append(result, s)
		}
	}
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supCom map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Printf("help: %s\n", supCom["help"].description)
	fmt.Printf("exit: %s\n", supCom["exit"].description)
	return nil
}
