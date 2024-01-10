package main

import "fmt"

/*

2385. Amount of Time for Binary Tree to Be Infected

You are given the root of a binary tree with unique values, and an integer start. At minute 0, an infection starts from the node with value start.

Each minute, a node becomes infected if:

The node is currently uninfected.
The node is adjacent to an infected node.
Return the number of minutes needed for the entire tree to be infected.



Example 1:


Input: root = [1,5,3,null,4,10,6,9,2], start = 3
Output: 4
Explanation: The following nodes are infected during:
- Minute 0: Node 3
- Minute 1: Nodes 1, 10 and 6
- Minute 2: Node 5
- Minute 3: Node 4
- Minute 4: Nodes 9 and 2
It takes 4 minutes for the whole tree to be infected so we return 4.
Example 2:


Input: root = [1], start = 1
Output: 0
Explanation: At minute 0, the only node in the tree is infected so we return 0.


Constraints:

The number of nodes in the tree is in the range [1, 105].
1 <= Node.val <= 105
Each node has a unique value.
A node with a value of start exists in the tree.

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBST(input []interface{}) *TreeNode {
	root := &TreeNode{Val: input[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		// pop first element
		node := queue[0]
		queue = queue[1:]

		if i < len(input) && input[i] != nil {
			node.Left = &TreeNode{Val: input[i].(int)}
			queue = append(queue, node.Left)
		}

		i++

		if i < len(input) && input[i] != nil {
			node.Right = &TreeNode{Val: input[i].(int)}
			queue = append(queue, node.Right)
		}

		i++

	}

	return root
}

type Graph map[int][]int

func buildUndirectedGraph(root *TreeNode) Graph {
	graph := make(Graph)

	var buildRelation func(node *TreeNode)

	buildRelation = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildRelation(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildRelation(node.Right)
		}
	}
	buildRelation(root)

	return graph
}

func amountOfTime(root *TreeNode, start int) int {
	if root == nil {
		return 0
	}

	// convert the tree to  the undirected graph
	// it means mapping each node with relation
	// something like this
	// map[1:[5 3] 2:[4] 3:[1 10 6] 4:[5 9 2] 5:[1 4] 6:[3] 9:[4] 10:[3]
	undirectedGraph := buildUndirectedGraph(root)

	minutes := 0
	mapVisited := make(map[int]bool)
	// initiate queue and visited from start
	queue := []int{start}
	mapVisited[start] = true

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			for _, num := range undirectedGraph[node] {
				// only check the node that not visited yet
				if !mapVisited[num] {
					mapVisited[num] = true
					queue = append(queue, num)
				}
			}
		}
		// pop the element based on latest size
		// FIFO (first in first out)
		queue = queue[size:]
		// as long as there is queue to check
		// we adding the minute
		if len(queue) > 0 {
			minutes++
		}
	}

	return minutes
}

func main() {
	input := []interface{}{1, 5, 3, nil, 4, 10, 6, 9, 2}
	root := createBST(input)

	fmt.Println(amountOfTime(root, 3))

	input2 := []interface{}{2, 5}
	root2 := createBST(input2)
	fmt.Println(amountOfTime(root2, 5))
}
