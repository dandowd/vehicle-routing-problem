package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestDriverUtilizationDispatcherShouldMaxDriverUse(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 0, Y: 0}, entities.Point{X: 100, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 0, Y: 0}, entities.Point{X: 100, Y: 0}),
		entities.NewLoad(3, entities.Point{X: 0, Y: 0}, entities.Point{X: 260, Y: 0}),
	}

	drivers := NewDriverUtilizationDispatch(loads).SearchForRoutes()

	if len(drivers) != 2 {
		t.Errorf("Expected 2 driver, got %d", len(drivers))
	}

	if drivers[0].GetTotalTime() != 720 {
		t.Errorf("Expected driver 1 to have 720 minutes, got %f", drivers[0].GetTotalTime())
	}
}
