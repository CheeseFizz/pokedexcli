package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/CheeseFizz/pokedexcli/internal/apitools"
	"github.com/CheeseFizz/pokedexcli/internal/pokecache"
)

type config struct {
	next     string `default:""`
	previous string `default:""`
}

func (c *config) SetNext(url string) {
	c.next = url
}
func (c *config) SetPrevious(url string) {
	c.previous = url
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
	config      *config
}

var cliRegistry = make(map[string]cliCommand)
var apiCache = pokecache.NewCache(time.Duration(10 * time.Second))
var userPokedex = make(map[string]apitools.Pokemon)

func commandMap(config *config, _ []string) error {
	// get url from api registry if no config cache
	var err error
	url := config.next
	if len(url) == 0 {
		url, err = apitools.GetPokeApiUrlPath("LocationAreas")
		if err != nil {
			return err
		}
	}

	// get list of location areas
	res, err := apitools.GetPokeApiResourceList(url, apiCache)
	if err != nil {
		return err
	}
	config.SetNext(res.Next)
	config.SetPrevious(res.Previous)

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func commandMapb(config *config, _ []string) error {
	if len(config.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.previous

	// get list of location areas
	res, err := apitools.GetPokeApiResourceList(url, apiCache)
	if err != nil {
		return err
	}
	config.SetNext(res.Next)
	config.SetPrevious(res.Previous)

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func commandExplore(_ *config, locationArea []string) error {
	location := apitools.LocationArea{}

	url, err := apitools.GetPokeApiUrlPath("LocationAreas")
	if err != nil {
		return err
	}
	url = fmt.Sprintf("%s/%s", url, locationArea[0])

	err = apitools.GetPokeApiResource(url, apiCache, &location)
	if err != nil {
		return err
	}

	for _, enc := range location.Pokemon_encounters {
		fmt.Println(enc.Pokemon.Name)
	}

	return nil
}

func commandCatch(_ *config, pokemonName []string) error {
	pokemon := apitools.Pokemon{}
	species := apitools.PokemonSpecies{}

	url, err := apitools.GetPokeApiUrlPath("Pokemon")
	if err != nil {
		return err
	}
	url = fmt.Sprintf("%s/%s", url, pokemonName[0])

	err = apitools.GetPokeApiResource(url, apiCache, &pokemon)
	if err != nil {
		return err
	}
	err = apitools.GetPokeApiResource(pokemon.Species.Url, apiCache, &species)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	try := rand.Intn(300)

	//time.Sleep(2 * time.Second)

	if try < species.Capture_rate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		userPokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(_ *config, pokemonName []string) error {
	pokemon, ok := userPokedex[pokemonName[0]]
	if !ok {
		fmt.Printf("%s not in your Pokedex!\n", pokemonName[0])
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, pstat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", pstat.Stat.Name, pstat.Base_stat)
	}
	fmt.Print("Type:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}

func commandHelp(_ *config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, entry := range cliRegistry {
		_, err := fmt.Printf("%s: %s\n", entry.name, entry.description)
		if err != nil {
			return err
		}
	}
	fmt.Println()
	return nil
}

func commandExit(_ *config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Something went wrong; the program didn't exit.")
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))

	return result
}

func main() {

	mapConfig := &config{"", ""}

	cliRegistry = map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List next 20 map locations",
			callback:    commandMap,
			config:      mapConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 map locations",
			callback:    commandMapb,
			config:      mapConfig,
		},
		"explore": {
			name:        "explore",
			description: "List Pokemon in the location. Use: explore <map-location>",
			callback:    commandExplore,
			config:      &config{},
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon! Use: catch <pokemon-name>",
			callback:    commandCatch,
			config:      &config{},
		},
		"inspect": {
			name:        "inspect",
			description: "See information about Pokemon you've caught. Use: inspect <pokemon-name>",
			callback:    commandInspect,
			config:      &config{},
		},
		"help": {
			name:        "help",
			description: "Print 'help' message",
			callback:    commandHelp,
			config:      &config{},
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
			config:      &config{},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	var args = []string{}
	for {
		args = []string{}
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		if len(input) > 1 {
			args = input[1:]
		}
		command, ok := cliRegistry[input[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(command.config, args)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}

}
