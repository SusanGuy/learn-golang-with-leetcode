/**
Given the root of a binary search tree, and an integer k, return the kth (1-indexed) smallest element in the tree.

Example 1:

Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count int = 0
var result int = -1

func kthSmallest(root *TreeNode, k int) int {
	inOrderTraverse(root, k)
	return result
}

func inOrderTraverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inOrderTraverse(root.Left, k)
	count++
	if count == k {
		result = root.Val
		return
	}
	inOrderTraverse(root.Right, k)
}