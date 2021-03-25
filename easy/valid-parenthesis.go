
/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.



Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

Example 4:

Input: s = "([)]"
Output: false

Example 5:

Input: s = "{[]}"
Output: true

*/

type Stack []byte

func (this *Stack) push(letter byte) {
	*this = append(*this, letter)
}

func (this *Stack) pop() byte {
	val := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return val
}

func (this *Stack) peek() byte {
	return (*this)[len(*this)-1]
}

func (this *Stack) isEmpty() bool {
	return len(*this) == 0
}

func (this *Stack) size() int {
	return len(*this)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := new(Stack)
	for _, val := range []byte(s) {
		if stack.isEmpty() {
			stack.push(val)
			continue
		}
		top := stack.peek()
		if top == '(' && val == ')' || top == '{' && val == '}' || top == '[' && val == ']' {
			stack.pop()
		} else {
			stack.push(val)
		}
	}
	return stack.isEmpty()
}