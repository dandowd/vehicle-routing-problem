package entities

import "math"

type Load struct {
	LoadNumber int
	Pickup     Point
	Dropoff    Point
}

func (l *Load) GetTime() float64 {
	return math.Sqrt(math.Pow(l.Pickup.X-l.Dropoff.X, 2) + math.Pow(l.Pickup.Y-l.Dropoff.Y, 2))
}

func NewLoad(loadNumber int, pickup Point, dropoff Point) Load {
	return Load{loadNumber, pickup, dropoff}
}
