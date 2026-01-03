package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/migueldor/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	mycaller := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.PokemonResponse{},
	}
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
			description: "Shows the pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "",
			callback:    commandPokedex,
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
				cmd.callback(mycaller, "")
			} else {
				fmt.Println("Unknown command")
			}
		}
		if len(cleanLine) > 1 {
			cmdName := cleanLine[0]
			option := cleanLine[1]
			if cmd, ok := supCom[cmdName]; ok {
				cmd.callback(mycaller, option)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
