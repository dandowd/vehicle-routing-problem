package main

import (
	"testing"
	"vehicle-routing-problem/entities"
)

func TestSearchForRoutesShouldSendOneTruck(t *testing.T) {
	loadOne := &entities.Load{LoadNumber: 1, Pickup: entities.Point{X: 1, Y: 1}, Dropoff: entities.Point{X: 0, Y: 0}}

	dispatch := NewDispatch([]*entities.Load{
		loadOne,
	})

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 1 {
		t.Error("Expected 1 driver, got ", len(drivers))
	}

	if len(drivers[0].GetCompletedLoads()) != 1 {
		t.Error("Expected 1 completed load, got ", len(drivers[0].GetCompletedLoads()))
	}
}

func TestSearchForRoutesShouldSendTwoTrucks(t *testing.T) {
	loadOne := &entities.Load{LoadNumber: 1, Pickup: entities.Point{X: 1, Y: 1}, Dropoff: entities.Point{X: -45, Y: 30}}
	loadTwo := &entities.Load{LoadNumber: 2, Pickup: entities.Point{X: 300, Y: 300}, Dropoff: entities.Point{X: 150, Y: 150}}

	dispatch := NewDispatch([]*entities.Load{
		loadOne,
		loadTwo,
	})

	drivers := dispatch.SearchForRoutes()

	if len(drivers) != 2 {
		t.Error("Expected 2 drivers, got ", len(drivers))
	}

	if len(drivers[0].GetCompletedLoads()) != 1 {
		t.Error("Expected 1 completed load, got ", len(drivers[0].GetCompletedLoads()))
	}

	if len(drivers[1].GetCompletedLoads()) != 1 {
		t.Error("Expected 1 completed load, got ", len(drivers[1].GetCompletedLoads()))
	}
}
