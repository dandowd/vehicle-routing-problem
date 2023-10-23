package main

import (
	"fmt"
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
				Value:   0,
				Usage:   "Number of drivers per route file",
			},
			&cli.BoolFlag{
				Name:    "fast",
				Aliases: []string{"f"},
				Usage:   "Run closest load algorithm only",
			},
			&cli.BoolFlag{
				Name:    "print-costs",
				Aliases: []string{"c"},
				Usage:   "Print costs of each algorithm",
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
				annealingDrivers := dispatchers.Annealing(loads, 20000, 1000, 0.97, 100)

				if c.Bool("print-costs") {
					fmt.Println("Annealing Costs:", dispatchers.GetTotalCost(annealingDrivers))
				}

				if dispatchers.GetTotalCost(annealingDrivers) < dispatchers.GetTotalCost(bestDrivers) {
					bestDrivers = annealingDrivers
				}
			}

			routeFile := c.String("driver-route-file")
			if routeFile != "" {
				if c.Int("driver-count") == 0 {
					visualization.Route(bestDrivers, "Route", routeFile)
				} else {
					visualization.SplitRoutes(bestDrivers, routeFile, c.Int("driver-count"))
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
