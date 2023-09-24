package lem_test

import (
	lem "lem/methods"
	"testing"
)

// TestCheckPathIntersection checks the CheckPathIntersection function
func TestCheckPathIntersection(t *testing.T) {
	room1 := &lem.Room{Key: "Room1"}
	room2 := &lem.Room{Key: "Room2"}
	room3 := &lem.Room{Key: "Room3"}

	uniqueRooms := map[string]*lem.Room{
		"Room1": room1,
		"Room2": room2,
	}

	path := &lem.Path{
		P: []*lem.Room{room1, room2, room3},
	}

	// Expecting intersections, so the result should be false
	intersects := lem.CheckPathIntersection(uniqueRooms, path)
	if intersects {
		t.Error("Expected path to intersect, but got true")
	}

	// Make path intersect by reusing a room already present in uniqueRooms
	path.P[2] = room2
	// Expecting intersections, so the result should be false
	intersects = lem.CheckPathIntersection(uniqueRooms, path)
	if intersects {
		t.Error("Expected path to intersect, but got true")
	}
}
