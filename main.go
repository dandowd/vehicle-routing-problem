package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
	"vehicle-routing-problem/entities"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program takes exactly one argument, the path to the problem file.")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	drivers := dispatchers.NewNearestLoadDispatch(loads).SearchForRoutes()

	cli.FormatDrivers(drivers)
}

func RunDispatchers(loads []*entities.Load) []*entities.Driver {
	dispatchers := []dispatchers.Dispatcher{
		dispatchers.NewDriverUtilizationDispatcher(loads),
		dispatchers.NewNearestLoadDispatch(loads),
	}

	var bestDrivers []*entities.Driver
	for _, dispatcher := range dispatchers {
		drivers := dispatcher.SearchForRoutes()

		totalCost := 0.0
		for _, driver := range drivers {
			totalCost += driver.GetTotalTime()
		}

		totalCost += 500 * float64(len(drivers))

		fmt.Printf("Total cost: %f, %T\n", totalCost, dispatcher)
	}

	return bestDrivers
}
