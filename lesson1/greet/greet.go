package greet

import "fmt"

func Greet(character string) string {
	return fmt.Sprintf("Hello, %s! \n", character)
}
