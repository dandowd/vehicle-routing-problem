package cli

import "testing"

func TestParseLineShouldReturnValidLoad(t *testing.T) {
	load := parseLine("1 (1,1) (0,0)")

	if load.LoadNumber != 1 {
		t.Error("Expected load number to be 1, got ", load.LoadNumber)
	}

	if load.Pickup.X != 1 {
		t.Error("Expected pickup X to be 1, got ", load.Pickup.X)
	}

	if load.Pickup.Y != 1 {
		t.Error("Expected pickup Y to be 1, got ", load.Pickup.Y)
	}

	if load.Dropoff.X != 0 {
		t.Error("Expected dropoff X to be 0, got ", load.Dropoff.X)
	}

	if load.Dropoff.Y != 0 {
		t.Error("Expected dropoff Y to be 0, got ", load.Dropoff.Y)
	}
}

func TestParseLineShouldConvertFloatsProperly(t *testing.T) {
	load := parseLine("1 (-264.91649183111383,-55.303821595572096) (-121.15241691952006,52.59477077550236)")

	if load.Pickup.X != -264.91649183111383 {
		t.Error("Expected pickup X to be -264.91649183111383, got ", load.Pickup.X)
	}

	if load.Pickup.Y != -55.303821595572096 {
		t.Error("Expected pickup Y to be -55.303821595572096, got ", load.Pickup.Y)
	}

	if load.Dropoff.X != -121.15241691952006 {
		t.Error("Expected dropoff X to be -121.15241691952006, got ", load.Dropoff.X)
	}

	if load.Dropoff.Y != 52.59477077550236 {
		t.Error("Expected dropoff Y to be 52.59477077550236, got ", load.Dropoff.Y)
	}
}
