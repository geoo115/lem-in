package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestAddRoom(t *testing.T) {
	graph := &lem.Graph{}
	err := lem.AddRoom(graph, "Room1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = lem.AddRoom(graph, "Room1")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
