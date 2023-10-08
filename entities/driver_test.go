package entities

import (
	"testing"
)

func TestDriverShouldNotBeAbleToPickup(t *testing.T) {
	driver := &Driver{currentPoint: Point{X: 360, Y: 0}}
	load := &Load{Pickup: Point{X: -360, Y: 0}, Dropoff: Point{X: 360, Y: 360}}

	if driver.CanMoveLoad(load) {
		t.Error("Driver should not be able to pickup")
	}
}

func TestMoveLoadShouldCorrectlyAddTime(t *testing.T) {
	driver := &Driver{currentPoint: Point{X: 0, Y: 0}, completedLoads: make(map[int]*Load)}
	load := &Load{Pickup: Point{X: 0, Y: 10}, Dropoff: Point{X: 0, Y: 15}}

	driver.MoveLoad(load)

	if driver.totalTime != 15 {
		t.Error("Driver time should be 15 received", driver.totalTime)
	}
}

func TestDriverRoundTripShouldAddCorrectly(t *testing.T) {
	driver := &Driver{currentPoint: Point{X: 0, Y: 0}, completedLoads: make(map[int]*Load)}
	load := &Load{Pickup: Point{X: 0, Y: 10}, Dropoff: Point{X: 0, Y: 15}}

	driver.MoveLoad(load)
	driver.ReturnToOrigin()

	if driver.totalTime != 30 {
		t.Error("Driver time should be 30 received", driver.totalTime)
	}
}

func TestCanMoveLoadShouldReturnFalseIfTheRoundTripIsTooLong(t *testing.T) {
	driver := &Driver{currentPoint: Point{X: 0, Y: 0}, completedLoads: make(map[int]*Load)}
	load := &Load{Pickup: Point{X: 0, Y: 360}, Dropoff: Point{X: 0, Y: 15}}
	loadTwo := &Load{Pickup: Point{X: 0, Y: 360}, Dropoff: Point{X: 0, Y: 15}}

	driver.MoveLoad(load)

	if driver.CanMoveLoad(loadTwo) {
		t.Error("Driver should not be able to pickup")
	}
}
