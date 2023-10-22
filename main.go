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
		},
		Action: func(c *cli.Context) error {
			filepath := c.Args().First()
			loads := utils.ParseLoadFile(filepath)

			nearestLoadDrivers := dispatchers.NewNearestLoadDispatch(loads).SearchForRoutes()
			fmt.Println("Nearest Load Total Cost:", dispatchers.GetTotalCost(nearestLoadDrivers))

			drivers := dispatchers.Annealing(loads)
			fmt.Println("Annealing Total Cost:", dispatchers.GetTotalCost(drivers))

			routeFile := c.String("route")
			if routeFile != "" {
				title := fmt.Sprintf("Total Cost: %f", dispatchers.GetTotalCost(drivers))
				visualization.Route(drivers, title, routeFile)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
