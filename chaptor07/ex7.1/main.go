package main

import "fmt"

func GetLight() string {
	return "blue"
}

func main() {
	if light := GetLight(); light == "red" {
		fmt.Printf("빨간불")
	} else if light == "blue" {
		fmt.Printf("파란불")
	} else {
		fmt.Printf("녹색불")
	}
}