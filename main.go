package main

import (
	"fmt"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
)

func main() {
	loads := cli.ParseLoadFile("/Users/danieldowd/Downloads/problem12.txt")
	dispatcher := dispatchers.NewBruteForceDispatch(loads)
	drivers := dispatcher.SearchForRoutes()

	totalTime := 0.0
	for _, driver := range drivers {
		totalTime += driver.GetTotalTime()
	}

	total := float64(500*len(drivers)) + totalTime
	cli.FormatDrivers(drivers)
	fmt.Println("Total cost:", total)
}
