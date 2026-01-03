package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/migueldor/pokedexcli/internal/pokeapi"
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
	callback    func(caller *config, opt string) error
}

var supCom map[string]cliCommand

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.PokemonResponse
	Previous      string
	Next          string
}

func commandExit(caller *config, opt string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(caller *config, opt string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Printf("help: %s\n", supCom["help"].description)
	fmt.Printf("exit: %s\n", supCom["exit"].description)
	return nil
}

func commandMap(caller *config, opt string) error {
	client := caller.pokeapiClient
	if caller.Next == "" {
		caller.Next = "https://pokeapi.co/api/v2/location-area"
	}
	locations, err := client.LocationArea(caller.Next)
	if err != nil {
		return err
	}
	if locations.Previous != nil {
		caller.Previous = *locations.Previous
	}
	if locations.Next != nil {
		caller.Next = *locations.Next
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(caller *config, opt string) error {
	client := caller.pokeapiClient
	if caller.Previous == "" {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	locations, err := client.LocationArea(caller.Previous)
	if err != nil {
		return err
	}
	if locations.Previous != nil {
		caller.Previous = *locations.Previous
	}
	if locations.Next != nil {
		caller.Next = *locations.Next
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(caller *config, opt string) error {
	client := caller.pokeapiClient

	location, err := client.Location(opt)
	if err != nil {
		fmt.Println("something went wrong")
		return err
	}

	fmt.Printf("Exploring %s...\n", opt)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(caller *config, opt string) error {
	client := caller.pokeapiClient

	pokemon, err := client.Pokemon(opt)
	if err != nil {
		fmt.Println("something went wrong")
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", opt)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	caller.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func commandInspect(caller *config, opt string) error {
	client := caller.pokeapiClient
	types := ""
	pokemon, err := client.Pokemon(opt)
	if err != nil {
		fmt.Println("something went wrong")
		return err
	}
	if _, ok := caller.caughtPokemon[pokemon.Name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("ID number: %d\n", pokemon.ID)
		for i := 0; i < len(pokemon.Types); i++ {
			types = types + " " + pokemon.Types[i].Type.Name
		}
		fmt.Printf("Types: %v\n", types)
		fmt.Printf("Height: %d cm\n", pokemon.Height)
	} else {
		fmt.Printf("%s not caught yet\n", pokemon.Name)
	}

	return nil
}

func commandPokedex(caller *config, opt string) error {
	fmt.Println("Your pokedex: ")
	for _, pokemon := range caller.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
