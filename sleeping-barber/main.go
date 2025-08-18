package main

import (
	"fmt"
	"math/rand"
	"time"
)

const waiting_room_size = 10
const arrival_rate = 100
const cut_duration = 1000 * time.Millisecond
const time_open = 10 * time.Second

func main() {
	ColorPrint(Cyan, "Dominus Iesus Christus")
	ColorPrint(Cyan, "----------------------")
	fmt.Println()

	client_channel := make(chan string, waiting_room_size)
	done_channel := make(chan bool)

	shop := BarberShop{
		ShopCapacity:       waiting_room_size,
		CutDuration:        cut_duration,
		Barbers:            0,
		ClientsChannel:     client_channel,
		BarbersDoneChannel: done_channel,
		IsOpen:             true,
	}

	ColorPrint(Green, "Shop is open...")
	shop.AddBarber("Dave")
	shop.AddBarber("Noah")
	shop.AddBarber("Shana")
	shop.AddBarber("Donny")
	shop.AddBarber("Milton")

	shop_is_closing_channel := make(chan bool)
	shop_is_closed := make(chan bool)

	go func() {
		<-time.After(time_open)
		shop_is_closing_channel <- true
		shop.CloseShop()
		shop_is_closed <- true
	}()

	i := 1
	go func() {
		for {
			dice_roll := time.Duration(rand.Int() % (2 * arrival_rate))
			select {
			case <-shop_is_closing_channel:
				return
			case <-time.After(time.Millisecond * dice_roll):
				shop.AddClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	<-shop_is_closed
}

type Barber struct {
	IsAsleep bool
}

func (b *Barber) Sleep() {
	b.IsAsleep = true
}

func CustomerArrivesSystem(barber_is_asleep chan bool) {
	for {
		is_asleep := <-barber_is_asleep
		if is_asleep {

		}
	}
}
