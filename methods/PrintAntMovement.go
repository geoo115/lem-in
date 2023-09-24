package lem

import (
	"fmt"
	"strconv"
)

// PrintAntMigration prints the ant migration along the optimal path using ant allocations
func PrintAntMovement(p *Allpaths, ants int) {
	// Initialize a 2D slice to store ant movements for each tick
	movementOutput := make([][]string, p.MinTravelTime-1)
	antName := 1

	// Allocate ants to paths using the AllocateAnts function
	antAllocations := AllocateAnts(p, ants)

	// Remove the start room from the optimal path
	RemoveStart(p)

	// Iterate over allocated ants for each path
	for i, antCount := range antAllocations {
		// Iterate over each ant's movements along the path
		for j := 0; j < antCount; j++ {
			// Iterate over rooms in the path
			for k, room := range p.OptPath[i].P {
				// Append ant movement to the corresponding tick in the movementOutput
				movementOutput[k+j] = append(movementOutput[k+j], "L"+strconv.Itoa(antName)+"-"+room.Key)
			}
			// Increment the ant identifier
			antName++
		}
	}

	// Print the ant movements
	fmt.Println("----------- Ants Movements ------------- ")
	for _, level := range movementOutput {
		for _, movement := range level {
			fmt.Print(movement + " ")
		}
		fmt.Println()
	}
	fmt.Println("Number of turns:", len(movementOutput))
}
