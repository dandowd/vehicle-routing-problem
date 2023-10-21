package dispatchers

import "vehicle-routing-problem/entities"

func GetTotalCost(drivers []*entities.Driver) float64 {
	totalCost := 0.0
	for _, driver := range drivers {
		totalCost += driver.GetTotalTime()
	}

	totalCost += 500 * float64(len(drivers))

	return totalCost
}

func CombineDriverLoads(drivers []*entities.Driver) []*entities.Load {
	loads := []*entities.Load{}
	for _, driver := range drivers {
		loads = append(loads, driver.GetPath()...)
	}

	return loads
}
