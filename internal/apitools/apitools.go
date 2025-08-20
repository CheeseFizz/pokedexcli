package apitools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const pokeApi = "https://pokeapi.co/api/v2/"

var pokeApiReference = map[string]string{
	"LocationAreas": fmt.Sprintf("%slocation-area", pokeApi),
}

type NamedApiResource struct {
	Name string
	Url  string
}

type NamedApiResourceList struct {
	Count    int
	Next     string
	Previous string
	Results  []NamedApiResource
}

func GetPokeApiResourceList(url string) (NamedApiResourceList, error) {
	var result NamedApiResourceList
	var zero NamedApiResourceList

	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return zero, err
	}

	return result, nil
}

func GetPokeApiUrlPath(name string) (string, error) {
	result, ok := pokeApiReference[name]
	if !ok {
		return "", fmt.Errorf("resource not found: %s", name)
	}
	return result, nil
}
