package main

import (
	"fmt"
	"strings"
)

func callbackInspect(cf *config, args ...string) error {
	if len(cf.caugthPokemon) == 0 {
		return fmt.Errorf("No Pokemon caught yet!\n")
	}
	pokemonName := strings.Join(args, " ")

	pokemon, ok := cf.caugthPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("No Pokemon with Name:%s was caught\n", pokemonName)
	}
	fmt.Printf("%s Information Details:\n", pokemon.Name)
	fmt.Printf("Height:%d\n", pokemon.Height)
	fmt.Printf("Weight:%d\n", pokemon.Weight)
	fmt.Printf("Base Experience:%d\n", pokemon.BaseExperience)
	fmt.Printf("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %v\n", stat)
	}
	fmt.Printf("Types:\n")
	for _, val := range pokemon.Types {
		fmt.Printf("- %v\n", val.Type.Name)
	}
	return nil
}
