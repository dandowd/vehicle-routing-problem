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
