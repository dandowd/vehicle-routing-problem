package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/entities"
	"vehicle-routing-problem/visualization"
)

var logger = cli.NewFileLogger("annealing_dispatch.log")

func Annealing(startingLoads []*entities.Load) []*entities.Driver {
	costIterationLog := visualization.NewGraphLog()

	explorationDrivers := []*entities.Driver{}
	bestExplorationCost := math.MaxFloat64

	bestOverallDrivers := []*entities.Driver{}
	bestOverallCost := math.MaxFloat64
	path := startingLoads

	temperature := 1000.0
	coolingRate := 0.97

	totalIterations := 1000 * len(startingLoads)
	for i := 0; i <= totalIterations; i++ {
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

		if i%100 == 0 {
			temperature *= coolingRate
		}
	}

	logger.Println("Total cost:", bestExplorationCost)
	logger.Println(cli.FormatPath(explorationDrivers))
	costIterationLog.CreateFile("annealing_dispatch_cost_graph")

	return bestOverallDrivers
}

func shouldExploreNewPath(oldCost float64, newCost float64, temperature float64) bool {
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
