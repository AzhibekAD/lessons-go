package main

import (
	"bufio"
	"fmt"
	"github.com/talgat-ruby/lessons-go/greet"
	"os"
)

func init() {
	fmt.Println("Initialization...")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tableScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your name below: ")
	for scanner.Scan() {
		var table string
		var text string

		text = scanner.Text()

		if text != "" {
			fmt.Println("Please enter your table number below: ")
			for tableScanner.Scan() {
				table = tableScanner.Text()
				if table == "" {
					fmt.Println("Please enter your name again.")
					break
				} else {
					greetings := greet.Greet(text, table)
					fmt.Printf(greetings)
				}
			}
		} else {
			fmt.Println("Please enter your name.")
		}

	}
}
