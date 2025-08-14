package main

import (
	"fmt"
	"sync"
)

var s string
var wg sync.WaitGroup

func main() {
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
