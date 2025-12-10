package main

import (
	"dsa/graphs"
	"fmt"
)

func main() {
	// fmt.Println("Solving DSA problems like a ninja")
	// fmt.Println("You will find below the description of the problem and the output coming from the solution")
	// problems.LengthOfLongestSubstringProblem()

	// inputArray := []int{1, 2, -1, -1, 3}
	// inputArray := []int{3, 1, -1, -1, 2}
	// inputArray := []int{3, 1, 4, -1, -1, 2}
	// root := problems.BuildTree(inputArray)
	// fmt.Println(problems.InOrderTraversal(root))
	// fmt.Println(problems.PreOrderTraversal(root))
	// fmt.Println(problems.PostOrderTraversal(root))

	// Example usage of MinHeap
	// minH := heap.NewMinHeap()

	// Insert some values
	// minH.Insert(5)
	// minH.Insert(2)
	// minH.Insert(8)
	// minH.Insert(1)
	// minH.Insert(9)

	// fmt.Println("MinHeap operations:")

	// Extract minimum values
	// for {
	// 	if val, ok := minH.ExtractMin(); ok {
	// 		fmt.Printf("Extracted: %d\n", val)
	// 	} else {
	// 		break
	// 	}
	// }

	// maxH := heap.NewMaxHeap()
	// maxH.Insert(2)
	// maxH.Insert(10)
	// maxH.Insert(9)
	// maxH.Insert(5)
	// maxH.Insert(3)
	// maxH.Insert(4)

	// fmt.Println("MaxHeap operations:")

	// Extract minimum values
	// for {
	// 	if val, ok := maxH.ExtractMax(); ok {
	// 		fmt.Printf("Extracted: %d\n", val)
	// 	} else {
	// 		break
	// 	}
	// }

	graph := graphs.NewGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(1, 7)
	graph.AddEdge(7, 8)

	fmt.Println("Adjacency list :", graph.Adj)

	graph.BFS(1)
	graph.DFS(1)

	matrixGraph := graphs.NewMatrixGraph(9)

	matrixGraph.AddEdge(1, 2)
	matrixGraph.AddEdge(2, 3)
	matrixGraph.AddEdge(3, 4)
	matrixGraph.AddEdge(3, 5)
	matrixGraph.AddEdge(4, 5)
	matrixGraph.AddEdge(5, 6)
	matrixGraph.AddEdge(1, 7)
	matrixGraph.AddEdge(7, 8)

	fmt.Println("Adjacency Matrix (For Reference):")
	for _, row := range matrixGraph.Adj {
		fmt.Println(row)
	}

	matrixGraph.BFS(1)
	matrixGraph.DFS(1)
}
