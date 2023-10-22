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
				annealingDrivers := dispatchers.Annealing(loads)
				if c.Bool("print-costs") {
					fmt.Println("Annealing Costs:", dispatchers.GetTotalCost(annealingDrivers))
				}

				if dispatchers.GetTotalCost(annealingDrivers) < dispatchers.GetTotalCost(bestDrivers) {
					bestDrivers = annealingDrivers
				}
			}

			routeFile := c.String("route")
			if routeFile != "" {
				title := fmt.Sprintf("Total Cost: %f", dispatchers.GetTotalCost(bestDrivers))
				visualization.Route(bestDrivers, title, routeFile)
			}

			utils.PrintRoutes(bestDrivers)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
