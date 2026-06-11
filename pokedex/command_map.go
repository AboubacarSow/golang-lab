package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func callbackMap(cf *config) error {

	resp, err := pokeapi.ListLocationAreas(cf.pokeapiClient, cf.nextLocationAreasUrl)
	if err != nil {
		return err
	}
	cf.nextLocationAreasUrl = resp.LocationAreas.Next
	cf.previousLocationsUrl = resp.LocationAreas.Previous
	for _, area := range resp.LocationAreas.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
