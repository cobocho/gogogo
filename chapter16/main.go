package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := []Student{
		{"홍길동", 23}, {"김길동", 21}, {"이길동", 25},
		{"박길동", 22}, {"최길동", 24}, {"정길동", 20},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}