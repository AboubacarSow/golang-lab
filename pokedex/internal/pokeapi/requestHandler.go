package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func(client *Client) ListLocationAreas( pagUrl *string) (Response, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pagUrl != nil {
		fullUrl = *pagUrl
	}
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, nil
	}
	response := Response{}
	err = json.Unmarshal(data, &response.LocationAreas)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to unmarshall data:%s", err)
	}

	return response, nil

}
