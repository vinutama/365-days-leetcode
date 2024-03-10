package main

/*
876. Middle of the Linked List
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.



Example 1:


Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
Example 2:


Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	nodes := make([]*ListNode, 0)

	curr := head
	for curr != nil {
		node := &ListNode{Val: curr.Val, Next: curr.Next}
		nodes = append(nodes, node)
		curr = curr.Next
	}
	middle_node_pos := len(nodes) / 2
	return nodes[middle_node_pos-1].Next
}
