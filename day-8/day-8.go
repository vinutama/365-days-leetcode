package main

import "fmt"

/*
*
938. Range Sum of BST
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].

Example 1:

Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.

Constraints:

The number of nodes in the tree is in the range [1, 2 * 104].
1 <= Node.val <= 105
1 <= low <= high <= 105
All Node.val are unique.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}

	return root
}

func createBST(input []interface{}) *TreeNode {
	root := &TreeNode{Val: input[0].(int)}

	for i := 1; i < len(input); i++ {
		if input[i] != nil {
			root = insertBST(root, input[i].(int))
		}
	}
	return root
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}

}

func main() {
	type Case struct {
		input []interface{}
		low   int
		high  int
	}

	case1 := Case{
		input: []interface{}{10, 5, 15, 3, 7, nil, 18},
		low:   7,
		high:  15,
	}

	case2 := Case{
		input: []interface{}{10, 5, 15, 3, 7, 13, 18, 1, nil, 6},
		low:   6,
		high:  10,
	}

	root1 := createBST(case1.input)

	fmt.Println(rangeSumBST(root1, case1.low, case1.high))

	root2 := createBST(case2.input)
	fmt.Println(rangeSumBST(root2, case2.low, case2.high))

}
