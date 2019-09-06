package main

import (
	"fmt"
	"ratio-finder/rategraph"
)

func main() {
	rates := []rategraph.Rate{
		rategraph.Rate{"light year", "km", 9.461e12},
		rategraph.Rate{"km", "m", 1000},
		rategraph.Rate{"mile", "km", 1600},
		rategraph.Rate{"mile", "foot", 5280},
		rategraph.Rate{"foot", "inch", 12},
		rategraph.Rate{"inch", "hand", 0.25},
		rategraph.Rate{"inch", "cm", 2.54},
		rategraph.Rate{"m", "cm", 100},
	}

	rg := rategraph.NewRateGraph(rates)

	fmt.Println(rg.GetNodes())

	fmt.Println(rg.GetNeighbors("mile"))
	fmt.Println(rg.GetNeighbors("inch"))

	fmt.Println(rg.GetNeighbors("hand"))

	rg.AddConversion("hand", "mm", 101.6)

	fmt.Println(rg.GetNeighbors("hand"))
}
