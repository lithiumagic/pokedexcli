package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(config *configStruct, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide one argument")
	}

	pokemonName := args[0]
	requestURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	var body []byte
	cachedData, ok := config.Cached.Get(requestURL)
	if ok {
		body = cachedData
	} else {
		res, err := http.Get(requestURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		config.Cached.Add(requestURL, body)
	}

	var data Pokemon
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	roll := rand.Intn(data.BaseExperience + 1)
	threshold := 50
	if roll < threshold {
		// caught
		fmt.Printf("%s was caught!\n", pokemonName)
		config.Pokedex[pokemonName] = data
	} else {
		// escaped
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

type Pokemon struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
