package main

import "fmt"

const ResetColor = "\033[0m"

const BrightGreen = "\033[92m"
const Cyan = "\033[36m"
const Green = "\033[32m"
const Red = "\033[31m"

func ColorPrint(color, message string) {
	fmt.Printf("%s%s%s\n", color, message, ResetColor)
}
