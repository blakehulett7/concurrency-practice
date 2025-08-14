package main

import (
	"sync"
	"testing"
)

func Test_PrintThis(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go PrintThis("Dominus Iesus Christus", &wg)
	wg.Wait()
}
