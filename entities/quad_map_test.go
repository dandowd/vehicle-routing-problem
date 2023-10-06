package entities

import (
  "testing"
)

func TestFindBounds(t *testing.T) {
  loadOne := NewLoad(1, Point{1, 1}, Point{2, 2})
  loadTwo := NewLoad(2, Point{-2, -2}, Point{3, 2})

  loads := []*Load{&loadOne, &loadTwo}
  
  minX, maxX, minY, maxY := findBounds(loads)

  if minX != -2 {
    t.Error("Expected minX to be -2, got ", minX)
  }

  if maxX != 3 {
    t.Error("Expected maxX to be 3, got ", maxX)
  }

  if minY != -2 {
    t.Error("Expected minY to be -2, got ", minY)
  }

  if maxY != 2 {
    t.Error("Expected maxY to be 2, got ", maxY)
  }
}

func TestGenerateChildrenShouldAddLoadsToCorrectQuad(t *testing.T) {
  loadOne := NewLoad(1, Point{-1, 1}, Point{-2, 2})
  loadTwo := NewLoad(2, Point{2, 2}, Point{3, 2})
  loadThree := NewLoad(3, Point{-4, -3}, Point{-5, -2})
  loadFour := NewLoad(4, Point{7, -3}, Point{5, -2})

  loads := []*Load{&loadOne, &loadTwo, &loadThree, &loadFour}
  
  minX, maxX, minY, maxY := findBounds(loads)
  quadMap := QuadMap{loads: loads, minX: minX, maxX: maxX, minY: minY, maxY: maxY}

  quadMap.GenerateChildren()

  if len(quadMap.subMaps[0].loads) != 1 {
    t.Error("Expected subMap[0] to have 1 load, got ", len(quadMap.subMaps[0].loads))
  }

  if len(quadMap.subMaps[1].loads) != 1 {
    t.Error("Expected subMap[1] to have 1 load, got ", len(quadMap.subMaps[1].loads))
  }

  if len(quadMap.subMaps[2].loads) != 1 {
    t.Error("Expected subMap[2] to have 1 load, got ", len(quadMap.subMaps[2].loads))
  }

  if len(quadMap.subMaps[3].loads) != 1 {
    t.Error("Expected subMap[3] to have 1 load, got ", len(quadMap.subMaps[3].loads))
  }
}
