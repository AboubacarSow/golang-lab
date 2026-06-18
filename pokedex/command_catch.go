package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func callbackCatch(cf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No argument for pokeman name provided")
	}
	pokemanName := strings.Join(args, " ")
	pokemon, err := cf.pokeapiClient.GetPokeman(pokemanName)
	if err != nil {
		return err
	}

	
	const treshHold = 60
	randNumber := rand.Intn(pokemon.BaseExperience)

	if randNumber > treshHold{
		return fmt.Errorf("Failed to catch pokeman %s\n",pokemanName)
	}

	if _, ok:= cf.caugthPokemon[pokemanName]; ok==false {
		cf.caugthPokemon[pokemanName]=pokemon
	}
	fmt.Printf("%s pokeman was catched\n",pokemanName)
	return nil
}
