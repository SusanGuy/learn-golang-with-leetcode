/**
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
Example 1:
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:
Input: numRows = 1
Output: [[1]]

*/

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}

	for i := 1; i < numRows; i++ {
		temp := []int{1}
		prev := result[i-1]
		for j := 1; j < i; j++ {
			temp = append(temp, prev[j]+prev[j-1])
		}
		temp = append(temp, 1)
		result = append(result, temp)
	}
	return result
}
