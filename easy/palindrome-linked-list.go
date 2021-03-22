/**
Given the head of a singly linked list, return true if it is a palindrome.



Example 1:

Input: head = [1,2,2,1]
Output: true

Example 2:

Input: head = [1,2]
Output: false


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	fast = head
	var prev *ListNode = nil
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	for fast != nil && prev != nil {
		if fast.Val != prev.Val {
			return false
		}
		fast = fast.Next
		prev = prev.Next
	}
	return true

}