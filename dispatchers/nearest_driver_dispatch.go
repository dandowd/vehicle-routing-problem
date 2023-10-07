package dispatchers

import "vehicle-routing-problem/entities"

type NearestDriverDispatch struct {
  loads []*entities.Load
}

func (d *NearestDriverDispatch) SearchForRoutes(loads []*entities.Load) []*entities.Driver {
  var drivers []*entities.Driver
  for _, load := range d.loads {
    var closestDriver *entities.Driver

    for _, driver := range drivers {
      if driver.CanMoveLoad(load) {
        distToPickup := driver.DistanceTo(load.Pickup)

        if closestDriver == nil {
          closestDriver = driver
        } else if distToPickup < closestDriver.DistanceTo(load.Pickup) {
          closestDriver = driver
        }
      }
    }

    if closestDriver != nil {
      closestDriver = entities.CreateNewDriver()
    }
  }

  return drivers
}
