package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestIsValidCoords(t *testing.T) {
	data := &lem.Input{
		Coords: [][]string{{"1", "2"}, {"5", "6"}},
	}

	// Valid coordinates
	coords := []string{"3", "4"}
	result := lem.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for valid coordinates, but got false")
	}

	// Invalid coordinates (duplicate)
	coords = []string{"1", "1"}
	result = lem.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for duplicate coordinates, but got false")
	}

	// Invalid coordinates (negative X)
	coords = []string{"-1", "2"}
	result = lem.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for negative X coordinate, but got false")
	}

	// Invalid coordinates (negative Y)
	coords = []string{"1", "-2"}
	result = lem.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for negative Y coordinate, but got false")
	}
}
