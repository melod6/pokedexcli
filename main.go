package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hello, World!")
}

cleanInput(text string) []string {
	stringParts := text.split(text, " ")
	return stringParts
}
