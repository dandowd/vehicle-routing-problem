package dispatchers

import "vehicle-routing-problem/entities"

type Dispatcher interface {
	SearchForRoutes() []*entities.Driver
}

func returnAllDriversToOrigin(drivers []*entities.Driver) []*entities.Driver {
	for _, driver := range drivers {
		driver.ReturnToOrigin()
	}
	return drivers
}