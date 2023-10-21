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
    
      lineLine := getLoadLine(pickup, dropoff, driverIndex)
      p.Add(lineLine)

      travelLine := getTravelLine(lastPoint, pickup, driverIndex)
      p.Add(travelLine)

      lastPoint = dropoff
    }
     
    line := getTravelLine(lastPoint, plotter.XY{X: 0, Y: 0}, driverIndex)
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

func getLoadLine(pickup plotter.XY, dropoff plotter.XY, driverIndex int) *plotter.Line {
    line, err := plotter.NewLine(plotter.XYs{pickup, dropoff})
    if err != nil {
      panic(err)
    }
    
    line.Color = plotutil.Color(driverIndex)

    return line
}

func getTravelLine(lastPoint plotter.XY, pickup plotter.XY, driverIndex int) *plotter.Line {
    travelLine, err := plotter.NewLine(plotter.XYs{lastPoint, pickup})
    travelLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
    if err != nil {
      panic(err)
    }

    travelLine.Color = plotutil.Color(driverIndex)
    travelLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

    return travelLine
}
