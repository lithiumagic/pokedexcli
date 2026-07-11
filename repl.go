package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin) // when it runs it blocks and waits for input until user presses Enter
	config := configStruct{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.callback(&config)
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
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configStruct) error
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
	}
}
