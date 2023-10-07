package cli

import (
	"fmt"
	"vehicle-routing-problem/entities"
)

func PrintRoutes(drivers []*entities.Driver) {
	for _, driver := range drivers {
		fmt.Println(driver.GetCompletedLoads())
	}
}

func FormatDrivers(drivers []*entities.Driver) {
	for _, driver := range drivers {
    driverLoads := "[" 
		for _, load := range driver.GetCompletedLoads() {
      driverLoads += fmt.Sprintf("%v,", load.LoadNumber)
		}

    driverLoads = driverLoads[:len(driverLoads)-1]
    driverLoads += "]"

		fmt.Println(driverLoads)
	}
}
