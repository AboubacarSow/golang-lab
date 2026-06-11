package main

import (
	"errors"
	"fmt"
)

func callbackMapBack(cf *config) error {
	if cf.previousLocationsUrl == nil{
		return errors.New("No previous! You're in the first page")
	}
	resp, err := cf.pokeapiClient.ListLocationAreas(cf.previousLocationsUrl)
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
