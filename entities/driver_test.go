package entities

import "testing"

func TestDriverShouldNotBeAbleToPickup(t *testing.T) {
  driver := &Driver{currentPoint: Point{X: 360, Y: 0}}
  load := &Load{Pickup: Point{X: -360, Y: 0}, Dropoff: Point{X: 360, Y: 360}}

  if driver.CanMoveLoad(load) {
    t.Error("Driver should not be able to pickup")
  }
}
