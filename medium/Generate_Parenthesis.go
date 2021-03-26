/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]

*/

func generateParenthesis(n int) []string {
	var result []string
	if n == 0 {
		return result
	}
	helper(&result, "", n, 0, 0)
	return result
}

func helper(result *[]string, s string, n, Open, Close int) {
	if len(s) == 2*n {
		*result = append(*result, s)
		return
	}
	if Open < n {
		helper(result, s+"(", n, Open+1, Close)
	}
	if Close < Open {
		helper(result, s+")", n, Open, Close+1)
	}
}