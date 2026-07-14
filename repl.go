package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lithiumagic/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin) // when it runs it blocks and waits for input until user presses Enter
	config := configStruct{
		Cached:  pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := words[1:]

		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.callback(&config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type configStruct struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
	Cached   *pokecache.Cache
	Pokedex  map[string]Pokemon
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *configStruct, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "displays a list of all the Pokémon located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch chosen Pokémon located there",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects an already caught Pokémon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "prints a list of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}
