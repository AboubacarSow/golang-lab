package main

import (
	"fmt"
	
)

func callbackPokedex(cf *config, args ...string) error {
	if len(cf.caugthPokemon) == 0 {
		return fmt.Errorf("No Pokemon caught yet!\n")
	}
	fmt.Println("Pokedex:")
	index := 0
	for _, pokemon := range cf.caugthPokemon {
		index++
		fmt.Printf("- %d. %s\n", index, pokemon.Name)
	}
	return nil
}


