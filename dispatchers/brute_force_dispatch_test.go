package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestBruteForceShouldTravelToAllLoads(t *testing.T) {
	loadOne := &entities.Load{LoadNumber: 1, Pickup: entities.Point{X: 1, Y: 1}, Dropoff: entities.Point{X: -32, Y: 24}}
	loadTwo := &entities.Load{LoadNumber: 2, Pickup: entities.Point{X: -72, Y: 80}, Dropoff: entities.Point{X: 23, Y: 4}}

	dispatch := NewBruteForceDispatch([]*entities.Load{
		loadOne,
		loadTwo,
	})

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 1 {
		t.Error("Expected 1 driver, got ", len(drivers))
	}

	if len(drivers[0].GetPath()) != 2 {
		t.Error("Expected 2 completed loads, got ", len(drivers[0].GetPath()))
	}
}
