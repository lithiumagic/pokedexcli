package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *configStruct) error {
	requestURL := ""
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		requestURL = config.Previous
	}

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

	var data configStruct
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}
