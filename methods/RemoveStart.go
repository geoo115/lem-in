package lem

// RemoveStart removes the start room from the paths in the optimal path
func RemoveStart(p *Allpaths) {
	for _, path := range p.OptPath {
		path.P = path.P[1:] // Remove the first room (start room) from the path
	}
}
