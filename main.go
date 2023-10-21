package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
	"vehicle-routing-problem/visualization"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program takes exactly one argument, the path to the problem file.")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	drivers := dispatchers.Annealing(loads)
	title := fmt.Sprintf("Total Cost: %f", dispatchers.GetTotalCost(drivers))

	visualization.Route(drivers, title)
}
