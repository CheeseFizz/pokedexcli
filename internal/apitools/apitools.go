package apitools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CheeseFizz/pokedexcli/internal/pokecache"
)

const pokeApi = "https://pokeapi.co/api/v2/"

var pokeApiReference = map[string]string{
	"Berries":                  fmt.Sprintf("%sberry", pokeApi),
	"BerryFirmness":            fmt.Sprintf("%sberry-firmness", pokeApi),
	"BerryFlavors":             fmt.Sprintf("%sberry-flavor", pokeApi),
	"ContestTypes":             fmt.Sprintf("%scontest-type", pokeApi),
	"ContestEffects":           fmt.Sprintf("%scontest-effect", pokeApi),
	"SuperContestEffects":      fmt.Sprintf("%ssuper-contest-effect", pokeApi),
	"EncounterMethods":         fmt.Sprintf("%sencounter-method", pokeApi),
	"EncounterConditions":      fmt.Sprintf("%sencounter-condition", pokeApi),
	"EncounterConditionValues": fmt.Sprintf("%sencounter-condition-value", pokeApi),
	"EvolutionChains":          fmt.Sprintf("%sevolution-chain", pokeApi),
	"EvolutionTriggers":        fmt.Sprintf("%sevolution-trigger", pokeApi),
	"Generations":              fmt.Sprintf("%sgeneration", pokeApi),
	"Pokedexes":                fmt.Sprintf("%spokedex", pokeApi),
	"Versions":                 fmt.Sprintf("%sversion", pokeApi),
	"VersionGroups":            fmt.Sprintf("%sversion-group", pokeApi),
	"Items":                    fmt.Sprintf("%sitem", pokeApi),
	"ItemAttributes":           fmt.Sprintf("%sitem-attribute", pokeApi),
	"ItemCategories":           fmt.Sprintf("%sitem-category", pokeApi),
	"ItemFlingEffects":         fmt.Sprintf("%sitem-fling-effect", pokeApi),
	"ItemPockets":              fmt.Sprintf("%sitem-pocket", pokeApi),
	"Locations":                fmt.Sprintf("%slocation", pokeApi),
	"LocationAreas":            fmt.Sprintf("%slocation-area", pokeApi),
	"PalParkAreas":             fmt.Sprintf("%spal-park-area", pokeApi),
	"Regions":                  fmt.Sprintf("%sregion", pokeApi),
	"Machines":                 fmt.Sprintf("%smachine", pokeApi),
	"Moves":                    fmt.Sprintf("%smove", pokeApi),
	"MoveAilments":             fmt.Sprintf("%smove-ailment", pokeApi),
	"MoveBattleStyle":          fmt.Sprintf("%smove-battle-style", pokeApi),
	"MoveCategories":           fmt.Sprintf("%smove-category", pokeApi),
	"MoveDamageClasses":        fmt.Sprintf("%smove-damage-class", pokeApi),
	"MoveLearnMethods":         fmt.Sprintf("%smove-learn-method", pokeApi),
	"MoveTargets":              fmt.Sprintf("%smove-target", pokeApi),
	"Abilities":                fmt.Sprintf("%sability", pokeApi),
	"Characteristics":          fmt.Sprintf("%scharacteristic", pokeApi),
	"EggGroups":                fmt.Sprintf("%segg-group", pokeApi),
	"Genders":                  fmt.Sprintf("%sgender", pokeApi),
	"GrowthRates":              fmt.Sprintf("%sgrowth-rate", pokeApi),
	"Natures":                  fmt.Sprintf("%snature", pokeApi),
	"PokeathlonStats":          fmt.Sprintf("%spokeathlon-stat", pokeApi),
	"Pokemon":                  fmt.Sprintf("%spokemon", pokeApi),
	"PokemonLocationAreas":     fmt.Sprintf("%spokemon/REPLACE/encounters", pokeApi),
	"PokemonColors":            fmt.Sprintf("%spokemon-color", pokeApi),
	"PokemonForms":             fmt.Sprintf("%spokemon-form", pokeApi),
	"PokemonHabitats":          fmt.Sprintf("%spokemon-habitat", pokeApi),
	"PokemonShapes":            fmt.Sprintf("%spokemon-shape", pokeApi),
	"PokemonSpecies":           fmt.Sprintf("%spokemon-species", pokeApi),
	"Stats":                    fmt.Sprintf("%sstat", pokeApi),
	"Types":                    fmt.Sprintf("%stype", pokeApi),
	"Languages":                fmt.Sprintf("%slanguage", pokeApi),
}

func GetPokeApiResourceList(url string, c *pokecache.Cache) (NamedApiResourceList, error) {
	var result NamedApiResourceList
	var zero NamedApiResourceList

	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return result, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return zero, err
		}

		c.Add(url, body)
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
