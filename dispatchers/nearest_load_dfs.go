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

// Recursively check paths until the path with the least waste is found
// Base case: once a driver reaches a node with all adjacent nodes being unable to be moved
func (d *NearestLoadDFSDispatch) search(driver *entities.Driver, travelCosts float64) (float64, *entities.Driver) {
  nearestLoads := d.getNearestLoads(driver)
  bestDriver := driver.MakeCopy()
  bestCosts := math.MaxFloat64

  for _, loadMap := range nearestLoads {
    if !driver.CanMoveLoad(loadMap) {
      continue
    }

    costs, subPath := d.search(driver.MakeCopy(), travelCosts)

    if costs < bestCosts {
      bestDriver = subPath
    }
  }

  if bestCosts == math.MaxFloat64 {
    return travelCosts, bestDriver
  }

  return travelCosts + bestCosts, bestDriver
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

func copyPath(path map[int]*entities.Load) map[int]*entities.Load {
  newPath := make(map[int]*entities.Load)
  for k, v := range path {
    newPath[k] = v
  }
  return newPath
}
