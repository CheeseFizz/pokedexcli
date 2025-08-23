package main

import (
	"bufio"
	"fmt"
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
