package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonBaseExperience(name string) (int, Pokemon, error) {
	url := fmt.Sprintf("%s%s", pokemonURL, name)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return 0, Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return 0, Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	pokemon_response := pokemonAPIResponse{}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return 0, Pokemon{}, err
	}

	err = json.Unmarshal(body, &pokemon_response)

	if err != nil {
		return 0, Pokemon{}, err
	}

	return pokemon_response.BaseExperience, pokemon_response.Pokemon, nil
}
