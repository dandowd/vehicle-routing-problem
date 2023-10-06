package main

import (
	"vehicle-routing-problem/entities"
)

type Dispatch struct {
	allLoads       []*entities.Load
	completedLoads []*entities.Load
	quadMap        *entities.QuadMap
}

func NewDispatch(loads []*entities.Load) *Dispatch {
	return &Dispatch{loads, []*entities.Load{}, entities.NewQuadMap(loads)}
}

func (d *Dispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver

	for _, quad := range d.quadMap.GetChildren() {
		for _, load := range quad.GetLoads() {
			var driver *entities.Driver
			for _, d := range drivers {
				if d.CanPickup(load) {
					driver = d
					break
				}
			}

			if driver == nil {
				driver = &entities.Driver{}
				drivers = append(drivers, driver)
			}

			driver.AddLoad(load)
		}
	}

	return drivers 
}
