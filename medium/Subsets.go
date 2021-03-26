/**
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

Input: nums = [0]
Output: [[],[0]]


*/

func subsets(nums []int) [][]int {
	result := new([][]int)
	helper(result, nums, new([]int), 0)
	return *result
}

func helper(result *[][]int, nums []int, subset *[]int, index int) {
	temp := make([]int, len(*subset))
	copy(temp, *subset)
	*result = append(*result, temp)
	for i := index; i < len(nums); i++ {
		*subset = append(*subset, nums[i])
		helper(result, nums, subset, i+1)
		*subset = (*subset)[:len(*subset)-1]
	}
}