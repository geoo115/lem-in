package lem

func GenerateCombinations(p *Allpaths, graph *Graph) {
	// Iterate through each path in the list of all paths
	for k := range p.Paths {
		// Create a map to track unique rooms in the combination
		uniqueRooms := make(map[string]*Room)
		// Temporary list to store paths in the current combination
		tmp := []*Path{}

		// Check if the current path (k) can be added to the combination
		if CheckPathIntersection(uniqueRooms, p.Paths[k]) {
			// Add the current path to the temporary list
			AddPathToUniqueRooms(uniqueRooms, p.Paths[k])
			tmp = append(tmp, p.Paths[k])
		}

		// Iterate over the remaining paths to find compatible paths for the combination
		for i := k + 1; i < len(p.Paths); i++ {
			// Check if the current path (i) can be added to the combination
			if CheckPathIntersection(uniqueRooms, p.Paths[i]) {
				// Add the current path to the temporary list
				AddPathToUniqueRooms(uniqueRooms, p.Paths[i])
				tmp = append(tmp, p.Paths[i])
			} else {
				continue
			}
		}

		// If the current path (k) is a direct route from start to end, create a separate combination
		if len(p.Paths[k].P) == 2 && p.Paths[k].P[0] == graph.Start && p.Paths[k].P[1] == graph.End {
			tmpDirect := []*Path{p.Paths[k]}
			p.Combo = append(p.Combo, tmpDirect)
		}

		// Create a new combination with the paths in the temporary list
		p.Combo = append(p.Combo, tmp)
	}
}
