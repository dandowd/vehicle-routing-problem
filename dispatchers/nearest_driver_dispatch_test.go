package dispatchers

import (
	"testing"
  "vehicle-routing-problem/entities"
)

func TestSearchForRoutesShouldUseNearestDriver(t *testing.T) {
  loadOne := entities.Load{Pickup: entities.Point{X: 30, Y: 30}, Dropoff: entities.Point{X: 200, Y: 200}} 
  loadTwo := entities.Load{Pickup: entities.Point{X: -200, Y: 200}, Dropoff: entities.Point{X: 150, Y: 150}}

  targetLoad := entities.Load{Pickup: entities.Point{X: 300, Y: 300}, Dropoff: entities.Point{X: 2, Y: 2}}

  dispatch := NearestDriverDispatch{loads: []*entities.Load{&loadOne, &loadTwo, &targetLoad}}

  drivers := dispatch.SearchForRoutes()

  if len(drivers) != 2 {
    t.Errorf("Expected 2 drivers, got %d", len(drivers))
  }
}
