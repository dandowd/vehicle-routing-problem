package strategies

import "vehicle-routing-problem/entities"

func NearestLoadStrategy(driver *entities.Driver, availableLoads map[int]*entities.Load) *entities.Load {
	var nearestLoad *entities.Load

	for _, load := range availableLoads {
		if !driver.CanMoveLoad(load) {
			continue
		}

		if nearestLoad == nil {
			nearestLoad = load
		}

		if load.Pickup.DistanceTo(driver.GetCurrentPoint()) < nearestLoad.Pickup.DistanceTo(driver.GetCurrentPoint()) {
			nearestLoad = load
		}
	}

	return nearestLoad
}
