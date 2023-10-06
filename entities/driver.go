package entities

const MAX_DRIVE_TIME = 720

type Driver struct {
	id             int
	currentPoint   Point
	completedLoads []Load
	totalTime      float64
}

func (d *Driver) AddLoad(l *Load) {
	d.completedLoads = append(d.completedLoads, *l)
	d.currentPoint = l.Dropoff
	d.totalTime += l.GetTime()
}

func (d *Driver) GetCompletedLoads() []Load {
	return d.completedLoads
}

func (d *Driver) CanPickup(load *Load) bool {
	return d.currentPoint.DistanceTo(load.Pickup)+load.GetTime()+load.Dropoff.DistanceTo(Point{0, 0}) <= MAX_DRIVE_TIME
}
