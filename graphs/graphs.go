package graphs

import "fmt"

type Graph struct {
	Adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u)
}

func (g *Graph) BFS(startVertex int) {
	if _, exits := g.Adj[startVertex]; !exits {
		return
	}
	visited := make(map[int]bool)
	queue := []int{}

	queue = append(queue, startVertex)
	visited[startVertex] = true

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]

		fmt.Println(vertex)
		neighbors := g.Adj[vertex]
		for _, n := range neighbors {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	fmt.Println("BFS Complete")
}

func (g *Graph) DFS_Recursive(node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	fmt.Println(node)
	visited[node] = true
	if neighbors, ok := g.Adj[node]; ok {
		for _, vertex := range neighbors {
			g.DFS_Recursive(vertex, visited)
		}
	}
}

func (g *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)
	fmt.Println("Checking DFS Way")
	g.DFS_Recursive(startVertex, visited)
}
