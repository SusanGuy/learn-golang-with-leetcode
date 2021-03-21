
/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2
*/

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	maxNumber := nums[0]
	max := 1
	dict[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; ok {
			dict[nums[i]] = dict[nums[i]] + 1
		} else {
			dict[nums[i]] = 1
		}
		if dict[nums[i]] > max {
			max = dict[nums[i]]
			maxNumber = nums[i]
		}
	}
	return maxNumber
}