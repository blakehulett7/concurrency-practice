package main

const NumberOfPizzas = 10

var PizzasMade int
var PizzasFailed int
var TotalPizzas int

type PizzaOrder struct {
	Number       int
	Message      string
	IsSuccessful bool
}

func main() {
	ColorPrint("Dominus Iesus Christus", Cyan)
	ColorPrint("----------------------", Cyan)

	pizza_job := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go Pizzeria(pizza_job)
}

func Pizzeria(pizza_maker *Producer) {
	for {

	}
}
