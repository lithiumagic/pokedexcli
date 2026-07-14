package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *configStruct, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide one argument")
	}

	pokemonName := args[0]
	// check if pokemon in pokedex
	if pokemon, ok := config.Pokedex[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
