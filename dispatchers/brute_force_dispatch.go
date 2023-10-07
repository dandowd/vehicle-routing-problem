package dispatchers

import (
	"vehicle-routing-problem/entities"
)

type BruteForceDispatch struct {
	loadQueue []*entities.Load
}

func NewBruteForceDispatch(loads []*entities.Load) *BruteForceDispatch {
	return &BruteForceDispatch{loads}
}

func (d *BruteForceDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver

	for _, load := range d.loadQueue {
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

  for _, driver := range drivers {
    driver.ReturnToOrigin()
  }

	return drivers
}
