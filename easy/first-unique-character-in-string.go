/**
Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode"
return 2.


*/

func firstUniqChar(s string) int {
	set := make(map[rune]int)
	for _, a := range s {
		if _, ok := set[a]; ok {
			set[a]++
		} else {
			set[a] = 1
		}
	}
	for index, a := range s {
		if set[a] == 1 {
			return index
		}
	}

	return -1
}