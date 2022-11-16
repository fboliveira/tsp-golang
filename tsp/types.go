package tsp

type TSPData struct {
	id     int
	pointx float64
	pointy float64
}

type TSPInstance struct {
	name        string
	problemType string
	comment     string
	Dimension   int
	edgeType    string
	Data        []TSPData
	Distance    [][]int
}

// Replaced by solution.go
// type TSPSolution struct {
// 	Tour []int
// 	Cost int
// }

type Parameters struct {
	FileName string
}
