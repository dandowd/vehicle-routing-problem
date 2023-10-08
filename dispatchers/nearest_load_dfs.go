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
func (d *NearestLoadDFSDispatch) search(currentPoint entities.Point, travelCosts float64, totalCost float64, path map[int]*entities.Load) (float64, map[int]*entities.Load) {
  nearestLoads := d.getNearestLoads(currentPoint, path)
  bestPath := path
  bestCosts := math.MaxFloat64

  for _, loadMap := range nearestLoads {
    pickupCosts := currentPoint.DistanceTo(loadMap.Pickup)
    returnCosts := loadMap.Dropoff.DistanceTo(entities.Point{X: 0, Y: 0})

    tripCosts := pickupCosts + returnCosts + totalCost + loadMap.GetTime()

    if tripCosts > entities.MAX_DRIVE_TIME {
      continue
    }

    currentPath := copyPath(path)

    costs, subPath := d.search(loadMap.Dropoff, travelCosts, tripCosts + totalCost, currentPath)

    if costs < bestCosts {
      bestPath = subPath
    }
  }

  if bestCosts == math.MaxFloat64 {
    return travelCosts, path
  }

  return travelCosts + bestCosts, bestPath
}

func (d *NearestLoadDFSDispatch) getNearestLoads(point entities.Point, path map[int]*entities.Load) []*entities.Load {
	var nearestLoads []*entities.Load

	for _, load := range d.loads {
		if path[load.LoadNumber] != nil {
			continue
		}

		if len(nearestLoads) < d.numberOfAdjacent {
			nearestLoads = append(nearestLoads, load)
			continue
		}

		dist := point.DistanceTo(load.Pickup)
		for i, currentNearest := range nearestLoads {
      currentDist := point.DistanceTo(currentNearest.Pickup)
			if dist < currentDist {
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
