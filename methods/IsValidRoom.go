package lem

import (
	"errors"
	"strconv"
	"strings"
)

// IsValidRoom validates and processes a room entry
func IsValidRoom(v, s string, data *Input) error {
	node := strings.Split(v, " ")
	if len(node) != 3 {
		return errors.New("❌ERROR: invalid format of the room")
	}
	if strings.ContainsAny(node[0], " #") || node[0][0] == 'L' {
		return errors.New("❌ERROR: invalid data format, the room can't start with L, # or space")
	}
	x, err := strconv.Atoi(node[1])
	if err != nil || x < 0 {
		return errors.New("❌ERROR: invalid data format, the coord X must be a non-negative number")
	}
	y, err := strconv.Atoi(node[2])
	if err != nil || y < 0 {
		return errors.New("❌ERROR: invalid data format, the coord Y must be a non-negative number")
	}
	nodeCoord := []string{node[1], node[2]}
	if !IsValidCoords(nodeCoord, data) {
		return errors.New("❌ERROR: invalid data format, room with such coordinates already exists")
	}
	switch s {
	case "start":
		data.StartR = node[0]
	case "end":
		data.EndR = node[0]
	default:
		data.Rooms = append(data.Rooms, node[0])
		data.Coords = append(data.Coords, []string{node[1], node[2]})
	}
	return nil
}
