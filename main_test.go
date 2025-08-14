package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_PrintThis(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go PrintThis("Dominus Iesus Christus", &wg)
	wg.Wait()
}

func Test_UpdateThis(t *testing.T) {
	var m sync.Mutex

	wg.Add(2)
	go UpdateThis("1", &m)
	go UpdateThis("2", &m)
	wg.Wait()

	wg.Add(2)
	go func() {
		s = "3"
		wg.Done()
	}()
	go func() {
		s = "4"
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(s)
}
