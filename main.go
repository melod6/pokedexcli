package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v\n", input[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
