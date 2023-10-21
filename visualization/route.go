package visualization

import (
	"fmt"
	"vehicle-routing-problem/entities"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Route(drivers []*entities.Driver, title string) {
  p := plot.New()
  p.Title.Text = title

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
     
    line, err := plotter.NewLine(plotter.XYs{lastPoint, plotter.XY{X: 0, Y: 0}})
    if err != nil {
      panic(err)
    }

    line.Color = plotutil.Color(driverIndex)
    p.Add(line)
    p.Legend.Add(fmt.Sprint(driverIndex), line)
  }
  
  scatter, err := plotter.NewScatter(points)
  if err != nil {
    panic(err)
  }

  scatter.GlyphStyle.Radius = vg.Points(10)
  p.Add(scatter)

  if err := p.Save(20*vg.Inch, 20*vg.Inch, fmt.Sprint("route.png")); err != nil {
		panic("Could not save the plot to a PNG file")
	}
}
