package main

import "fmt"

const ResetColor = "\033[0m"
const Cyan = "\033[36m"

func ColorPrint(message, color string) {
	fmt.Printf("%s%s%s\n", color, message, ResetColor)
}
