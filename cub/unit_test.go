package main

import (
	"sync"
	"testing"
)

func Test_Email(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	SendEmail("me@here.com", "In Nomine Patris...", "Et Filii...", make(chan error), &wg)
}
