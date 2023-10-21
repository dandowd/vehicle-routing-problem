package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
	"vehicle-routing-problem/visualization"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <load-file> [visualization-name]")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	drivers := dispatchers.Annealing(loads)

	if len(os.Args) == 3 {
		visualizationName := os.Args[2]
		
		title := fmt.Sprintf("Total Cost: %f", dispatchers.GetTotalCost(drivers))
		visualization.Route(drivers, title, fmt.Sprint("annealing-", visualizationName))
	}
}
