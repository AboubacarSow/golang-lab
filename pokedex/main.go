package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient *pokeapi.Client
	nextLocationAreasUrl *string
	previousLocationsUrl *string
}
func buildConfif(interval time.Duration) *config{
	return &config{
		pokeapiClient: pokeapi.NewClient(interval),
	}
}
func main() {
	var duration = 45 * time.Second
	cf:=buildConfif(duration)
	startRepl(cf)
}
