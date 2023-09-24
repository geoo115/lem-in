package lem

// Checks the adding Room for existency
func RoomExists(s []*Room, key string) bool {
	for _, v := range s {
		if v.Key == key {
			return true
		}
	}
	return false
}
