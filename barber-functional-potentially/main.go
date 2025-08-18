package main

import (
	"fmt"
	"math/rand"
	"time"
)

const customer_arrival_rate = 500 * time.Millisecond
const shop_day_length = 10 * time.Second
const waiting_room_capacity = 10

func main() {
	ColorPrint(Cyan, "Dominus Iesus Christus")
	ColorPrint(Cyan, "----------------------")
	fmt.Println()

	shop := BarberShop{
		Barbers:         []string{"Dave", "Noah", "Chris", "John"},
		BarberIsDone:    make(chan bool),
		CustomerChannel: make(chan string, waiting_room_capacity),
		IsClosing:       make(chan bool),
		IsClosed:        make(chan bool),
	}

	go StartBarberSystem(shop)
	go CustomerSystem(shop)
	go OpenShop(shop)

	<-shop.IsClosed

	fmt.Println()
	ColorPrint(Cyan, "----------------------")
	ColorPrint(Cyan, "Dominus Iesus Christus")
}

func BarberSystem(name string, shop BarberShop) {
	is_asleep := false

	for {
		if len(shop.CustomerChannel) == 0 {
			is_asleep = true
			ColorPrint(Yellow, fmt.Sprintf("No customers, barber %s takes a nap", name))
		}

		customer, shop_is_open := <-shop.CustomerChannel
		if !shop_is_open {
			ColorPrint(Green, fmt.Sprintf("No more customers, barber %s is going home", name))
			shop.BarberIsDone <- true
			return
		}

		if is_asleep {
			is_asleep = false
			ColorPrint(Yellow, fmt.Sprintf("%s wakes up barber %s", customer, name))
		}

		dice_roll := (rand.Intn(5) + 1) * 100
		haircut_duration := time.Millisecond * time.Duration(dice_roll)

		ColorPrint(Green, fmt.Sprintf("Barber %s has begun %s's haircut", name, customer))
		time.Sleep(haircut_duration)
		ColorPrint(Green, fmt.Sprintf("Barber %s has finished %s's haircut", name, customer))
	}
}

func CloseShop(shop BarberShop) {
	close(shop.CustomerChannel)

	barber_counter := len(shop.Barbers)
	for {
		if barber_counter == 0 {
			return
		}

		<-shop.BarberIsDone
		barber_counter--
	}
}

func CustomerSystem(shop BarberShop) {
	i := 1
	for {
		dice_roll := (rand.Intn(3) + 4) * 10
		arrival_time := time.Millisecond * time.Duration(dice_roll)

		select {
		case <-shop.IsClosing:
			return
		case <-time.After(arrival_time):
			customer_name := fmt.Sprintf("Customer %d", i)
			i++
			ColorPrint(Blue, fmt.Sprintf("%s has arrived", customer_name))
			select {
			case shop.CustomerChannel <- customer_name:
				ColorPrint(Blue, fmt.Sprintf("%s takes a seat", customer_name))
			default:
				ColorPrint(Red, fmt.Sprintf("No seats available, %s has left", customer_name))
			}
		}
	}
}

func OpenShop(shop BarberShop) {
	<-time.After(shop_day_length)
	shop.IsClosing <- true
	CloseShop(shop)
	shop.IsClosed <- true
}

func StartBarberSystem(shop BarberShop) {
	for _, barber := range shop.Barbers {
		go BarberSystem(barber, shop)
	}
}
