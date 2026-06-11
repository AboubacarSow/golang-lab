package main

import "pokedex/internal/pokeapi"

type config struct {
	pokeapiClient *pokeapi.Client
	nextLocationAreasUrl *string
	previousLocationsUrl *string
}
func buildConfif() *config{
	return &config{
		pokeapiClient: pokeapi.NewClient(),
	}
}
func main() {
	cf:=buildConfif()
	startRepl(cf)
}
