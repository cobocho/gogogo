package main

import (
	"chaptor14/usepkg/custompkg"
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3,4,5,6,8,6,5,8,5,10,2,7,2,5,6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}