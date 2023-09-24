package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestAllocateAnts(t *testing.T) {
	room1 := &lem.Room{}
	room2 := &lem.Room{}
	path1 := &lem.Path{P: []*lem.Room{room1, room2}}
	path2 := &lem.Path{P: []*lem.Room{room2, room1}}
	optimalPath := []*lem.Path{path1, path2}
	pathsInfo := &lem.Allpaths{
		OptPath:       optimalPath,
		MinTravelTime: 5,
	}
	ants := 10
	antAllocations := lem.AllocateAnts(pathsInfo, ants)
	expectedPath1Allocation := 6
	expectedPath2Allocation := 4

	if antAllocations[0] != expectedPath1Allocation {
		t.Errorf("Expected allocation for path1: %d, but got %d", expectedPath1Allocation, antAllocations[0])
	}
	if antAllocations[1] != expectedPath2Allocation {
		t.Errorf("Expected allocation for path2: %d, but got %d", expectedPath2Allocation, antAllocations[1])
	}
}
