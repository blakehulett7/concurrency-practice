package main

type BarberShop struct {
	Barbers         []string
	BarberIsDone    chan bool
	CustomerChannel chan string
	IsClosing       chan bool
	IsClosed        chan bool
}
