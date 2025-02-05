package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locations_response := LocationAreasResponse{}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	json.Unmarshal(body, &locations_response)

	if res.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locations_response, nil

}
