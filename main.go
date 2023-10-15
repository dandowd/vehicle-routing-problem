package main

import (
	"fmt"
	"math"
	"math/rand"
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

	drivers := RunDispatchers(loads)

	cli.Logger.Println("Total cost:", getTotalCost(drivers))
	cli.Logger.Println(cli.FormatPath(drivers))
}

func annealing(startingLoads []*entities.Load) []*entities.Driver {
	bestDrivers := dispatchers.NewNearestLoadDispatch(startingLoads).SearchForRoutes()
	totalCost := getTotalCost(bestDrivers)
	path := combineDriverLoads(bestDrivers)

	temperature := 100.0
	coolingRate := 0.95

	for i := 0; i <= 200000; i++ {
		randomSwap(path)

		newDrivers := driveRoute(path)
		newCost := getTotalCost(newDrivers)

		if shouldTakeNewPath(totalCost, newCost, temperature) {
			totalCost = newCost
			bestDrivers = newDrivers
		}

		path = combineDriverLoads(bestDrivers)

		if i % 100 == 0 {
			temperature *= coolingRate
		}
	}


	return bestDrivers
}

func shouldTakeNewPath(oldCost float64, newCost float64, temperature float64) bool {
	if newCost < oldCost {
		return true
	}

	probability := math.Exp(-(newCost - oldCost) / temperature)
	return rand.Float64() < probability
}

func randomSwap(loads []*entities.Load) {
	firstIndex := rand.Intn(len(loads))
	secondIndex := rand.Intn(len(loads))

	temp := loads[firstIndex]

	loads[firstIndex] = loads[secondIndex]
	loads[secondIndex] = temp
}

func driveRoute(loads []*entities.Load) []*entities.Driver {
	drivers := []*entities.Driver{}

	for len(loads) > 0 {
		driver := entities.NewDriver()

		for len(loads) > 0 && driver.CanMoveLoad(loads[0]) {
			driver.MoveLoad(loads[0])
			loads = loads[1:]
		}
		driver.ReturnToOrigin()

		drivers = append(drivers, driver)
	}

	return drivers
}


func combineDriverLoads(drivers []*entities.Driver) []*entities.Load {
	loads := []*entities.Load{}
	for _, driver := range drivers {
		loads = append(loads, driver.GetPath()...)
	}

	return loads
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
