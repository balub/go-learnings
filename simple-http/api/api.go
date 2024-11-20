package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseURL = "https://pokeapi.co/api/v2/pokemon"

func GetPokemonDetails(name string) (*Pokemon, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("Invalid name")
	}

	pokemonAPIUrl := fmt.Sprintf("%s/%s", baseURL, name)

	resp, err := http.Get(pokemonAPIUrl)
	if err != nil {
		return nil, fmt.Errorf("Error making request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Something Failed")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing response body")
	}
	var pokemon Pokemon

	if err := json.Unmarshal(bodyBytes, &pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil

}
