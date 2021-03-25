/**
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.



Example 1:

Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:

Input: root = [1,2]
Output: 1


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := calculateMax(root.Left)
	right := calculateMax(root.Right)
	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiameter := diameterOfBinaryTree(root.Right)
	return max(max(leftDiameter, rightDiameter), left+right)
}

func calculateMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(calculateMax(root.Left), calculateMax(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}