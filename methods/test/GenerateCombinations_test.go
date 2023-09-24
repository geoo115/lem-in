package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestGenerateCombinations(t *testing.T) {
	graph := &lem.Graph{}
	_ = lem.AddRoom(graph, "Room1")
	_ = lem.AddRoom(graph, "Room2")
	_ = lem.AddRoom(graph, "Room3")
	_ = lem.AddEdge(graph, "Room1", "Room2")
	_ = lem.AddEdge(graph, "Room2", "Room3")

	pathsInfo := &lem.Allpaths{
		Paths: []*lem.Path{
			{P: []*lem.Room{lem.GetRoom(graph, "Room1"), lem.GetRoom(graph, "Room2"), lem.GetRoom(graph, "Room3")}},
			{P: []*lem.Room{lem.GetRoom(graph, "Room1"), lem.GetRoom(graph, "Room2")}},
		},
	}

	lem.GenerateCombinations(pathsInfo, graph)

	if len(pathsInfo.Combo) != 2 {
		t.Errorf("Expected 2 combinations, but got %d", len(pathsInfo.Combo))
	}

	expectedCombination1 := []string{"Room1", "Room2", "Room3"}
	for i, path := range pathsInfo.Combo[0][0].P {
		if path.Key != expectedCombination1[i] {
			t.Errorf("Expected room %s at position %d in combination 1, but got %s", expectedCombination1[i], i, path.Key)
		}
	}

	expectedCombination2 := []string{"Room1", "Room2"}
	for i, path := range pathsInfo.Combo[1][0].P {
		if path.Key != expectedCombination2[i] {
			t.Errorf("Expected room %s at position %d in combination 2, but got %s", expectedCombination2[i], i, path.Key)
		}
	}
}
