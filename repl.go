package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EveBisk/pokedex/internal/apiHelpers"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    apiHelpers.CommandMap,
		},
	}
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	registry := getCommands()

	for _, val := range registry {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandMap() error {

	return nil
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	registry := getCommands()

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		entry, ok := registry[commandName]

		if !ok {
			fmt.Printf("Command not found\n")
		} else {
			err := entry.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
