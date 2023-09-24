package lem

// AddPathToUniqueRooms adds the rooms in the given path to the uniqueRooms map
func AddPathToUniqueRooms(uniqueRooms map[string]*Room, path *Path) {
	// Iterate through the rooms in the path
	for _, room := range path.P {
		uniqueRooms[room.Key] = room // Add the room to the uniqueRooms map
	}
}
