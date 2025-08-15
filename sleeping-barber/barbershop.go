package main

import (
	"fmt"
	"time"
)

type BarberShop struct {
	ShopCapacity       int
	CutDuration        time.Duration
	Barbers            int
	BarbersDoneChannel chan bool
	ClientsChannel     chan string
	IsOpen             bool
}

func (shop *BarberShop) AddBarber(name string) {
	shop.Barbers++

	go func() {
		is_sleeping := false
		ColorPrint(Yellow, fmt.Sprintf("%s goes to the waiting room to check for clients", name))

		for {
			if len(shop.ClientsChannel) == 0 {
				is_sleeping = true
				ColorPrint(Yellow, fmt.Sprintf("No customers... barber %s takes a nap", name))
			}

			client, shop_open := <-shop.ClientsChannel
			if !shop_open {
				shop.Barbers--
				ColorPrint(Red, fmt.Sprintf("Barber %s has gone home", name))
				return
			}

			if is_sleeping {
				is_sleeping = false
				ColorPrint(Yellow, fmt.Sprintf("%s wakes up barber %s", client, name))
			}

			ColorPrint(Yellow, fmt.Sprintf("Barber %s begins %s's hair cut", name, client))
			time.Sleep(shop.CutDuration)
			ColorPrint(Green, fmt.Sprintf("Barber %s has finished %s's hair cut", name, client))
		}
	}()
}
