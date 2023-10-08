package dispatchers

import (
	"vehicle-routing-problem/entities"
)

type QuadDispatch struct {
	allLoads []*entities.Load
	quadMap  *entities.QuadMap
}

func NewQuadDispatch(loads []*entities.Load) *QuadDispatch {
	return &QuadDispatch{loads, entities.NewQuadMap(loads)}
}

func (d *QuadDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver

	for _, quad := range d.quadMap.GetChildren() {
		for _, load := range quad.GetLoads() {
			var driver *entities.Driver
			for _, d := range drivers {
				if d.CanMoveLoad(load) {
					driver = d
					break
				}
			}

			if driver == nil {
				driver = entities.NewDriver()
				drivers = append(drivers, driver)
			}

			driver.MoveLoad(load)
		}
	}

	returnAllDriversToOrigin(drivers)

	return drivers
}
