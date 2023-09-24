package lem

func FindOptimalPath(p *Allpaths, ticks []int, graph *Graph) {
	// Initialize the minimum travel time and optimal path with values from the first combination
	p.MinTravelTime = ticks[0]
	p.OptPath = p.Combo[0]

	// If the optimal path starts with the graph's start room and ends with the end room
	if p.OptPath[0].P[0] == graph.Start && p.OptPath[0].P[1] == graph.End {
		// Modify the first combination to only contain the start and end rooms
		p.Combo[0] = p.Combo[0][0:1]
		// Update the optimal path with the modified first combination
		p.OptPath = p.Combo[0]
	}

	// Iterate over the travel times and combinations to find the optimal path
	for i, num := range ticks {
		// If the current travel time is less than the previously known optimal travel time
		if num < p.MinTravelTime {
			// Update the optimal travel time
			p.MinTravelTime = num
			// Update the optimal path with the current combination's path
			p.OptPath = p.Combo[i]
		}
	}
}
