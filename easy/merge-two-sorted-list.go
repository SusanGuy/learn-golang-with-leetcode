/**
Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.



Example 1:

Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:

Input: l1 = [], l2 = []
Output: []

Example 3:

Input: l1 = [], l2 = [0]
Output: [0]

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{Val: -1, Next: nil}
	temp := &result
	for l1 != nil && l2 != nil {
		temp.Next = new(ListNode)
		if l1.Val < l2.Val {
			temp.Next.Val = l1.Val
			l1 = l1.Next
		} else {
			temp.Next.Val = l2.Val
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	} else if l2 != nil {
		temp.Next = l2
	}

	return result.Next
}