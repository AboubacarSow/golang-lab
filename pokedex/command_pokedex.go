package main

import "fmt"

func callbackPokedex(cf *config, args ...string) error {

	size := len(cf.caugthPokemon)
	if size == 0 {
		return fmt.Errorf("Not Pokemon caught yet!\n")
	}
	fmt.Println("List of pokemon caught:")
	index := 0
	for _, pokemon := range cf.caugthPokemon {
		index++
		fmt.Printf("%d. %s\n", index, pokemon.Name)
	}
	return nil
}
