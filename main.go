package main

import "fmt"

func main() {
	go PrintThis("In Nomine Patris...")
	PrintThis("Et Filii...")
	PrintThis("Et Spiritus Sancti!")
}

func PrintThis(s string) {
	fmt.Println(s)
}
