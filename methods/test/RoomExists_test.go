package lem_test

import (
	lem "lem/methods"
	"testing"
)

// Checks the adding Room for existency
func TestRoomExists(t *testing.T) {
	rooms := []*lem.Room{
		{Key: "Room1"},
		{Key: "Room2"},
	}

	exists := lem.RoomExists(rooms, "Room1")
	if !exists {
		t.Error("Expected room to exist, but got false")
	}

	exists = lem.RoomExists(rooms, "Room3")
	if exists {
		t.Error("Expected room not to exist, but got true")
	}
}
