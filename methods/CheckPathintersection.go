package lem

// CheckPathIntersection checks whether the given path intersects with rooms in the uniqueRooms map
func CheckPathIntersection(uniqueRooms map[string]*Room, path *Path) bool {
	// Iterate through the rooms in the path, excluding the start and end rooms
	for i := 1; i < len(path.P)-1; i++ {
		// If the room is already present in uniqueRooms, there is an intersection
		if _, ok := uniqueRooms[path.P[i].Key]; ok {
			return false // There is an intersection, so return false
		}
	}
	return true // No intersections found, so return true
}
