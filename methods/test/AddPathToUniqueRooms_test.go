package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestAddPathToUniqueRooms(t *testing.T) {
	room1 := &lem.Room{Key: "Room1"}
	room2 := &lem.Room{Key: "Room2"}
	room3 := &lem.Room{Key: "Room3"}

	uniqueRooms := map[string]*lem.Room{
		"Room1": room1,
	}

	path := &lem.Path{
		P: []*lem.Room{room2, room3},
	}

	// Call the function to add rooms from the path to uniqueRooms
	lem.AddPathToUniqueRooms(uniqueRooms, path)

	// Check if all rooms from the path are added to uniqueRooms
	if uniqueRooms["Room2"] != room2 {
		t.Errorf("Expected Room2 to be in uniqueRooms, but it's not")
	}
	if uniqueRooms["Room3"] != room3 {
		t.Errorf("Expected Room3 to be in uniqueRooms, but it's not")
	}
}
