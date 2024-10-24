package main

import "fmt"

func main() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y =", x + y) // 10
	fmt.Println("x - y =", x - y) // 4
	fmt.Println("x * y =", x * y) // 21
	fmt.Println("x / y =", x / y) // 2
	fmt.Println("x % y =", x % y) // 1

	fmt.Println("s * t =", s * t) // 15.7
	fmt.Println("s / t =", s / t) // 0.628
}