package visualization

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type GraphLog struct {
  points plotter.XYs
}

func NewGraphLog() *GraphLog {
  return &GraphLog{points: make(plotter.XYs, 0)}
}

func (g *GraphLog) AddPoint(x, y float64) {
  g.points = append(g.points, plotter.XY{X: x, Y: y})
}

func (g *GraphLog) CreateFile(filename string) {
  p := plot.New()

  scatter, err := plotter.NewScatter(g.points)
  if err != nil {
    panic(err)
  }

  p.Add(scatter)

	if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprint(filename, ".png")); err != nil {
		fmt.Printf("Could not save the plot to a PNG file: %v\n", err)
		panic(err)
	}
}
