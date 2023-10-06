package main

import (
  "testing"
)

func TestGetPickupsInArea(t *testing.T) {
  loadOne := NewLoad(1, Point{1, 1}, Point{2, 2})
  loadTwo := NewLoad(2, Point{-2, -2}, Point{3, 2})

  loadMap := LoadMap{[]Load{loadOne, loadTwo}}

  pickups := loadMap.GetPickupsInArea(Point{0, 0}, 1)

  t.Log(pickups)
}
