package main

import (
	"errors"
	"fmt"
	"strings"
)

func callbackExplore(cf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No argument for location name provided")
	}
	name := strings.Join(args, " ")
	resp, err := cf.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}
	fmt.Printf("Pokeman in :%s\n", resp.Location.Name)
	for _, pokemanencounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemanencounter.Pokemon.Name)
	}
	return nil
}
