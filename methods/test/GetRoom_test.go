package lem_test

import (
	lem "lem/methods"
	"testing"
)

func TestGetRoom(t *testing.T) {
	graph := &lem.Graph{}
	_ = lem.AddRoom(graph, "Room1")
	_ = lem.AddRoom(graph, "Room2")

	room := lem.GetRoom(graph, "Room1")
	if room == nil {
		t.Error("Expected non-nil room, but got nil")
	}
}
