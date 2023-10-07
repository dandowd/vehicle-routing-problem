package dispatchers 

import (
	"vehicle-routing-problem/entities"
)

type QuadDispatch struct {
	allLoads       []*entities.Load
	remainingLoads []*entities.Load
	quadMap        *entities.QuadMap
}

func NewQuadDispatch(loads []*entities.Load) *QuadDispatch {
	return &QuadDispatch{loads, []*entities.Load{}, entities.NewQuadMap(loads)}
}

func (d *QuadDispatch) SearchForRoutes() []*entities.Driver {
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

	for _, driver := range drivers {
		driver.ReturnToOrigin()
	}

	return drivers
}
