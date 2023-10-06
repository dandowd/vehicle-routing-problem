package entities

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
// TODO: Add dropoff location for better heuristics
func (m* QuadMap) GenerateChildren() *QuadMap {
  zeroMap := &QuadMap{[4]*QuadMap{}, nil, m.minX, m.maxX / 2, m.minY, m.maxY / 2}
  oneMap := &QuadMap{[4]*QuadMap{}, nil, m.maxX / 2, m.maxX, m.minY, m.maxY / 2}
  twoMap := &QuadMap{[4]*QuadMap{}, nil, m.minX, m.maxX / 2, m.maxY / 2, m.maxY}
  threeMap := &QuadMap{[4]*QuadMap{}, nil, m.maxX / 2, m.maxX, m.maxY / 2, m.maxY}

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

// Finds bounds of loads using pickup and dropoff locations
func findBounds(loads* []*Load) (float64, float64, float64, float64) {
  var minX, maxX, minY, maxY float64
  for _, l := range *loads {
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
