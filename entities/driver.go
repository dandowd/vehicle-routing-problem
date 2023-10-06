package entities

type Driver struct {
	id             int
	currentPoint   Point
	completedLoads []Load
}

func (d *Driver) CompleteLoad(l Load) {
	d.completedLoads = append(d.completedLoads, l)
}
