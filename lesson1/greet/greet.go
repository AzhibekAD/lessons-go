package greet

import "fmt"

func Greet(character string, table string) string {
	return fmt.Sprintf("Hello, %s! Your table is %s!", character, table)
}
