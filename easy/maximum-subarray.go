/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:

Input: nums = [1]
Output: 1

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23


*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = calculateMax(nums[i], nums[i]+dp[i-1])
		max = calculateMax(dp[i], max)
	}

	return max
}

func calculateMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

