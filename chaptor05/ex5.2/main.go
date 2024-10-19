package main

import "fmt"

func Decide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Decide(9, 3)
	fmt.Println(c, success) // 3 true

	d, success := Decide(9, 0)
	fmt.Println(d, success) // 0 false
}