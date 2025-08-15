package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Dominus Iesus Christus")

	ping := make(chan string)
	pong := make(chan string)

	defer close(ping)
	defer close(pong)

	go shout(ping, pong)

	fmt.Println("Type something...")

	for {
		fmt.Print("-> ")
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			return
		}

		ping <- input

		response := <-pong
		fmt.Printf("Response: %s\n", response)
	}
}

func shout(ping, pong chan string) {
	for {
		s := <-ping                                   // When you get something from channel ping, put it in variable s
		pong <- fmt.Sprintf("%s", strings.ToUpper(s)) // Send something to channel pong
	}
}
