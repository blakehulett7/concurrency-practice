package main

import (
	"fmt"
	"strings"
	"time"
)

// Buffered Channels
func main() {
	ch := make(chan int, 10)

	go ListenToChan(ch)

	for i := range 100 {
		fmt.Println("sending", i, "to channel")
		ch <- i
		fmt.Println("sent", i)
	}

	close(ch)
}

func ListenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got data:", i)

		time.Sleep(time.Second)
	}
}

func channel_select() {
	fmt.Println("Jesus is Lord!")
	fmt.Println("--------------")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	for {
		select {
		case s1 := <-ch1:
			fmt.Println("Case 1", s1)
		case s2 := <-ch1:
			fmt.Println("Case 2", s2)
		case s3 := <-ch2:
			fmt.Println("Case 3", s3)
		case s4 := <-ch2:
			fmt.Println("Case 4", s4)
		}
	}
}

func server1(ch chan string) {
	for {
		time.Sleep(time.Second * 6)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func shout_prompt() {
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

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping                                   // When you get something from channel ping, put it in variable s
		pong <- fmt.Sprintf("%s", strings.ToUpper(s)) // Send something to channel pong
	}
}
