/**
Given a string columnTitle that represents the column title as appear in an Excel sheet, return its corresponding column number.
For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:
Input: columnTitle = "A"
Output: 1

Example 2:
Input: columnTitle = "AB"
Output: 28

Example 3:
Input: columnTitle = "ZY"
Output: 701

Example 4:
Input: columnTitle = "FXSHRXW"
Output: 2147483647
*/

func titleToNumber(columnTitle string) int {
	total := 0
	for _, char := range columnTitle {
		val := int(char-'A') + 1
		total = total*26 + val
	}
	return total
}