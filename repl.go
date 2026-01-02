package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/migueldor/pokedexcli/internal/pokeapi"
	"github.com/migueldor/pokedexcli/internal/pokecache"
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
	Previous      string
	Next          string
}

func NewClient() pokeapi.Client {
	cache := pokecache.NewCache(5 * time.Second)

	return pokeapi.Client{
		Cache:      cache,
		HttpClient: http.Client{},
	}
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
	client := NewClient()
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
	client := NewClient()
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
	client := NewClient()

	location, err := client.GetLocation(opt)
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
