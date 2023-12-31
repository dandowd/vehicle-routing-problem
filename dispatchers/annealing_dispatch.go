package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/entities"
	"vehicle-routing-problem/visualization"
)

type AnnealingOptions struct {
	Iterations   int
	StartingTemp float64
	CoolingRate  float64
	Schedule     int
}

func AnnealDrivers(drivers []*entities.Driver, options *AnnealingOptions) []*entities.Driver {
	optimzedDrivers := []*entities.Driver{}
	for _, driver := range drivers {
		if len(driver.GetPath()) > 1 {
			drivers := Annealing(
				driver.GetPath(),
				options,
			)
			optimzedDrivers = append(optimzedDrivers, drivers...)
		} else {
			optimzedDrivers = append(optimzedDrivers, driver)
		}
	}

	return optimzedDrivers
}

func Annealing(startingLoads []*entities.Load, options *AnnealingOptions) []*entities.Driver {
	graphLog := visualization.NewGraphLog()
	tempLog := visualization.NewGraphLog()

	explorationDrivers := driveRoute(startingLoads)
	explorationCost := GetTotalCost(explorationDrivers)

	bestOverallDrivers := explorationDrivers
	bestOverallCost := explorationCost
	path := startingLoads

	for reheat := 0; reheat < 5; reheat++ {
		temperature := options.StartingTemp

		for i := 0; i <= options.Iterations; i++ {
			randomSwap(path)

			newDrivers := driveRoute(path)
			newCost := GetTotalCost(newDrivers)

			if shouldExploreNewPath(explorationCost, newCost, temperature) {
				graphLog.AddPoint(float64(i), newCost)

				explorationCost = newCost
				explorationDrivers = newDrivers
			}

			path = CombineDriverLoads(explorationDrivers)

			if i%options.Schedule == 0 {
				temperature *= options.CoolingRate
			}
			tempLog.AddPoint(float64(i), temperature)

			if newCost < bestOverallCost {
				bestOverallCost = newCost
				bestOverallDrivers = newDrivers
				i = 0
			}
		}
	}

	tempLog.CreateFile("annealing_temp")
	graphLog.CreateFile("annealing_graph")
	return bestOverallDrivers
}

func multipleShuffle(loads []*entities.Load, count int) []*entities.Load {
	for i := 0; i < count; i++ {
		loads = shuffleLoads(loads)
	}

	return loads
}

func shuffleLoads(loads []*entities.Load) []*entities.Load {
	shuffledLoads := []*entities.Load{}
	for len(loads) > 0 {
		index := rand.Intn(len(loads))
		shuffledLoads = append(shuffledLoads, loads[index])
		loads = append(loads[:index], loads[index+1:]...)
	}

	return shuffledLoads
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
