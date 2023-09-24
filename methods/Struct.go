package lem

type Graph struct {
	Rooms      []*Room
	Start, End *Room
}

type Room struct {
	Key      string
	Children []*Room
}

type Path struct {
	P []*Room
}

type Allpaths struct {
	Paths         []*Path
	Combo         [][]*Path
	OptPath       []*Path
	MinTravelTime int
}

type Input struct {
	Ants   int
	StartR string
	EndR   string
	Rooms  []string
	Coords [][]string
	Links  [][]string
}
