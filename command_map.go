package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *configStruct) error {
	requestURL := ""
	if config.Next == "" {
		requestURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		requestURL = config.Next
	}

	res, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data configStruct
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	return nil
}
