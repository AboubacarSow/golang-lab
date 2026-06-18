package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocationAreas(pagUrl *string) (Response, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pagUrl != nil {
		fullUrl = *pagUrl
	}
	data, ok := client.Cache.Get(fullUrl)
	if !ok {
		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			return Response{}, nil
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return Response{}, nil
		}
		if resp.StatusCode > 399 {
			return Response{}, fmt.Errorf("Bad request. Status code %d", resp.StatusCode)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, nil
		}
		client.Cache.Add(data, fullUrl)
	}

	response := Response{}
	err := json.Unmarshal(data, &response.LocationAreas)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to unmarshall data:%s", err)
	}

	return response, nil

}

func (client *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	data, ok := client.Cache.Get(fullUrl)
	if !ok {
		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			return LocationArea{}, nil
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, nil
		}
		if resp.StatusCode > 399 {
			return LocationArea{}, fmt.Errorf("Bad request. Status code %d", resp.StatusCode)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, nil
		}
		client.Cache.Add(data, fullUrl)
	}

	locationArea := LocationArea{}
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Failed to unmarshall data:%s", err)
	}

	return locationArea, nil

}

func (client *Client) GetPokeman(pokemanName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemanName
	fullUrl := baseUrl + endpoint

	data, ok := client.Cache.Get(fullUrl)
	if !ok {
		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			return Pokemon{}, nil
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, nil
		}
		if resp.StatusCode > 399 {
			return Pokemon{}, fmt.Errorf("Bad request. Status code %d\n", resp.StatusCode)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, nil
		}
		client.Cache.Add(data, fullUrl)
	}

	pokeman := Pokemon{}
	err := json.Unmarshal(data, &pokeman)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to unmarshall data:%s", err)
	}

	return pokeman, nil

}
