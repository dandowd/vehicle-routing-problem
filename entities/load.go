package entities

type Load struct {
	LoadNumber int
	Pickup     Point
	Dropoff    Point
}

func (l *Load) GetTime() float64 {
	return l.Pickup.DistanceTo(l.Dropoff)
}

func NewLoad(loadNumber int, pickup Point, dropoff Point) *Load {
	return &Load{loadNumber, pickup, dropoff}
}
