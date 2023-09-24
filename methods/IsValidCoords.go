package lem

// IsValidCoords checks if the room coordinates are valid and unique
func IsValidCoords(nodeCoord []string, data *Input) bool {
	for _, coord := range data.Coords {
		if coord[0] == nodeCoord[0] && coord[1] == nodeCoord[1] {
			return false
		}
	}
	return true
}
