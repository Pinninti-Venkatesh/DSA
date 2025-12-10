package graphs

import "fmt"

type MatrixGraph struct {
	Adj         [][]int
	NumVertices int
}

func NewMatrixGraph(V int) *MatrixGraph {
	matrix := make([][]int, V)
	for i := range matrix {
		matrix[i] = make([]int, V)
	}
	return &MatrixGraph{
		Adj:         matrix,
		NumVertices: V,
	}
}

func (g *MatrixGraph) AddEdge(u, v int) {
	if u >= g.NumVertices || v >= g.NumVertices {
		return
	}
	g.Adj[u][v] = 1
	g.Adj[v][u] = 1
}

func (g *MatrixGraph) BFS(startVetex int) {
	V := g.NumVertices
	if startVetex >= V {
		return
	}
	visited := make(map[int]bool)
	visited[startVetex] = true
	queue := []int{startVetex}

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		fmt.Println(u)
		for v := 0; v < V; v++ {
			if g.Adj[u][v] == 1 && !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	fmt.Println("BFS Complete")

}
func (g *MatrixGraph) DFS_Recursive(node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	fmt.Println(node)
	visited[node] = true
	for v := 0; v < g.NumVertices; v++ {
		if g.Adj[node][v] == 1 {
			g.DFS_Recursive(v, visited)
		}

	}
}

func (g *MatrixGraph) DFS(startVetex int) {
	visited := make(map[int]bool)
	g.DFS_Recursive(startVetex, visited)
}
