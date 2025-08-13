package main

import (
	"fmt"
)

func main() {
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
		go PrintThis(fmt.Sprintf("%d. %s", i, phrase))
	}

	PrintThis("In Nomine Patris...")
	PrintThis("Et Filii...")
	PrintThis("Et Spiritus Sancti!")
}

func PrintThis(s string) {
	fmt.Println(s)
}
