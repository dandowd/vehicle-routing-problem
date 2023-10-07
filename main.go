package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
)

func main() {
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
