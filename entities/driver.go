package entities

const MAX_DRIVE_TIME = 720

type Driver struct {
	currentPoint   Point
	path           []*Load
	completedLoads map[int]*Load
	totalTime      float64
}

func NewDriver() *Driver {
	return &Driver{currentPoint: Point{0, 0}, completedLoads: make(map[int]*Load), totalTime: 0}
}

func (d *Driver) MoveLoad(l *Load) {
	d.totalTime += d.currentPoint.DistanceTo(l.Pickup) + l.GetTime()
	d.currentPoint = l.Dropoff
	d.completedLoads[l.LoadNumber] = l
	d.path = append(d.path, l)
}

func (d *Driver) GetCurrentPoint() Point {
	return d.currentPoint
}

func (d *Driver) GetPath() []*Load {
	return d.path
}

func (d *Driver) GetTotalTime() float64 {
	return d.totalTime
}

func (d *Driver) CanMoveLoad(load *Load) bool {
	return d.totalTime+d.currentPoint.DistanceTo(load.Pickup)+load.GetTime()+load.Dropoff.DistanceTo(Point{0, 0}) <= MAX_DRIVE_TIME
}

func (d *Driver) DistanceTo(p Point) float64 {
	return d.currentPoint.DistanceTo(p)
}

func (d *Driver) ReturnToOrigin() {
	d.totalTime += d.currentPoint.DistanceTo(Point{0, 0})
	d.currentPoint = Point{0, 0}
}

func (d *Driver) HasCompletedLoad(l *Load) bool {
	return d.completedLoads[l.LoadNumber] != nil
}

func (d *Driver) MakeCopy() *Driver {
	newDriver := NewDriver()
	newDriver.currentPoint = d.currentPoint
	newDriver.completedLoads = make(map[int]*Load)
	for k, v := range d.completedLoads {
		newDriver.completedLoads[k] = v
	}
	newDriver.totalTime = d.totalTime

	// TODO: check why copy will not work
	for i := range d.path {
		newDriver.path = append(newDriver.path, d.path[i])
	}
	return newDriver
}
