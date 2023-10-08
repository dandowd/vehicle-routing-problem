package dispatchers

import (
	"math"
	"vehicle-routing-problem/entities"
)

type NearestLoadDFSDispatch struct {
	loads            map[int]*entities.Load
	numberOfAdjacent int
}

func NewNearestLoadDFSDispatch(loads []*entities.Load, numberOfAdjacent int) *NearestLoadDFSDispatch {
	dispatch := &NearestLoadDFSDispatch{loads: make(map[int]*entities.Load), numberOfAdjacent: numberOfAdjacent}
	for _, load := range loads {
		dispatch.loads[load.LoadNumber] = load
	}

	return dispatch
}

func (d *NearestLoadDFSDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver
	for len(d.loads) > 0 {
		_, driver := d.search(entities.NewDriver(), 0)
		d.removeDriverLoads(driver)

		if len(driver.GetCompletedLoads()) != 0 {
			drivers = append(drivers, driver)
		}
	}

	return drivers
}

// Recursively check paths until the path with the least waste is found
// Base case: once a driver reaches a node with all adjacent nodes being unable to be moved
func (d *NearestLoadDFSDispatch) search(driver *entities.Driver, betweenTravelTime float64) (float64, *entities.Driver) {
	bestDriver := driver
	bestTravelCosts := math.MaxFloat64

	nearestLoads := d.getNearestLoads(driver)
	for _, load := range nearestLoads {
		if !driver.CanMoveLoad(load) {
			continue
		}

		copiedDriver := driver.MakeCopy()
		copiedDriver.MoveLoad(load)
		subCosts, subDriver := d.search(copiedDriver, betweenTravelTime+driver.DistanceTo(load.Pickup))

		if subCosts < bestTravelCosts {
			bestTravelCosts = subCosts
			bestDriver = subDriver
		}
	}

	if bestTravelCosts == math.MaxFloat64 {
		completedDriver := driver.MakeCopy()
		completedDriver.ReturnToOrigin()

		return betweenTravelTime, completedDriver
	}

	return bestTravelCosts, bestDriver
}

func (d *NearestLoadDFSDispatch) getNearestLoads(driver *entities.Driver) []*entities.Load {
	var nearestLoads []*entities.Load
	for _, load := range d.loads {
		if driver.HasCompletedLoad(load) {
			continue
		}

		if len(nearestLoads) < d.numberOfAdjacent {
			nearestLoads = append(nearestLoads, load)
			continue
		}

		newDist := driver.DistanceTo(load.Pickup)

		for i, currentNearest := range nearestLoads {
			currentDist := driver.DistanceTo(currentNearest.Pickup)
			if newDist < currentDist {
				nearestLoads[i] = load
			}
		}
	}

	return nearestLoads
}

func (d *NearestLoadDFSDispatch) removeDriverLoads(driver *entities.Driver) {
	for _, load := range driver.GetCompletedLoads() {
		delete(d.loads, load.LoadNumber)
	}
}
