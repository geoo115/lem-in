package lem

func AllocateAnts(p *Allpaths, ants int) []int {
	// Create a slice to store the number of ants allocated to each path
	antAllocations := make([]int, len(p.OptPath))

	// Continue allocating ants until there are no ants left to allocate
	for ants != 0 {
		// Iterate over each path in the optimal path combination
		for i, path := range p.OptPath {
			// Calculate the maximum number of ants that can be allocated to the current path
			maxAnts := p.MinTravelTime - (len(path.P) - 1)

			// If the calculated number of ants is greater than the remaining ants
			if maxAnts > ants {
				maxAnts = ants
			}

			// Allocate the calculated number of ants to the current path
			antAllocations[i] += maxAnts

			// Reduce the remaining ants by the allocated amount
			ants -= maxAnts
		}
	}

	// Return the slice containing the final ant allocations for each path
	return antAllocations
}
