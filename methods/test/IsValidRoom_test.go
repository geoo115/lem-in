package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestIsValidRoom(t *testing.T) {
	data := &lem.Input{
		Rooms: []string{"RoomA", "1", "2"},
		// Coords: [][]string{{"1", "2"}, {"3", "4"}},
	}

	// Valid room entry
	err := lem.IsValidRoom("RoomA 1 2", "", data)
	if err != nil {
		t.Errorf("Expected no error for valid room entry, but got: %v", err)
	}

	// Invalid room entry: name starting with 'L'
	err = lem.IsValidRoom("LRoomA 1 2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, the room can't start with L, # or space" {
		t.Errorf("Expected error for room name starting with 'L', but got: %v", err)
	}

	// Invalid room entry: name starting with '#'
	err = lem.IsValidRoom("#RoomA 1 2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, the room can't start with L, # or space" {
		t.Errorf("Expected error for room name starting with '#', but got: %v", err)
	}

	// Invalid room entry: name starting with space
	err = lem.IsValidRoom("#RoomA 1 2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, the room can't start with L, # or space" {
		t.Errorf("Expected error for room name starting with space, but got: %v", err)
	}

	// Invalid room entry: invalid format
	err = lem.IsValidRoom("RoomA", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid format of the room" {
		t.Errorf("Expected error for invalid room format, but got: %v", err)
	}

	// Invalid room entry: negative X coordinate
	err = lem.IsValidRoom("RoomA -1 2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, the coord X must be a non-negative number" {
		t.Errorf("Expected error for negative X coordinate, but got: %v", err)
	}

	// Invalid room entry: negative Y coordinate
	err = lem.IsValidRoom("RoomA 1 -2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, the coord Y must be a non-negative number" {
		t.Errorf("Expected error for negative Y coordinate, but got: %v", err)
	}

	// Invalid room entry: duplicate coordinates
	err = lem.IsValidRoom("RoomB 1 2", "", data)
	if err == nil || err.Error() != "❌ERROR: invalid data format, room with such coordinates already exists" {
		t.Errorf("Expected error for duplicate coordinates, but got: %v", err)
	}

	// Valid start room entry
	err = lem.IsValidRoom("StartRoom 0 5", "start", data)
	if err != nil {
		t.Errorf("Expected no error for valid start room entry, but got: %v", err)
	}

	// Valid end room entry
	err = lem.IsValidRoom("EndRoom 3 4", "end", data)
	if err != nil {
		t.Errorf("Expected no error for valid end room entry, but got: %v", err)
	}
}
