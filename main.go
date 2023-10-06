package main

import (
	"vehicle-routing-problem/cli"
)

func main() {
	loads := cli.ParseLoadFile("/Users/danieldowd/Downloads/problem1.txt")
	dispatcher := NewDispatch(loads)
	drivers := dispatcher.SearchForRoutes()

	cli.PrintRoutes(drivers)
}
