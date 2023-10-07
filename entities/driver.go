package entities

import "fmt"

const MAX_DRIVE_TIME = 720

type Driver struct {
	currentPoint   Point
	completedLoads []*Load
	totalTime      float64
}

func CreateNewDriver() *Driver {
	return &Driver{currentPoint: Point{0, 0}, completedLoads: []*Load{}, totalTime: 0}
}

func (d *Driver) MoveLoad(l *Load) {
	d.totalTime += d.currentPoint.DistanceTo(l.Pickup)+l.GetTime()
	d.currentPoint = l.Dropoff
	d.completedLoads = append(d.completedLoads, l)

	if d.currentPoint.DistanceTo(Point{0, 0}) > MAX_DRIVE_TIME {
		fmt.Println("Driver has exceeded max drive time")
	}
}

func (d *Driver) GetCompletedLoads() []*Load {
	return d.completedLoads
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
