package main

import (
	"vehicle-routing-problem/cli"
	"vehicle-routing-problem/dispatchers"
)

func main() {
	loads := cli.ParseLoadFile("/Users/danieldowd/Downloads/problem12.txt")
	dispatcher := dispatchers.NewBruteForceDispatch(loads)
	drivers := dispatcher.SearchForRoutes()

	cli.FormatDrivers(drivers)
}
