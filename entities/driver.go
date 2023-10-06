package entities

type Driver struct {
	id             int
	currentPoint   Point
	completedLoads []Load
	totalTime		float64	
}

func (d *Driver) AddLoad(l *Load) {
	d.completedLoads = append(d.completedLoads, *l)
	d.currentPoint = l.Dropoff
	d.totalTime += l.GetTime()
}

func (d *Driver) CompleteLoad(l Load) {
	d.completedLoads = append(d.completedLoads, l)
}

func (d *Driver) GetCompletedLoads() []Load {
	return d.completedLoads
}
