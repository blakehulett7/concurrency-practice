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

	barber_counter := 0
	barber_is_done := make(chan bool)
	shop_is_closing := make(chan bool)
	shop_is_closed := make(chan bool)
	customers_channel := make(chan string, 10)

	barber_counter++
	go BarberArrives("Dave", barber_is_done, customers_channel)
	barber_counter++
	go BarberArrives("Noah", barber_is_done, customers_channel)
	barber_counter++
	go BarberArrives("Chris", barber_is_done, customers_channel)
	barber_counter++
	go BarberArrives("John", barber_is_done, customers_channel)
	go CustomerSystem(customers_channel, shop_is_closing)
	go func() {
		<-time.After(shop_day_length)
		shop_is_closing <- true
		CloseShop(barber_counter, barber_is_done, customers_channel)
		shop_is_closed <- true
	}()

	<-shop_is_closed

	fmt.Println()
	ColorPrint(Cyan, "----------------------")
	ColorPrint(Cyan, "Dominus Iesus Christus")
}

func BarberArrives(name string, barber_is_done chan bool, customer_channel chan string) {
	is_asleep := false

	for {
		if len(customer_channel) == 0 {
			is_asleep = true
			ColorPrint(Yellow, fmt.Sprintf("No customers, barber %s takes a nap", name))
		}

		customer, shop_is_open := <-customer_channel
		if !shop_is_open {
			ColorPrint(Green, fmt.Sprintf("No more customers, barber %s is going home", name))
			barber_is_done <- true
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

func CloseShop(barber_counter int, barber_is_done chan bool, customer_channel chan string) {
	close(customer_channel)
	for {
		if barber_counter == 0 {
			return
		}

		<-barber_is_done
		barber_counter--
	}
}

func CustomerSystem(customer_channel chan string, shop_is_closing chan bool) {
	i := 1
	for {
		dice_roll := (rand.Intn(3) + 4) * 10
		arrival_time := time.Millisecond * time.Duration(dice_roll)

		select {
		case <-shop_is_closing:
			return
		case <-time.After(arrival_time):
			customer_name := fmt.Sprintf("Customer %d", i)
			i++
			ColorPrint(Blue, fmt.Sprintf("%s has arrived", customer_name))
			select {
			case customer_channel <- customer_name:
				ColorPrint(Blue, fmt.Sprintf("%s takes a seat", customer_name))
			default:
				ColorPrint(Red, fmt.Sprintf("No seats available, %s has left", customer_name))
			}
		}
	}
}
