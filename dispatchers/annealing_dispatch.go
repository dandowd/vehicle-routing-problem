package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/entities"
)

func Annealing(startingLoads []*entities.Load) []*entities.Driver {
	bestDrivers := []*entities.Driver{}
	totalCost := math.MaxFloat64
	path := startingLoads 

	temperature := 1000.0
	coolingRate := 0.98

	for i := 0; i <= 20000; i++ {
		randomSwap(path)

		newDrivers := driveRoute(path)
		newCost := GetTotalCost(newDrivers)

		if shouldTakeNewPath(totalCost, newCost, temperature) {
			totalCost = newCost
			bestDrivers = newDrivers
		}

		path = CombineDriverLoads(bestDrivers)

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

	probability := math.Exp((oldCost - newCost) / temperature)
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