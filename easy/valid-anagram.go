/**
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

*/

func isAnagram(s string, t string) bool {
	arr1 := []rune(s)
	arr2 := []rune(t)
	set := make(map[rune]int)
	for _, a := range arr1 {
		if _, ok := set[a]; ok {
			set[a] = set[a] + 1
		} else {
			set[a] = 1
		}
	}
	for _, a := range arr2 {
		_, ok := set[a]
		if !ok {
			return false
		}
		set[a] = set[a] - 1
		if set[a] == 0 {
			delete(set, a)
		}
	}

	return len(set) == 0
}