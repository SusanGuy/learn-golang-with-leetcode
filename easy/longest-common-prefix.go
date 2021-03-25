/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

*/

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	common := strs[0]
	for i := 1; i < len(strs); i++ {

		for strings.Index(strs[i], common) != 0 {
			common = common[:len(common)-1]
		}
		if len(common) == 0 {
			return common
		}
	}
	return common
}