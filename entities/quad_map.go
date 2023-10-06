package entities

import "math"

const MAX_DRIVE_TIME = 720
const MAX_DISTANCE = MAX_DRIVE_TIME / 2

type QuadMap struct {
  subMaps [4]*QuadMap
  loads []*Load
  minX, maxX, minY, maxY float64
}

func (m* QuadMap) HasPoint(p Point) bool {
  return p.X >= m.minX && p.X <= m.maxX && p.Y >= m.minY && p.Y <= m.maxY
}

func (m* QuadMap) AddLoad(l* Load) {
  m.loads = append(m.loads, l)
}

// Generates children and adds loads based off pickup location
// This map has blind spots based off the smallest non-zero float64
// TODO: Add dropoff location for better heuristics
func (m* QuadMap) GenerateChildren() *QuadMap {
  zeroMap := &QuadMap{[4]*QuadMap{}, nil, -MAX_DISTANCE, 0, 0, MAX_DISTANCE}
  oneMap := &QuadMap{[4]*QuadMap{}, nil, math.SmallestNonzeroFloat64, MAX_DISTANCE, 0, MAX_DISTANCE}
  twoMap := &QuadMap{[4]*QuadMap{}, nil, math.SmallestNonzeroFloat64, MAX_DISTANCE, -MAX_DISTANCE, -math.SmallestNonzeroFloat64}
  threeMap := &QuadMap{[4]*QuadMap{}, nil, -MAX_DISTANCE, 0, -MAX_DISTANCE, -math.SmallestNonzeroFloat64}

  m.subMaps[0] = zeroMap
  m.subMaps[1] = oneMap
  m.subMaps[2] = twoMap
  m.subMaps[3] = threeMap

  for _, l := range m.loads {
    if zeroMap.HasPoint(l.Pickup) {
      zeroMap.AddLoad(l)
    } else if oneMap.HasPoint(l.Pickup) {
      oneMap.AddLoad(l)
    } else if twoMap.HasPoint(l.Pickup) {
      twoMap.AddLoad(l)
    } else if threeMap.HasPoint(l.Pickup) {
      threeMap.AddLoad(l)
    }
  }
  return m
}

func NewQuadMap(loads []*Load, depth int) *QuadMap {
  minX, maxX, minY, maxY := findBounds(loads)
  quadMap := QuadMap{[4]*QuadMap{}, loads, minX, maxX, minY, maxY}
  quadMap.GenerateChildren()

  for i := 0; i < 4; i++ {

  }
  return &quadMap
}

// Finds bounds of loads using pickup and dropoff locations
func findBounds(loads []*Load) (float64, float64, float64, float64) {
  var minX, maxX, minY, maxY float64
  for _, l := range loads {
    if l.Pickup.X < minX {
      minX = l.Pickup.X
    }
    if l.Pickup.X > maxX {
      maxX = l.Pickup.X
    }
    if l.Pickup.Y < minY {
      minY = l.Pickup.Y
    }
    if l.Pickup.Y > maxY {
      maxY = l.Pickup.Y
    }

    if l.Dropoff.X < minX {
      minX = l.Dropoff.X
    }
    if l.Dropoff.X > maxX {
      maxX = l.Dropoff.X
    }
    if l.Dropoff.Y < minY {
      minY = l.Dropoff.Y
    }
    if l.Dropoff.Y > maxY {
      maxY = l.Dropoff.Y
    }
  }
  return minX, maxX, minY, maxY
}
