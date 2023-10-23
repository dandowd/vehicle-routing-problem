package dispatchers

import (
	"math"
	"math/rand"
	"vehicle-routing-problem/entities"
	"vehicle-routing-problem/visualization"
)

func PathAnnealing(startingLoads []*entities.Load, iterations int, startingTemp float64, coolingRate float64, schedule int) []*entities.Driver {
	pathTracker := NewPathCostTracker(startingLoads)

	costIterationLog := visualization.NewGraphLog()

	bestExplorationCost := math.MaxFloat64
	temperature := startingTemp

	for i := 0; i <= iterations; i++ {
		pathTracker.RandomSwap()
		newCost := pathTracker.GetCost()
		costIterationLog.AddPoint(float64(i), newCost)

		if shouldExploreNewPath(bestExplorationCost, newCost, temperature) {
			bestExplorationCost = newCost
		} else {
			pathTracker.UndoSwap()
		}

		if i%schedule == 0 {
			temperature *= coolingRate
		}
	}

	costIterationLog.CreateFile("annealing_dispatch_cost_graph")

	return driveRoute(pathTracker.path)
}

type PathCostTracker struct {
	path  []*entities.Load
	cost  float64
	swapI int
	swapK int
}

func NewPathCostTracker(path []*entities.Load) *PathCostTracker {
	tracker := &PathCostTracker{path: []*entities.Load{}}

	for _, load := range path {
		tracker.AddLoad(load)
	}

	return tracker
}

func (tracker *PathCostTracker) EstimateDriverCost() float64 {
	return GetTotalCost(driveRoute(tracker.path))
}

func (tracker *PathCostTracker) RandomSwap() {
	tracker.swapI = rand.Intn(len(tracker.path))
	tracker.swapK = rand.Intn(len(tracker.path))

	tracker.SwapLoads(tracker.swapI, tracker.swapK)
}

func (tracker *PathCostTracker) UndoSwap() {
	tracker.SwapLoads(tracker.swapI, tracker.swapK)
}

func (tracker *PathCostTracker) GetCost() float64 {
	return tracker.cost
}

func (tracker *PathCostTracker) AddLoad(load *entities.Load) {
	tracker.path = append(tracker.path, load)
	tracker.addCost(len(tracker.path) - 1)
}

func (tracker *PathCostTracker) SwapLoads(i, k int) {
	temp := tracker.path[i]

	tracker.removeCost(i)
	tracker.path[i] = tracker.path[k]
	tracker.addCost(i)

	tracker.removeCost(k)
	tracker.path[k] = temp
	tracker.addCost(k)
}

func (tracker *PathCostTracker) removeCost(i int) {
	tracker.cost -= tracker.path[i].GetTime()

	if i > 0 {
		tracker.cost -= tracker.path[i-1].Dropoff.DistanceTo(tracker.path[i].Pickup)
	}

	if i < len(tracker.path)-1 {
		tracker.cost -= tracker.path[i].Dropoff.DistanceTo(tracker.path[i+1].Pickup)
	}
}

func (tracker *PathCostTracker) addCost(i int) {
	tracker.cost += tracker.path[i].GetTime()

	if i > 0 {
		tracker.cost += tracker.path[i-1].Dropoff.DistanceTo(tracker.path[i].Pickup)
	}

	if i < len(tracker.path)-1 {
		tracker.cost += tracker.path[i].Dropoff.DistanceTo(tracker.path[i+1].Pickup)
	}
}

func (tracker *PathCostTracker) insertLoad(i int, load *entities.Load) {
	tracker.path[i] = load
	tracker.addCost(i)
}
