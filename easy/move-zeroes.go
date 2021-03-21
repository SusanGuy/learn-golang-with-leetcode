/**
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]

*/

func moveZeroes(nums []int) {
	index := 0
	for index < len(nums) {
		if nums[index] == 0 {
			break
		}
		index++
	}
	if index == len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			nums[i] = 0
			index++
		}
	}
}