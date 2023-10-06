package main

type LoadMap struct {
	loads []Load
}

func (m* LoadMap) GetPickupsInArea(p Point, radius float64) []Load {
	var loadsInArea []Load

	for _, l := range m.loads {
		if l.Pickup.X < p.X + radius && l.Pickup.X > p.X - radius && l.Pickup.Y < p.Y + radius && l.Pickup.Y > p.Y - radius {
			loadsInArea = append(loadsInArea, l)
		}
	}
	return loadsInArea
}
