package dispatchers

import (
	"vehicle-routing-problem/entities"
)

type BruteForceDispatch struct {
	loads []*entities.Load
}

func NewBruteForceDispatch(loads []*entities.Load) *BruteForceDispatch {
	return &BruteForceDispatch{loads}
}

func (d *BruteForceDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver

	for _, load := range d.loads {
		var selectedDriver *entities.Driver
		for _, currentDriver := range drivers {
			if currentDriver.CanMoveLoad(load) {
				selectedDriver = currentDriver
				break
			}
		}

		if selectedDriver == nil {
			selectedDriver = entities.NewDriver() 
			drivers = append(drivers, selectedDriver)
		}

		selectedDriver.MoveLoad(load)
	}

	returnAllDriversToOrigin(drivers)

	return drivers
}
