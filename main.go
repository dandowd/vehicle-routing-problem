package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
	"vehicle-routing-problem/entities"
	"vehicle-routing-problem/strategies"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program takes exactly one argument, the path to the problem file.")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	drivers := RunDispatchers(loads)

	cli.FormatDrivers(drivers)
}

func RunDispatchers(startingLoads []*entities.Load) []*entities.Driver {
	bestDrivers := dispatchers.NewNearestLoadDispatch(startingLoads).SearchForRoutes()
	bestTotalCost := getTotalCost(bestDrivers)

	driverUtil := dispatchers.NewDriverUtilizationDispatch(startingLoads).SearchForRoutes()
	if getTotalCost(driverUtil) < bestTotalCost {
		bestDrivers = driverUtil
	}

	for i := 0; i < len(startingLoads); i++ {
		loads := rotateSlice(startingLoads, i)
		dispatchers := []dispatchers.Dispatcher{
			dispatchers.NewBruteForceDispatch(loads),
			dispatchers.NewNearestDriverDispatch(loads, 400),
		}

		for _, dispatcher := range dispatchers {
			drivers := dispatcher.SearchForRoutes()

			totalCost := getTotalCost(drivers)

			if totalCost < bestTotalCost || bestTotalCost == 0 {
				bestTotalCost = totalCost
				bestDrivers = drivers
			}
		}
	}

	return bestDrivers
}

func RunStrategies(startingLoads []*entities.Load) []*entities.Driver {
	loads := NewLoadMap(startingLoads) 
	drivers := []*entities.Driver{}

	for len(loads) > 0 {
		driver := entities.NewDriver()
		drivers = append(drivers, driver)

		for len(loads) > 0 {
			load := strategies.NearestLoadStrategy(driver, loads)
			if load == nil {
				break
			}

			driver.MoveLoad(load)
			delete(loads, load.LoadNumber)
		}
	}

	return drivers
}

func NewLoadMap(loads []*entities.Load) map[int]*entities.Load {
	loadMap := map[int]*entities.Load{}
	for _, load := range loads {
		loadMap[load.LoadNumber] = load
	}

	return loadMap
}

func getTotalCost(drivers []*entities.Driver) float64 {
	totalCost := 0.0
	for _, driver := range drivers {
		totalCost += driver.GetTotalTime()
	}

	totalCost += 500 * float64(len(drivers))

	return totalCost
}

func rotateSlice(slice []*entities.Load, n int) []*entities.Load {
	return append(slice[n:], slice[:n]...)
}
