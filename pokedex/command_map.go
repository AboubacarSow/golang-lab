package main

import (
	"fmt"
)

func callbackMap(cf *config) error {

	resp, err := cf.pokeapiClient.ListLocationAreas(cf.nextLocationAreasUrl)
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
