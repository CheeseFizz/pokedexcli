package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CheeseFizz/pokedexcli/internal/apitools"
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
	callback    func(*config) error
	config      *config
}

var cliRegistry = make(map[string]cliCommand)

func commandMap(config *config) error {
	// get url from api registry if no config cache
	Url := config.next
	if len(Url) == 0 {
		url, err := apitools.GetPokeApiUrlPath("LocationAreas")
		if err != nil {
			return err
		}
		Url = url
	}

	// get list of location areas
	res, err := apitools.GetPokeApiResourceList(Url)
	if err != nil {
		return err
	}
	config.SetNext(res.Next)
	config.SetPrevious(res.Previous)

	fmt.Printf("in map config.next: %v\n", config.next)
	fmt.Printf("in map config.previous: %v\n", config.previous)

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if len(config.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.previous

	// get list of location areas
	res, err := apitools.GetPokeApiResourceList(url)
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

func commandHelp(config *config) error {
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

func commandExit(config *config) error {
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command, ok := cliRegistry[input[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(command.config)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}

}
