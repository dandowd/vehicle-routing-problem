package visualization

import (
	"vehicle-routing-problem/entities"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Route(drivers []*entities.Driver) {
  p := plot.New()

  points := plotter.XYs{}
  
  for driverIndex, driver := range drivers {
    lastPoint := plotter.XY{X: 0, Y: 0}
    for _, load := range driver.GetPath() {
      pickup := plotter.XY{X: load.Pickup.X, Y: load.Pickup.Y}
      dropoff := plotter.XY{X: load.Dropoff.X, Y: load.Dropoff.Y}
    
      line, err := plotter.NewLine(plotter.XYs{lastPoint, pickup, dropoff})
      if err != nil {
        panic(err)
      }

      line.Color = plotutil.Color(driverIndex)
      p.Add(line)

      lastPoint = dropoff
    }
  }
  
  scatter, err := plotter.NewScatter(points)
  if err != nil {
    panic(err)
  }

  p.Add(scatter)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "route.png"); err != nil {
		panic("Could not save the plot to a PNG file")
	}
}
