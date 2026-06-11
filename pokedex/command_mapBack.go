package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func callbackMapBack(cf *config) error {
	if cf.previousLocationsUrl == nil{
		fmt.Println("No previous! You're in the first page")
		return nil
	}
	resp, err := pokeapi.ListLocationAreas(cf.pokeapiClient, cf.previousLocationsUrl)
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
