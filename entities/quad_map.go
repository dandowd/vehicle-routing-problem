package entities

const MAX_DISTANCE = MAX_DRIVE_TIME / 2

type QuadMap struct {
	parent           *QuadMap
	zoomLevel        int
	subMaps          [4]*QuadMap
	loads            []*Load
	originX, originY float64
}

func (m *QuadMap) HasPoint(p Point) bool {
	var maxDistance = float64(MAX_DISTANCE / m.zoomLevel)
	return p.X >= m.originX-maxDistance && p.X <= m.originX+maxDistance && p.Y >= m.originY-maxDistance && p.Y <= m.originY+maxDistance
}

func (m *QuadMap) addLoad(l *Load) {
	m.loads = append(m.loads, l)
}

// Generates children and adds loads based off pickup location
// TODO: Add dropoff location for better heuristics
func (m *QuadMap) generateChildren() *QuadMap {
	var originOffset = float64(MAX_DISTANCE/m.zoomLevel) / 2

	zeroMap := &QuadMap{m, m.zoomLevel + 1, [4]*QuadMap{}, nil, m.originX + originOffset, m.originY + originOffset}
	oneMap := &QuadMap{m, m.zoomLevel + 1, [4]*QuadMap{}, nil, m.originX - originOffset, m.originY + originOffset}
	twoMap := &QuadMap{m, m.zoomLevel + 1, [4]*QuadMap{}, nil, m.originX - originOffset, m.originY - originOffset}
	threeMap := &QuadMap{m, m.zoomLevel + 1, [4]*QuadMap{}, nil, m.originX + originOffset, m.originY - originOffset}

	m.subMaps[0] = zeroMap
	m.subMaps[1] = oneMap
	m.subMaps[2] = twoMap
	m.subMaps[3] = threeMap

	for _, l := range m.loads {
		if zeroMap.HasPoint(l.Pickup) {
			zeroMap.addLoad(l)
		} else if oneMap.HasPoint(l.Pickup) {
			oneMap.addLoad(l)
		} else if twoMap.HasPoint(l.Pickup) {
			twoMap.addLoad(l)
		} else if threeMap.HasPoint(l.Pickup) {
			threeMap.addLoad(l)
		}
	}
	return m
}

func (m *QuadMap) GetChildren() [4]*QuadMap {
	if m.subMaps[0] == nil {
		m.generateChildren()
		return m.subMaps
	}

	return m.subMaps
}

func (m* QuadMap) GetLoads() []*Load {
	return m.loads
}

func NewQuadMap(loads []*Load) *QuadMap {
	quadMap := QuadMap{nil, 1, [4]*QuadMap{}, loads, 0, 0}

	return &quadMap
}
