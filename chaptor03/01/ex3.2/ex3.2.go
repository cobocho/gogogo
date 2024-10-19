package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 1234556789

	fmt.Printf("%5d, %5d\n", a, b) // "  123,   456"
	fmt.Printf("%05d, %05d\n", a, b) // "00123, 00456"
	fmt.Printf("%5d, %-05d\n", a, b) // "  123, 456  "

	fmt.Printf("%5d, %5d\n", c, c) // "1234556789, 1234556789"
	fmt.Printf("%05d, %05d\n", c, c) // "1234556789, 1234556789"
}