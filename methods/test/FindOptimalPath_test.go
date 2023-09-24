package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestFindOptimalPath(t *testing.T) {
	// Create some sample data for testing
	path1 := &lem.Path{P: []*lem.Room{{}, {}}}
	path2 := &lem.Path{P: []*lem.Room{{}, {}, {}}}
	path3 := &lem.Path{P: []*lem.Room{{}, {}, {}, {}}}

	combination1 := []*lem.Path{path1, path2}
	combination2 := []*lem.Path{path2, path3}
	combinations := [][]*lem.Path{combination1, combination2}

	pathsInfo := &lem.Allpaths{
		Combo: combinations,
	}

	ticks := []int{10, 5} // Sample travel times

	// Create a sample graph
	graph := &lem.Graph{
		Start: &lem.Room{},
		End:   &lem.Room{},
	}

	// Call the function being tested
	lem.FindOptimalPath(pathsInfo, ticks, graph)

	// Define the expected optimal path and travel time
	expectedOptimalPath := combination2
	expectedMinTravelTime := 5

	// Compare the lengths of the optimal path
	if len(pathsInfo.OptPath) != len(expectedOptimalPath) {
		t.Errorf("Expected optimal path length %d, but got %d", len(expectedOptimalPath), len(pathsInfo.OptPath))
	}

	// Compare the contents of the optimal path
	for i := range pathsInfo.OptPath {
		if pathsInfo.OptPath[i] != expectedOptimalPath[i] {
			t.Errorf("Expected optimal path element %v, but got %v", expectedOptimalPath[i], pathsInfo.OptPath[i])
		}
	}

	// Compare the expected minimum travel time with the calculated minimum travel time
	if pathsInfo.MinTravelTime != expectedMinTravelTime {
		t.Errorf("Expected minimum travel time %d, but got %d", expectedMinTravelTime, pathsInfo.MinTravelTime)
	}
}
