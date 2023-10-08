package entities

import "testing"

func TestPointDistanceTo(t *testing.T) {
	p1 := Point{X: 0, Y: 10}
	p2 := Point{X: 0, Y: 15}

	dist := p1.DistanceTo(p2)

	if dist != 5 {
		t.Error("Distance should be 5 received", dist)
	}
}

func TestPointToDistanceFloating(t *testing.T) {
	p1 := Point{X: 0, Y: 10}
	p2 := Point{X: 0, Y: 15.5}

	dist := p1.DistanceTo(p2)

	if dist != 5.5 {
		t.Error("Distance should be 5.5 received", dist)
	}
}
