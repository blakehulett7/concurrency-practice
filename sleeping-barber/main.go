package main

import "time"

const waiting_room_size = 10
const arrival_rate = 100
const cut_duration = 1000 * time.Millisecond
const time_open = 10 * time.Second

func main() {
	ColorPrint(Cyan, "Dominus Iesus Christus")
	ColorPrint(Cyan, "----------------------")

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

	time.Sleep(5 * time.Second)
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
