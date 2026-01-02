package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mycaller := config{}
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations on the map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "",
			callback:    commandExplore,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		if len(cleanLine) == 1 {
			cmdName := cleanLine[0]
			if cmd, ok := supCom[cmdName]; ok {
				cmd.callback(&mycaller, "")
			} else {
				fmt.Println("Unknown command")
			}
		}
		if len(cleanLine) > 1 {
			cmdName := cleanLine[0]
			option := cleanLine[1]
			if cmd, ok := supCom[cmdName]; ok {
				cmd.callback(&mycaller, option)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
