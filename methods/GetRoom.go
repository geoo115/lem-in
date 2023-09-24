package lem

// Returns a pointer to the Room with a key
func GetRoom(graph *Graph, k string) *Room {
	for i, v := range graph.Rooms {
		if v.Key == k {
			return graph.Rooms[i]
		}
	}
	return nil
}
