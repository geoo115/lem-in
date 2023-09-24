package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestCalculateTravelTimesForEachCombination(t *testing.T) {
	// Create some sample data for testing
	path1 := &lem.Path{P: []*lem.Room{{}, {}}}
	path2 := &lem.Path{P: []*lem.Room{{}, {}, {}}}
	path3 := &lem.Path{P: []*lem.Room{{}, {}, {}, {}}}

	combination1 := []*lem.Path{path1, path2}
	combination2 := []*lem.Path{path2, path3}
	combinations := [][]*lem.Path{combination1, combination2}

	pathsInfo := &lem.Allpaths{Combo: combinations}

	totalAnts := 10

	// Call the function being tested
	travelTimes := lem.CalculateTravelTimesForEachCombination(pathsInfo, totalAnts)

	// Define the expected travel times based on the given data
	expectedTravelTimes := []int{
		(totalAnts + (len(path1.P) + len(path2.P) - 1)) / len(combination1),
		(totalAnts + (len(path2.P) + len(path3.P) - 1)) / len(combination2),
	}

	// Compare the expected travel times with the calculated travel times
	for i, expected := range expectedTravelTimes {
		if travelTimes[i] != expected {
			t.Errorf("Expected travel time %d for combination %d, but got %d", expected, i, travelTimes[i])
		}
	}
}
