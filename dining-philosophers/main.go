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

	fmt.Printf("%s is seated at the table.\n", philosopher.Name)
	seated.Done()

	seated.Wait()

	for i := Hunger; i > 0; i-- {
		if philosopher.LeftFork > philosopher.RightFork {

			forks[philosopher.RightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.Name)

			forks[philosopher.LeftFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.Name)

		} else {

			forks[philosopher.LeftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.Name)

			forks[philosopher.RightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.Name)

		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.Name)
		time.Sleep(EatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.Name)
		time.Sleep(ThinkTime)

		forks[philosopher.LeftFork].Unlock()
		forks[philosopher.RightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.Name)
	}

	fmt.Printf("%s is satisfied.\n", philosopher.Name)
	fmt.Printf("%s left the table.\n", philosopher.Name)
}
