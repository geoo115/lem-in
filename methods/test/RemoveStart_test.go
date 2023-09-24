package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestRemoveStart(t *testing.T) {
	// Create a sample Allpaths struct and populate it with data for testing
	p := &lem.Allpaths{
		OptPath: []*lem.Path{
			{
				P: []*lem.Room{
					{Key: "Start"},
					{Key: "Room1"},
					{Key: "Room2"},
					{Key: "End"},
				},
			},
		},
	}

	// Call the function to be tested
	lem.RemoveStart(p)

	// Check if the start room has been removed from the path
	if len(p.OptPath[0].P) != 3 {
		t.Errorf("Expected path length: 3, got: %d", len(p.OptPath[0].P))
	}

	// Check if the first room is now "Room1"
	if p.OptPath[0].P[0].Key != "Room1" {
		t.Errorf("Expected first room: Room1, got: %s", p.OptPath[0].P[0].Key)
	}
}
