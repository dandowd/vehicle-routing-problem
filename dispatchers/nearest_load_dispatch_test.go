package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestNearestLoadDispatchShouldPickupInCorrectOrder(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 1, Y: 0}, entities.Point{X: 2, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 6, Y: 0}, entities.Point{X: 10, Y: 0}),
		entities.NewLoad(3, entities.Point{X: 2, Y: 0}, entities.Point{X: 3, Y: 3}),
	}

	dispatch := NewNearestLoadDispatch(loads)
	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 1 {
		t.Errorf("Expected 1 driver, got %d", len(drivers))
	}

	loadTwo := drivers[0].GetPath()[1]

	if loadTwo.LoadNumber != 3 {
		t.Errorf("The second load to be picked up should be 3, got %d", loadTwo.LoadNumber)
	}
}
