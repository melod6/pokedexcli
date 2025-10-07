package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	configFile = config{
		next:     "",
		previous: "",
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]
		_, supportedCommand := supportedCommands[command]
		if supportedCommand {
			supportedCommands[command].callback(&configFile)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

var supportedCommands map[string]cliCommand

var configFile config

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays 20 map locations. Subsequent uses display the next 20 locations.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 map locations.",
			callback: commandMapBack,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	next     string
	previous string
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, value := range supportedCommands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

type locationAreaResonse struct {
	Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results[]struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != "" {
		url = c.next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	locationAreaRes := locationAreaResonse{}
	err = json.Unmarshal(body, &locationAreaRes)
	c.next = fmt.Sprintf("%s%s", locationAreaRes.Next, locationAreaRes.Count)
	c.previous = locationAreaRes.Previous
	for _, location := range locationAreaRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.previous != "" {
		url = c.previous
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	locationAreaRes := locationAreaResonse{}
	err = json.Unmarshal(body, &locationAreaRes)
	c.next = fmt.Sprintf("%s%s", locationAreaRes.Next, locationAreaRes.Count)
	c.previous = locationAreaRes.Previous
	for _, location := range locationAreaRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
