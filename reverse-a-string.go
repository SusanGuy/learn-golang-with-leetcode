/*
Write a function that reverses a string. The input string is given as an array of characters s.
Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

*/

package main

func reverseString(s []byte) {
	pointer1 := 0
	pointer2 := len(s) - 1
	for pointer1 < pointer2 {
		temp := s[pointer1]
		s[pointer1] = s[pointer2]
		s[pointer2] = temp
		pointer1++
		pointer2--
	}
}
