package main

import (
	"fmt"
)

/*
1026. Maximum Difference Between Node and Ancestor

Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.

A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.



Example 1:


Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
Example 2:


Input: root = [1,null,2,null,0,3]
Output: 3


Constraints:

The number of nodes in the tree is in the range [2, 5000].
0 <= Node.val <= 105

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

type GraphDetails struct {
	maxVal int
	minVal int
}
type Graph map[int]GraphDetails

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func createNodeAttributes(root *TreeNode) Graph {
	graph := make(Graph)

	graph[root.Val] = GraphDetails{maxVal: root.Val, minVal: root.Val}

	var searchMinMaxPathNode func(node *TreeNode)

	searchMinMaxPathNode = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			graph[node.Left.Val] = GraphDetails{maxVal: max(node.Left.Val, graph[node.Val].maxVal), minVal: min(node.Left.Val, graph[node.Val].minVal)}
			searchMinMaxPathNode(node.Left)
		}

		if node.Right != nil {
			graph[node.Right.Val] = GraphDetails{maxVal: max(node.Right.Val, graph[node.Val].maxVal), minVal: min(node.Right.Val, graph[node.Val].minVal)}
			searchMinMaxPathNode(node.Right)
		}
	}
	searchMinMaxPathNode(root)

	return graph
}

func maxAncestorDiff(root *TreeNode) int {
	nodeAttributes := createNodeAttributes(root)
	maxDiffVal := 0

	for _, graph := range nodeAttributes {
		diff := graph.maxVal - graph.minVal
		maxDiffVal = max(diff, maxDiffVal)
	}

	return maxDiffVal

}

func main() {
	input := []interface{}{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13}

	root := createBST(input)

	// input2 := []interface{}{1, nil, 2, nil, 0, 3}
	// root2 := createBST(input2)

	fmt.Println(maxAncestorDiff(root))
	// fmt.Println(maxAncestorDiff(root2))

}
