package visualization

import (
	"fmt"
	"math"
	"vehicle-routing-problem/entities"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func SplitRoutes(drivers []*entities.Driver, filename string, driverCount int) {
	numberOfFiles := math.Ceil(float64(len(drivers) / driverCount))
	for i := 0; i < int(numberOfFiles); i++ {
		driverChunk := drivers[i*driverCount : (i+1)*driverCount]
		title := fmt.Sprintf("Route %d", i)
		filepath := fmt.Sprintf("%s_%d", filename, i)

		Route(driverChunk, title, filepath)
	}
}

func Route(drivers []*entities.Driver, title string, filename string) {
	p := plot.New()
	p.Title.Text = title

	pickupPoints := plotter.XYs{}
	dropoffPoints := plotter.XYs{}

	for driverIndex, driver := range drivers {
		lastPoint := plotter.XY{X: 0, Y: 0}

		for _, load := range driver.GetPath() {
			pickup := plotter.XY{X: load.Pickup.X, Y: load.Pickup.Y}
			pickupPoints = append(pickupPoints, pickup)

			dropoff := plotter.XY{X: load.Dropoff.X, Y: load.Dropoff.Y}
			dropoffPoints = append(dropoffPoints, dropoff)

			travelLine := newTravelLine(lastPoint, pickup, driverIndex)
			p.Add(travelLine)

			lineLine := newLoadLine(pickup, dropoff, driverIndex)
			p.Add(lineLine)

			lastPoint = dropoff
		}

		line := newTravelLine(lastPoint, plotter.XY{X: 0, Y: 0}, driverIndex)
		p.Add(line)

		p.Legend.Add(fmt.Sprintf("Driver number: %d, utilization: %f", driverIndex, driver.GetTotalTime()), line)
	}

	createPickupScatter(pickupPoints, p)

	createDropoffScatter(dropoffPoints, p)

	createOriginScatter(p)

	if err := p.Save(20*vg.Inch, 20*vg.Inch, fmt.Sprint(filename, ".png")); err != nil {
		fmt.Printf("Could not save the plot to a PNG file: %v\n", err)
		panic(err)
	}
}

func createPickupScatter(pickupPoints plotter.XYs, p *plot.Plot) *plotter.Scatter {
	scatter, err := plotter.NewScatter(pickupPoints)
	if err != nil {
		panic(err)
	}

	scatter.GlyphStyle.Shape = draw.BoxGlyph{}
	scatter.GlyphStyle.Radius = vg.Points(5)

	p.Add(scatter)
	p.Legend.Add("Pickup", scatter)

	return scatter
}

func createOriginScatter(p *plot.Plot) *plotter.Scatter {
	origin := plotter.XY{X: 0, Y: 0}

	scatter, err := plotter.NewScatter(plotter.XYs{origin})
	if err != nil {
		panic(err)
	}

	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	scatter.GlyphStyle.Radius = vg.Points(8)

	p.Add(scatter)
	p.Legend.Add("Origin", scatter)

	return scatter
}

func createDropoffScatter(dropoffPoints plotter.XYs, p *plot.Plot) *plotter.Scatter {
	scatter, err := plotter.NewScatter(dropoffPoints)
	if err != nil {
		panic(err)
	}

	scatter.GlyphStyle.Shape = draw.SquareGlyph{}
	scatter.GlyphStyle.Radius = vg.Points(5)

	p.Add(scatter)
	p.Legend.Add("Dropoff", scatter)

	return scatter
}

func newLoadLine(pickup plotter.XY, dropoff plotter.XY, driverIndex int) *plotter.Line {
	line, err := plotter.NewLine(plotter.XYs{pickup, dropoff})
	if err != nil {
		panic(err)
	}

	line.Color = plotutil.Color(driverIndex)

	return line
}

func newTravelLine(lastPoint plotter.XY, pickup plotter.XY, driverIndex int) *plotter.Line {
	travelLine, err := plotter.NewLine(plotter.XYs{lastPoint, pickup})
	travelLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	if err != nil {
		panic(err)
	}

	travelLine.Color = plotutil.Color(driverIndex)
	travelLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	return travelLine
}
