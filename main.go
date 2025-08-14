package main

import (
	"fmt"
	"sync"
)

var s string
var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	balance := 0
	var m sync.Mutex

	fmt.Printf("Initial balance: %d\n", balance)
	fmt.Println()

	incomes := []Income{
		{Source: "main", Amount: 500},
		{Source: "gifts", Amount: 10},
		{Source: "side", Amount: 50},
		{Source: "flow", Amount: 100},
	}

	for i, income := range incomes {
		wg.Add(1)
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				m.Lock()
				current := balance
				current += income.Amount
				balance = current
				m.Unlock()

				fmt.Printf("On week %d, we got %d from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final balance: %d\n", balance)
	fmt.Println()
}

func OldMain() {
	s = "Dominus Iesus Christus"

	var m sync.Mutex

	wg.Add(3)
	go UpdateThis("Gloria Patris", &m)
	go UpdateThis("Et Filii", &m)
	go UpdateThis("Et Spiritus Sancti", &m)
	wg.Wait()

	fmt.Println(s)

	hail_mary := []string{
		"Ave Maria, gratia plena",
		"Dominus tecum",
		"benedicta tu in mulieribus",
		"et benedictus fructus ventris tui, Iesus.",
		"Sacta Maria mater dei,",
		"ora pro nobis peccatoribus",
		"nunc et in hora a mortis nostrae",
		"Amen",
	}

	for i, phrase := range hail_mary {
		wg.Add(1)
		go PrintThis(fmt.Sprintf("%d. %s", i, phrase), &wg)
	}
	wg.Wait()

	wg.Add(3)
	PrintThis("In Nomine Patris...", &wg)
	PrintThis("Et Filii...", &wg)
	PrintThis("Et Spiritus Sancti!", &wg)
}

func PrintThis(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func UpdateThis(new_s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	defer m.Unlock()
	s = new_s
}
