package main

import "fmt"

const NumberOfPizzas = 10

var PizzasMade int
var PizzasFailed int
var TotalPizzas int

type PizzaOrder struct {
	Number       int
	Message      string
	IsSuccessful bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func main() {
	seed := ""

	fmt.Println("Dominus Iesus Christus")
}
