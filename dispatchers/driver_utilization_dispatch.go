package dispatchers

import "vehicle-routing-problem/entities"

type DriverUtilizationDispatcher struct {
  loads map[int]*entities.Load
}

func NewDriverUtilizationDispatcher(loads []*entities.Load) *DriverUtilizationDispatcher {
  loadsMap := make(map[int]*entities.Load)
  for _, load := range loads {
    loadsMap[load.LoadNumber] = load
  }

  return &DriverUtilizationDispatcher{loads: loadsMap}
}

func (d *DriverUtilizationDispatcher) SearchForRoutes() []*entities.Driver {
  var drivers []*entities.Driver

  for len(d.loads) > 0 {
    driver := entities.NewDriver() 
    
    load := d.getLongestLoad(driver)
    for load != nil {
      driver.MoveLoad(load)
      delete(d.loads, load.LoadNumber)

      load = d.getLongestLoad(driver)
    }

    driver.ReturnToOrigin()
    drivers = append(drivers, driver)
  }

  return drivers
}

func (d *DriverUtilizationDispatcher) getLongestLoad(driver *entities.Driver) *entities.Load {
  var longestLoad *entities.Load
  for _, load := range d.loads {
    if driver.CanMoveLoad(load) && longestLoad == nil {
      longestLoad = load
    }

    if driver.CanMoveLoad(load) && load.GetTime() > longestLoad.GetTime() {
      longestLoad = load
    }
  }

  return longestLoad
}
