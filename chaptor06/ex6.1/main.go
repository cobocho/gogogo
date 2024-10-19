package main

import "fmt"

const PI = 3.14 // untyped constant
const FloatPI float32 = 3.14

func main() {
	var a int = PI * 100
	var b int = FloatPI * 100 // 에러 발생

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}