package entities

import (
	"testing"
)

func TestGenerateChildrenShouldAddLoadsToCorrectQuad(t *testing.T) {
	loadOne := NewLoad(1, Point{1, 1}, Point{-2, 2})
	outOfBoundsLoad := NewLoad(2, Point{360.01, 2}, Point{3, 2})
	loadTwo := NewLoad(2, Point{-2, 2}, Point{3, 2})
	loadThree := NewLoad(3, Point{-4, -3}, Point{-5, -2})
	loadFour := NewLoad(4, Point{7, -3}, Point{5, -2})

	loads := []*Load{&loadOne, &loadTwo, &loadThree, &loadFour, &outOfBoundsLoad}

	quadMap := NewQuadMap(loads)

	quadMap.generateChildren()

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

func TestOriginPlacement(t *testing.T) {
	quadMap := NewQuadMap([]*Load{})

	if quadMap.originX != 0 {
		t.Error("Expected originX to be 0, got ", quadMap.originX)
	}

	if quadMap.originY != 0 {
		t.Error("Expected originY to be 0, got ", quadMap.originY)
	}

	quadZero := quadMap.GetChildren()[0]

	if quadZero.originX != 180 {
		t.Error("Expected quadZero originX to be 180, got ", quadZero.originX)
	}

	if quadZero.originY != 180 {
		t.Error("Expected quadZero originY to be 180, got ", quadZero.originY)
	}

	quadZeroZero := quadZero.GetChildren()[0]
	if quadZeroZero.originX != 270 {
		t.Error("Expected quadZeroZero originX to be 270, got ", quadZeroZero.originX)
	}

	if quadZeroZero.originY != 270 {
		t.Error("Expected quadZeroZero originY to be 270, got ", quadZeroZero.originY)
	}

	quadZeroOne := quadZero.GetChildren()[1]
	if quadZeroOne.originX != 90 {
		t.Error("Expected quadZeroOne originX to be 90, got ", quadZeroOne.originX)
	}

	if quadZeroOne.originY != 270 {
		t.Error("Expected quadZeroOne originY to be 270, got ", quadZeroOne.originY)
	}

	quadZeroTwo := quadZero.GetChildren()[2]
	if quadZeroTwo.originX != 90 {
		t.Error("Expected quadZeroTwo originX to be 90, got ", quadZeroTwo.originX)
	}

	if quadZeroTwo.originY != 90 {
		t.Error("Expected quadZeroTwo originY to be 90, got ", quadZeroTwo.originY)
	}

	quadZeroThree := quadZero.GetChildren()[3]
	if quadZeroThree.originX != 270 {
		t.Error("Expected quadZeroThree originX to be 270, got ", quadZeroThree.originX)
	}

	if quadZeroThree.originY != 90 {
		t.Error("Expected quadZeroThree originY to be 90, got ", quadZeroThree.originY)
	}
}

func TestZoomLevelsWork(t *testing.T) {
	loadOne := NewLoad(1, Point{0, 0}, Point{-2, 2})
	loadTwo := NewLoad(2, Point{180, 120}, Point{3, 2})
	loadThree := NewLoad(3, Point{181, 181}, Point{3, 2})
	loadFour := NewLoad(4, Point{-1, 2}, Point{3, 2})

	loads := []*Load{&loadOne, &loadTwo, &loadThree, &loadFour}

	quadMap := NewQuadMap(loads)
	if len(quadMap.loads) != 4 {
		t.Error("Expected quadMap to have 4 load, got ", len(quadMap.loads))
	}

	quadZero := quadMap.GetChildren()[0]
	if len(quadZero.loads) != 3 {
		t.Error("Expected quadZeroZero to have 3 load, got ", len(quadZero.loads))
	}

	quadZeroTwo := quadZero.GetChildren()[2]

	if len(quadZeroTwo.loads) != 2 {
		t.Error("Expected quadZeroTwo to have 2 load, got ", len(quadZeroTwo.loads))
	}

	quadZeroZero := quadZeroTwo.GetChildren()[0]

	if len(quadZeroZero.loads) != 1 {
		t.Error("Expected quadZeroZero to have 1 load, got ", len(quadZeroZero.loads))
	}
}
