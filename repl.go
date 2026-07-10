package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splitStringSlice := []string{}

	text = strings.ToLower(text)
	splitStringSlice = strings.Fields(text)
	return splitStringSlice
}
