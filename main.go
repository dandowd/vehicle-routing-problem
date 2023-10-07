package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program takes exactly one argument, the path to the problem file.")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	dispatcher := dispatchers.NewNearestDriverDispatch(loads, 400)
	drivers := dispatcher.SearchForRoutes()

	totalTime := 0.0
	for _, driver := range drivers {
		totalTime += driver.GetTotalTime()
	}

	total := float64(500*len(drivers)) + totalTime
	cli.FormatDrivers(drivers)
	fmt.Println("Total cost:", total)
}
