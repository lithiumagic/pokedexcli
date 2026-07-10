package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // when it runs it blocks and waits for input until user presses Enter
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.ToLower(userInput)
		cleanUserInput := strings.Fields(userInput)
		if len(cleanUserInput) == 0 {
			continue
		}
		firstWord := cleanUserInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}

}
