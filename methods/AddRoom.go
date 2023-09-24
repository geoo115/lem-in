package lem

import "fmt"

// AddRoom adds a Room to the Graph
func AddRoom(graph *Graph, k string) error {
	if RoomExists(graph.Rooms, k) {
		err := fmt.Errorf("❌ ERROR: room %v is already exists", k)
		return err
	} else {
		graph.Rooms = append(graph.Rooms, &Room{Key: k})
	}
	return nil
}
