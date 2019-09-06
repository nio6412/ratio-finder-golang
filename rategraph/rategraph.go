package rategraph

type DirectedEdge struct {
	End    string
	Weight float32
}

type Rate struct {
	Start string
	End   string
	Rate  float32
}

type RateGraph struct {
	graph map[string][]DirectedEdge
}

func NewRateGraph(rates []Rate) *RateGraph {
	rg := RateGraph{make(map[string][]DirectedEdge)}

	for _, r := range rates {
		rg.AddConversion(r.Start, r.End, r.Rate)
	}

	return &rg
}

func (rg *RateGraph) AddConversion(orig string, dest string, rate float32) {
	nodes, ok := rg.graph[orig]
	if !ok {
		rg.graph[orig] = []DirectedEdge{DirectedEdge{dest, rate}}
	} else {
		rg.graph[orig] = append(nodes, DirectedEdge{dest, rate})
	}
}

func (rg *RateGraph) GetNeighbors(node string) []DirectedEdge {
	return rg.graph[node]
}

func (rg *RateGraph) GetNodes() []string {
	keys := make([]string, len(rg.graph))
	i := 0
	for k := range rg.graph {
		keys[i] = k
		i++
	}

	return keys
}

func depthFirstSearch(rg *RateGraph, start string, dest string, rFromOrig float32, visited map[string]interface{}) {

}
