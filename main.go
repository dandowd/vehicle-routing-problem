package main

import (
	"fmt"
	"math"
	"os"
	"vehicle-routing-problem/dispatchers"
	"vehicle-routing-problem/utils"
	"vehicle-routing-problem/visualization"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "driver-route-file",
				Aliases: []string{"r"},
				Usage:   "Drivers routes file path",
			},
			&cli.IntFlag{
				Name:    "driver-count",
				Aliases: []string{"dc"},
				Value: 1,
				Usage:   "Number of drivers per route file",
			},
			&cli.StringFlag{
				Name:    "annealing-graph",
				Aliases: []string{"ag"},
				Usage:   "Annealing graph file path",
			},
			&cli.BoolFlag{
				Name:    "fast",
				Aliases: []string{"f"},
				Usage:   "Run closest load algorithm only",
			},
			&cli.BoolFlag{
				Name:    "print-costs",
				Aliases: []string{"c"},
				Usage:   "Run closest load algorithm only",
			},
		},
		Action: func(c *cli.Context) error {
			filepath := c.Args().First()
			loads := utils.ParseLoadFile(filepath)

			bestDrivers := dispatchers.NewNearestLoadDispatch(loads).SearchForRoutes()
			if c.Bool("print-costs") {
				fmt.Println("Nearest Load Costs:", dispatchers.GetTotalCost(bestDrivers))
			}

			if !c.Bool("fast") {
				annealingDrivers := dispatchers.Annealing(dispatchers.CombineDriverLoads(bestDrivers))
				for i := 0; i < 5; i++ {
					newAnnealingDrivers := dispatchers.Annealing(dispatchers.CombineDriverLoads(annealingDrivers))

					if dispatchers.GetTotalCost(newAnnealingDrivers) < dispatchers.GetTotalCost(annealingDrivers) {
						annealingDrivers = newAnnealingDrivers
					}
				}

				if c.Bool("print-costs") {
					fmt.Println("Annealing Costs:", dispatchers.GetTotalCost(annealingDrivers))
				}

				if dispatchers.GetTotalCost(annealingDrivers) < dispatchers.GetTotalCost(bestDrivers) {
					bestDrivers = annealingDrivers
				}
			}

			routeFile := c.String("driver-route-file")
			if routeFile != "" {
				numberOfFiles := math.Ceil(float64(len(bestDrivers) / c.Int("driver-count")))
				for i := 0; i < int(numberOfFiles); i++ {
					driverChunk := bestDrivers[i*c.Int("driver-count") : (i+1)*c.Int("driver-count")]
					title := fmt.Sprintf("Total Cost: %f", dispatchers.GetTotalCost(driverChunk))
					filepath := fmt.Sprintf("%s_%d", routeFile, i)
					visualization.Route(driverChunk, title, filepath)
				}
			}

			utils.PrintRoutes(bestDrivers)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
