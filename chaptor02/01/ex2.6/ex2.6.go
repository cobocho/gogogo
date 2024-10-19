package main

import "fmt"

var global = "global"

func main() {
	var function = "function"

	{
		var block = "block"
		fmt.Println(global, function, block) // global function block
	}

}