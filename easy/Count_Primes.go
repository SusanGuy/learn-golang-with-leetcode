/**
Count the number of prime numbers less than a non-negative number, n.



Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:

Input: n = 0
Output: 0

Example 3:

Input: n = 1
Output: 0

*/

func countPrimes(n int) int {
	arr := make([]bool, n)
	for i, _ := range arr {
		arr[i] = true
	}
	for i := 2; i*i < len(arr); i++ {
		if arr[i] {
			for j := i; j*i < len(arr); j++ {
				arr[i*j] = false
			}
		}
	}
	count := 0
	for i := 2; i < len(arr); i++ {
		if arr[i] {
			count++
		}
	}
	return count
}