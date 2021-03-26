/**
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

*/

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = nums[0]
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	nums[0] = right[1]
	nums[len(nums)-1] = left[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = left[i-1] * right[i+1]
	}
	return nums
}