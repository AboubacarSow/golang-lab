package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()

		input := scanner.Text()

		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := cleaned[1:]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]

		if !ok == true {
			(availableCommands["invalid"]).callback(cf)
			continue
		}
		if err := command.callback(cf, args...); err != nil {
			fmt.Printf("%v", err)
		}
	}
}

type cliCommmand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func getCommands() map[string]cliCommmand {
	maps := map[string]cliCommmand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Display some location areas",
			callback:    callbackMap,
		},
		"mapback": {
			name:        "mapBack",
			description: "Display Previous location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location_area_name}",
			description: "Display Pokenam name for given location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokeman_name}",
			description: "Attempt to catch pokeman!",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all caught Pokémon during catch attempts",
			callback:    callbackPokedex,
		},
		"invalid": {
			name:        "invalid command",
			description: "Displayed when user type unavailabe command",
			callback:    callbackInvalid,
		},
		"exit": {
			name:        "exit",
			description: "Turns off Pokedex",
			callback:    callbackExit,
		},
	}

	return maps
}

func cleanInput(input string) []string {
	lower_input := strings.ToLower(input)

	words := strings.Fields(lower_input)

	return words
}
