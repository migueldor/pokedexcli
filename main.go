package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		if line == "q" {
			break
		}
		cleanLine := cleanInput(line)
		if len(cleanLine) > 0 {
			fmt.Printf("Your command was: %s\n", cleanLine[0])
		}
	}
}
