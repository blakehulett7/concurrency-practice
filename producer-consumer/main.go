package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func MakePizza(pizza_number int) *PizzaOrder {
	pizza_number++

	if pizza_number > NumberOfPizzas {
		return &PizzaOrder{Number: pizza_number}
	}

	delay := rand.Intn(5) + 1
	fmt.Printf("Received order no. %d\n", pizza_number)
	fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizza_number, delay)

	time.Sleep(time.Duration(delay) * time.Second)

	dice_roll := rand.Intn(12) + 1
	success := dice_roll > 4
	if success {
		PizzasMade++
		TotalPizzas++
		return &PizzaOrder{
			Number:       pizza_number,
			Message:      fmt.Sprintf("pizza no %d is ready!", pizza_number),
			IsSuccessful: success,
		}
	}

	PizzasFailed++
	TotalPizzas++

	if dice_roll <= 2 {
		return &PizzaOrder{
			Number:       pizza_number,
			Message:      fmt.Sprintf("*** Ran out of ingredients for pizza no %d!", pizza_number),
			IsSuccessful: success,
		}
	}

	return &PizzaOrder{
		Number:       pizza_number,
		Message:      fmt.Sprintf("*** Cook quit while making pizza no %d!", pizza_number),
		IsSuccessful: success,
	}
}

func Pizzeria(pizza_maker *Producer) {
	i := 0

	for {
		current_pizza := MakePizza(i)
	}
}
