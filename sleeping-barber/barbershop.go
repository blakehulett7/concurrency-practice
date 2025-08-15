package main

import "time"

type BarberShop struct {
	ShopCapacity       int
	CutDuration        time.Duration
	Barbers            int
	BarbersDoneChannel chan bool
	ClientsChannel     chan string
	IsOpen             bool
}
