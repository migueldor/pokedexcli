package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	supCom = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		if len(cleanLine) > 0 {
			cmdName := cleanLine[0]
			if cmd, ok := supCom[cmdName]; ok {
				cmd.callback()
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}
