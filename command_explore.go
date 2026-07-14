package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(config *configStruct, args ...string) error {
	if len(args) != 1 {
		return errors.New("more than one argument")
	}

	areaName := args[0]
	requestURL := "https://pokeapi.co/api/v2/location-area/" + areaName

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

	var data exploreLocationResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

type exploreLocationResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
