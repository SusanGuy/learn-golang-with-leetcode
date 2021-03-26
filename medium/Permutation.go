/**
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]

*/

func permute(nums []int) [][]int {
	set := make(map[int]bool)
	var result [][]int
	var permutation []int
	permuteHelp(nums, &result, &permutation, set)
	return result
}

func permuteHelp(nums []int, list *[][]int, permutation *[]int, set map[int]bool) {
	if len(*permutation) == len(nums) {
		temp := make([]int, len(*permutation))
		copy(temp, *permutation)
		*list = append(*list, temp)
		return
	}

	for i, val := range nums {
		if _, ok := set[i]; !ok {
			set[i] = true
			*permutation = append(*permutation, val)
			permuteHelp(nums, list, permutation, set)
			*permutation = (*permutation)[:len(*permutation)-1]
			delete(set, i)
		}
	}
}