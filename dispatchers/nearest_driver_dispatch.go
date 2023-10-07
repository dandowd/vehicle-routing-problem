package dispatchers

import "vehicle-routing-problem/entities"

type NearestDriverDispatch struct {
  loads []*entities.Load
  maxRange float64
}

func NewNearestDriverDispatch(loads []*entities.Load, maxRange float64) *NearestDriverDispatch {
  return &NearestDriverDispatch{loads: loads, maxRange: maxRange}
}

func (d *NearestDriverDispatch) SearchForRoutes() []*entities.Driver {
  var drivers []*entities.Driver
  for _, load := range d.loads {
    var closestDriver *entities.Driver

    for _, driver := range drivers {
      distToPickup := driver.DistanceTo(load.Pickup)

      if distToPickup > d.maxRange {
        continue
      }

      if !driver.CanMoveLoad(load) {
        continue
      }

      if distToPickup < closestDriver.DistanceTo(load.Pickup) {
        closestDriver = driver
      }
    }

    if closestDriver == nil {
      closestDriver = entities.NewDriver()
      drivers = append(drivers, closestDriver)
    }
    
    closestDriver.MoveLoad(load)
  }

  return drivers
}
