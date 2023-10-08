package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestSearchForRoutesShouldReturnPathWithLeastTravelCosts(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 5, Y: 0}, entities.Point{X: 12, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 6, Y: 0}, entities.Point{X: 2, Y: 0}),
		entities.NewLoad(3, entities.Point{X: 2, Y: 0}, entities.Point{X: 3, Y: 0}),
	}

	dispatch := NewNearestLoadDFSDispatch(loads, 4)

	drivers := dispatch.SearchForRoutes()

	path := drivers[0].GetCompletedLoads()

	if path[0].LoadNumber != 2 {
		t.Errorf("Expected first load to be 2, got %d", path[0].LoadNumber)
	}

	if path[1].LoadNumber != 3 {
		t.Errorf("Expected second load to be 3, got %d", path[1].LoadNumber)
	}

	if path[2].LoadNumber != 1 {
		t.Errorf("Expected third load to be 1, got %d", path[2].LoadNumber)
	}
}

func TestSearchForRoutesShouldSpawnCorrectNumberOfDrivers(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 100, Y: 0}, entities.Point{X: 150, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 200, Y: 0}, entities.Point{X: 300, Y: 0}),
		entities.NewLoad(3, entities.Point{X: -200, Y: 0}, entities.Point{X: 0, Y: 0}),
	}

	dispatch := NewNearestLoadDFSDispatch(loads, 4)

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 2 {
		t.Errorf("Expected 2 drivers, got %d", len(drivers))
	}
}

func TestSearchShouldReturnBestPath(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 1, Y: 0}, entities.Point{X: 2, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 2, Y: 0}, entities.Point{X: 3, Y: 0}),
		entities.NewLoad(3, entities.Point{X: 3, Y: 0}, entities.Point{X: 4, Y: 0}),
		entities.NewLoad(4, entities.Point{X: 4, Y: 0}, entities.Point{X: 5, Y: 0}),
	}

	dispatch := NewNearestLoadDFSDispatch(loads, 4)

	_, driver := dispatch.search(entities.NewDriver(), 0)

	if driver.GetTotalTime() != 10 {
		t.Errorf("Expected total time of 10, got %f", driver.GetTotalTime())
	}
}

func TestSearchShouldReturnCorrectTravelCosts(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 1, Y: 0}, entities.Point{X: 1, Y: 0}),
		entities.NewLoad(2, entities.Point{X: 2, Y: 0}, entities.Point{X: 2, Y: 0}),
		entities.NewLoad(3, entities.Point{X: 3, Y: 0}, entities.Point{X: 3, Y: 0}),
		entities.NewLoad(4, entities.Point{X: 4, Y: 0}, entities.Point{X: 4, Y: 0}),
	}

	dispatch := NewNearestLoadDFSDispatch(loads, 4)

	travelCosts, _ := dispatch.search(entities.NewDriver(), 0)

	if travelCosts != 4 {
		t.Errorf("Expected travel costs of 4, got %f", travelCosts)
	}
}

func TestGetNearestLoadPickups(t *testing.T) {
	loads := []*entities.Load{
		entities.NewLoad(1, entities.Point{X: 6, Y: 6}, entities.Point{X: 2, Y: 2}),
		entities.NewLoad(2, entities.Point{X: 2, Y: 2}, entities.Point{X: 3, Y: 3}),
		entities.NewLoad(3, entities.Point{X: 3, Y: 3}, entities.Point{X: 4, Y: 4}),
		entities.NewLoad(4, entities.Point{X: 4, Y: 4}, entities.Point{X: 5, Y: 5}),
		entities.NewLoad(5, entities.Point{X: 5, Y: 5}, entities.Point{X: 6, Y: 6}),
		entities.NewLoad(6, entities.Point{X: 5, Y: 5}, entities.Point{X: 6, Y: 6}),
	}

	dispatch := NewNearestLoadDFSDispatch(loads, 4)

	nearestLoads := dispatch.getNearestLoads(entities.NewDriver())

	if len(nearestLoads) != 4 {
		t.Errorf("Expected 4 nearest loads, got %d", len(nearestLoads))
	}
}
