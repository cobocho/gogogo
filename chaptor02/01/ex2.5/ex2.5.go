package main

import "fmt"

func main() {
	a := 3 // int
	var b float64 = 3.5 // float64

	var c int = int(b) // int, 3
	d := float64(a * c) // int 3 * int 3 = int 9 -> float64 9.0

	var e int64 = 7
	f := int64(d) * e // int64 9 * int64 7 = int64 63

	var g int = int(b * 3) // float64 3.5 * 3 = float64 10.5 -> int 10
	var h int = int(b) * 3 // int 3 * 3 = int 9
	fmt.Println(g, h, f) // 10 9 63

	var i int16 = 3456
	var j int8 = int8(i) // int16 3456 -> int8 -128

	fmt.Println(i, j) // 3456 -128
}