package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var f = 10.8271

	fmt.Print("a:", a, "b:", b) // "a:10b:20"
	fmt.Println("a:", a, "b:", b, "f:", f) // "a: 10 b: 20 f: 10.8271"
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f) // "a: 10 b: 20 f: 10.827100"
}