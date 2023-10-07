package dispatchers

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

    canMoveNextLoad := true
    for canMoveNextLoad {
      load := d.getNearestLoadToPoint(driver.GetCurrentPoint())

      if load == nil {
        canMoveNextLoad = false
        break
      }

      if driver.CanMoveLoad(load) {
        driver.MoveLoad(load)
        delete(d.loads, load.LoadNumber)
      } else {
        canMoveNextLoad = false
      }
    }
  }

  return drivers
}

func (d *NearestLoadDispatch) getNearestLoadToPoint(point entities.Point) *entities.Load {
  var closestsToOrigin *entities.Load

  for _, load := range d.loads {
    if closestsToOrigin == nil {
      closestsToOrigin = load
      continue
    }

    if load.Pickup.DistanceTo(point) < closestsToOrigin.Pickup.DistanceTo(point) {
      closestsToOrigin = load
    }
  }

  return closestsToOrigin
}
