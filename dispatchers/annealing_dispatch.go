package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/entities"
)

type AnnealingOptions struct {
	Iterations   int
	StartingTemp float64
	CoolingRate  float64
	Schedule     int
}

func AnnealDrivers(drivers []*entities.Driver) []*entities.Driver {
	optimzedDrivers := []*entities.Driver{}
	for _, driver := range drivers {
		drivers := Annealing(
			driver.GetPath(),
			&AnnealingOptions{Iterations: 1000, StartingTemp: 1000, CoolingRate: 0.97, Schedule: 100},
		)

		optimzedDrivers = append(optimzedDrivers, drivers...)
	}

	return optimzedDrivers
}

func Annealing(startingLoads []*entities.Load, options *AnnealingOptions) []*entities.Driver {
	explorationDrivers := []*entities.Driver{}
	bestExplorationCost := math.MaxFloat64

	bestOverallDrivers := []*entities.Driver{}
	bestOverallCost := math.MaxFloat64
	path := startingLoads

	temperature := options.StartingTemp

	for i := 0; i <= options.Iterations; i++ {
		randomSwap(path)

		newDrivers := driveRoute(path)
		newCost := GetTotalCost(newDrivers)

		if bestExplorationCost < bestOverallCost {
			bestOverallCost = bestExplorationCost
			bestOverallDrivers = explorationDrivers
		}

		if shouldExploreNewPath(bestExplorationCost, newCost, temperature) {
			bestExplorationCost = newCost
			explorationDrivers = newDrivers
		}

		path = CombineDriverLoads(explorationDrivers)

		if i%options.Schedule == 0 {
			temperature *= options.CoolingRate
		}
	}

	return bestOverallDrivers
}

func shouldExploreNewPath(oldCost float64, newCost float64, temperature float64) bool {
	if newCost < oldCost {
		return true
	}

	probability := math.Exp((oldCost - newCost) / temperature)
	return rand.Float64() < probability
}

func reverseConsecutiveLoads(loads []*entities.Load) {
	firstIndex := rand.Intn(len(loads))

	if firstIndex == len(loads)-1 {
		return
	}

	temp := loads[firstIndex]
	loads[firstIndex] = loads[firstIndex+1]
	loads[firstIndex+1] = temp
}

func randomSwap(loads []*entities.Load) {
	firstIndex := rand.Intn(len(loads))
	secondIndex := rand.Intn(len(loads))

	temp := loads[firstIndex]

	loads[firstIndex] = loads[secondIndex]
	loads[secondIndex] = temp
}

func reverseRandomSegment(loads []*entities.Load) {
	firstIndex := rand.Intn(len(loads))
	secondIndex := rand.Intn(len(loads))

	if firstIndex > secondIndex {
		temp := firstIndex
		firstIndex = secondIndex
		secondIndex = temp
	}

	for i := 0; i < (secondIndex-firstIndex)/2; i++ {
		temp := loads[firstIndex+i]

		loads[firstIndex+i] = loads[secondIndex-i]
		loads[secondIndex-i] = temp
	}
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
