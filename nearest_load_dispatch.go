package main

import "vehicle-routing-problem/entities"

type NearestLoadDispatch struct {
	loads map[int]*entities.Load
}

func NewNearestLoadDispatch(loads []*entities.Load) *NearestLoadDispatch {
	dispatch := &NearestLoadDispatch{loads: make(map[int]*entities.Load)}
	for _, load := range loads {
		dispatch.loads[load.LoadNumber] = load
	}

	return dispatch
}

// This optimizes for distance between dropoff and next pickup, minimizing wasted driving.
func (d *NearestLoadDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver
	for len(d.loads) > 0 {
		driver := entities.NewDriver()
		drivers = append(drivers, driver)

		for {
			load := d.getNearestLoadToPoint(driver)

			if load == nil {
				break
			}

			driver.MoveLoad(load)
			delete(d.loads, load.LoadNumber)
		}
	}

	for _, driver := range drivers {
		driver.ReturnToOrigin()
	}

	return drivers
}

func (d *NearestLoadDispatch) getNearestLoadToPoint(driver *entities.Driver) *entities.Load {
	var nearestLoad *entities.Load

	for _, load := range d.loads {
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
