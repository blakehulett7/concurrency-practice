package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
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
