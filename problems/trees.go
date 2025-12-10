package Problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
}

func InOrderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, InOrderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InOrderTraversal(root.Right)...)
	return result
}

func PreOrderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, PreOrderTraversal(root.Left)...)
	result = append(result, PreOrderTraversal(root.Right)...)
	return result
}

func PostOrderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, PostOrderTraversal(root.Left)...)
	result = append(result, PostOrderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// BuildTree takes a level-order array representation (where -1 is nil)
// and returns the root of the constructed Binary Tree.
func BuildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	// 1. Create the root node
	root := &TreeNode{Val: arr[0]}

	// 2. Initialize a Queue for BFS
	queue := []*TreeNode{root}

	// 3. Start processing the array from the second element
	i := 1
	for len(queue) > 0 && i < len(arr) {

		// Dequeue the parent node
		parent := queue[0]
		queue = queue[1:]

		// --- Process Left Child ---
		if i < len(arr) {
			val := arr[i]
			if val != -1 {
				leftNode := &TreeNode{Val: val}
				parent.Left = leftNode
				queue = append(queue, leftNode) // Enqueue the new node
			}
			i++ // Move to the next array index
		}

		// --- Process Right Child ---
		if i < len(arr) {
			val := arr[i]
			if val != -1 {
				rightNode := &TreeNode{Val: val}
				parent.Right = rightNode
				queue = append(queue, rightNode) // Enqueue the new node
			}
			i++ // Move to the next array index
		}
	}

	return root
}
