package dispatchers

import "vehicle-routing-problem/entities"

type BruteForceDispatch struct {
	loadQueue []*entities.Load
}

func NewBruteForceDispatch(loads []*entities.Load) *BruteForceDispatch {
	return &BruteForceDispatch{loads}
}

func (d *BruteForceDispatch) SearchForRoutes() []*entities.Driver {
	var drivers []*entities.Driver

	for _, load := range d.loadQueue {
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

  for _, driver := range drivers {
    driver.ReturnToOrigin()
  }

	return drivers
}
