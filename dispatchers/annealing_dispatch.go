package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/entities"
	"vehicle-routing-problem/visualization"
)

func AnnealDrivers(drivers []*entities.Driver) []*entities.Driver {
	optimzedDrivers := []*entities.Driver{}
	for _, driver := range drivers {
		drivers := Annealing(driver.GetPath(), 20000, 1000, 0.95, 100)
		optimzedDrivers = append(optimzedDrivers, drivers...)
	}

	return optimzedDrivers
}

func Annealing(startingLoads []*entities.Load, iterations int, startingTemp float64, coolingRate float64, schedule int) []*entities.Driver {
	costIterationLog := visualization.NewGraphLog()

	explorationDrivers := []*entities.Driver{}
	bestExplorationCost := math.MaxFloat64

	bestOverallDrivers := []*entities.Driver{}
	bestOverallCost := math.MaxFloat64
	path := startingLoads

	temperature := startingTemp 

	for i := 0; i <= iterations; i++ {
		randomSwap(path)

		newDrivers := driveRoute(path)
		newCost := GetTotalCost(newDrivers)

		if bestExplorationCost < bestOverallCost {
			bestOverallCost = bestExplorationCost
			bestOverallDrivers = explorationDrivers
		}

		if shouldExploreNewPath(bestExplorationCost, newCost, temperature) {
			costIterationLog.AddPoint(float64(i), newCost)

			bestExplorationCost = newCost
			explorationDrivers = newDrivers
		}

		path = CombineDriverLoads(explorationDrivers)

		if i%schedule == 0 {
			temperature *= coolingRate
		}
	}

	costIterationLog.CreateFile("annealing_dispatch_cost_graph")

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
