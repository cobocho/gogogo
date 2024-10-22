package main

import "fmt"

func main() {
  var a int = 10
  var p *int

  p = &a



  fmt.Println("p의 값:", p) // p의 값: 0xc0000b6010
  fmt.Println("p가 가리키는 값:", *p) // p가 가리키는 값: 10
  *p = 100
  fmt.Println("a의 값:", a) // a의 값: 100
}