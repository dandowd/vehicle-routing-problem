package main

import (
	"fmt"
	"os"
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
)

var nearestLoadLogger = cli.NewFileLogger("./nearest_load.log")

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program takes exactly one argument, the path to the problem file.")
		os.Exit(1)
	}

	filepath := os.Args[1]
	loads := cli.ParseLoadFile(filepath)

	drivers := dispatchers.NewNearestLoadDispatch(loads).SearchForRoutes()

	nearestLoadLogger.Println("Total cost:", dispatchers.GetTotalCost(drivers))
	nearestLoadLogger.Println(cli.FormatPath(drivers))
}
