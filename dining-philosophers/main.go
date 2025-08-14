package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name      string
	LeftFork  int
	RightFork int
}

var Philosophers = []Philosopher{
	{Name: "Jesus Christ", LeftFork: 4, RightFork: 0},
	{Name: "Thomas Acquinas", LeftFork: 0, RightFork: 1},
	{Name: "CS Lewis", LeftFork: 1, RightFork: 2},
	{Name: "JRR Tolkien", LeftFork: 2, RightFork: 3},
	{Name: "Augustine", LeftFork: 3, RightFork: 4},
}

var Hunger = 3
var EatTime = time.Second
var ThinkTime = 3 * time.Second
var SleepTime = time.Second

func main() {
	fmt.Println("Dominus Iesus Christus")
	fmt.Println("----------------------")
	fmt.Println()
	fmt.Println("The table is empty")

	eating := &sync.WaitGroup{}
	seated := &sync.WaitGroup{}

	eating.Add(len(Philosophers))
	seated.Add(len(Philosophers))

	forks := map[int]*sync.Mutex{}
	for i := 0; i < len(Philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for _, philosopher := range Philosophers {
		go Dine(philosopher, eating, seated, forks)
	}
	eating.Wait()

	fmt.Println("The table is empty")

}

func Dine(philosopher Philosopher, eating, seated *sync.WaitGroup, forks map[int]*sync.Mutex) {
	defer eating.Done()
}
