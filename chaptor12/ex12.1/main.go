package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func main() {
  p := new(Person)

  p.Name = "Alice"
  p.Age = 25

  fmt.Println(p.Name, p.Age) 
}