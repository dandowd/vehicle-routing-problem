package dispatchers

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestSearchShouldReturnBestPath(t *testing.T) {
  loads := []*entities.Load{
    entities.NewLoad(1, entities.Point{X: 1, Y: 0}, entities.Point{X: 1, Y: 0}),
    entities.NewLoad(2, entities.Point{X: 2, Y: 0}, entities.Point{X: 2, Y: 0}),
    entities.NewLoad(3, entities.Point{X: 3, Y: 0}, entities.Point{X: 3, Y: 0}),
    entities.NewLoad(4, entities.Point{X: 4, Y: 0}, entities.Point{X: 4, Y: 0}),
  }

  dispatch := NewNearestLoadDFSDispatch(loads, 4)

  _, driver := dispatch.search(entities.NewDriver(), 0)

  if driver.GetTotalTime() != 4 {
    t.Errorf("Expected total time of 4, got %f", driver.GetTotalTime())
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
