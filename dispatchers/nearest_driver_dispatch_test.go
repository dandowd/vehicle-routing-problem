package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestSearchForRoutesShouldUseNearestDriver(t *testing.T) {
	loadOne := entities.Load{Pickup: entities.Point{X: 30, Y: 30}, Dropoff: entities.Point{X: 1, Y: 1}}
	targetLoad := entities.Load{Pickup: entities.Point{X: 1, Y: 1}, Dropoff: entities.Point{X: 5, Y: 5}}

	dispatch := NewNearestDriverDispatch([]*entities.Load{&loadOne, &targetLoad}, 5)

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 1 {
		t.Errorf("Expected 1 drivers, got %d", len(drivers))
	}
}

func TestSearchForRoutesShouldCreateNewDriver(t *testing.T) {
	loadOne := entities.Load{Pickup: entities.Point{X: 30, Y: 30}, Dropoff: entities.Point{X: 1, Y: 1}}
	loadTwo := entities.Load{Pickup: entities.Point{X: -200, Y: 200}, Dropoff: entities.Point{X: 150, Y: 150}}

	targetLoad := entities.Load{Pickup: entities.Point{X: 1, Y: 1}, Dropoff: entities.Point{X: 5, Y: 5}}

	dispatch := NewNearestDriverDispatch([]*entities.Load{&loadOne, &loadTwo, &targetLoad}, 5)

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 2 {
		t.Errorf("Expected 2 drivers, got %d", len(drivers))
	}
}
