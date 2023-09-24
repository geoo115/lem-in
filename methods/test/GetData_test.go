package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestGetData(t *testing.T) {
	data := &lem.Input{}
	dataFile := []string{
		"4", // Valid number of ants
		"##start",
		"0 0 3", // Valid room entry
		"2 2 5", // Valid room entry
		"3 4 0", // Valid room entry
		"##end",
		"1 8 3", // Valid room entry
		"0-2",   // Valid room entry
		"2-3",   // Valid room entry
		"3-1",   // Valid room entry
	}

	err := lem.GetData(data, dataFile)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Validate extracted data
	if data.Ants != 4 {
		t.Errorf("Expected Ants to be 10, but got: %d", data.Ants)
	}
	if data.StartR != "0" {
		t.Errorf("Expected StartR to be 'RoomA', but got: %s", data.StartR)
	}
	if data.EndR != "1" {
		t.Errorf("Expected EndR to be 'RoomE', but got: %s", data.EndR)
	}
	if len(data.Rooms) != 4 {
		t.Errorf("Expected 8 rooms, but got: %d", len(data.Rooms))
	}
	if len(data.Coords) != 4 {
		t.Errorf("Expected 8 coordinates, but got: %d", len(data.Coords))
	}
	if len(data.Links) != 3 {
		t.Errorf("Expected 3 links, but got: %d", len(data.Links))
	}
}
