package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/talgat-ruby/lessons-go/greet"
)

func init() {
	fmt.Println("Initialization...")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your name below: ")

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			greetings := greet.Greet(text)
			fmt.Printf(greetings)
		} else {
			fmt.Println("Please enter your name.")
		}

	}
}
