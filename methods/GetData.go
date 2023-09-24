package lem

import (
	"errors"
	"strconv"
	"strings"
)

// GetData extracts relevant data from the input file and populates the Input struct
func GetData(data *Input, dataFile []string) error {
	for i, v := range dataFile {
		switch {
		case i == 0:
			ants, err := strconv.Atoi(v)
			if err != nil || ants <= 0 {
				return errors.New("❌ERROR: invalid data format, invalid number of Ants")
			}
			data.Ants = ants
		case v == "##start" && i != len(dataFile)-1:
			if err := IsValidRoom(dataFile[i+1], "start", data); err != nil {
				return err
			}
		case v == "##end" && i != len(dataFile)-1:
			if err := IsValidRoom(dataFile[i+1], "end", data); err != nil {
				return err
			}
		case strings.Contains(v, " "):
			if err := IsValidRoom(v, "", data); err != nil {
				return err
			}
		case strings.Contains(v, "-"):
			node := strings.Split(v, "-")
			if len(node) != 2 {
				return errors.New("❌ERROR: invalid data format, invalid link between rooms")
			}
			data.Links = append(data.Links, node)
		}
	}
	return nil
}
