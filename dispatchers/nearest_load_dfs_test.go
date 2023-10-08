package dispatchers

import (
	"fmt"
	"testing"
	"vehicle-routing-problem/entities"
)

func TestSearchShouldReturnBestPath(t *testing.T) {
  loads := []*entities.Load{
    entities.NewLoad(1, entities.Point{X: 1, Y: 1}, entities.Point{X: 2, Y: 2}),
    entities.NewLoad(2, entities.Point{X: 2, Y: 2}, entities.Point{X: 3, Y: 3}),
    entities.NewLoad(3, entities.Point{X: 3, Y: 3}, entities.Point{X: 4, Y: 4}),
    entities.NewLoad(4, entities.Point{X: 4, Y: 4}, entities.Point{X: 5, Y: 5}),
  }

  dispatch := NewNearestLoadDFSDispatch(loads, 4)

  cost, paths := dispatch.search(entities.Point{X: 0, Y: 0}, 0, 0, make(map[int]*entities.Load))

  fmt.Println(cost, paths)
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

  nearestLoads := dispatch.getNearestLoads(entities.Point{X: 0, Y: 0}, make(map[int]*entities.Load))

  if len(nearestLoads) != 4 {
    t.Errorf("Expected 4 nearest loads, got %d", len(nearestLoads))
  }
}
