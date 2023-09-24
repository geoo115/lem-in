package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestBFS(t *testing.T) {
	graph := &lem.Graph{}
	_ = lem.AddRoom(graph, "Room1")
	_ = lem.AddRoom(graph, "Room2")
	_ = lem.AddRoom(graph, "Room3")
	_ = lem.AddEdge(graph, "Room1", "Room2")
	_ = lem.AddEdge(graph, "Room2", "Room3")

	start := "Room1"
	end := "Room3"
	paths := lem.BFS(graph, start, end)

	if len(paths) != 1 {
		t.Errorf("Expected 1 path, but got %d", len(paths))
	}

	expectedPath := []string{"Room1", "Room2", "Room3"}
	for i, room := range paths[0].P {
		if room.Key != expectedPath[i] {
			t.Errorf("Expected room %s at position %d, but got %s", expectedPath[i], i, room.Key)
		}
	}
}
