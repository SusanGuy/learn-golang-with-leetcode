/**
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


*/

type MinStack struct {
	stack   []int
	minimum []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:   []int{},
		minimum: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minimum) == 0 {
		this.minimum = append(this.minimum, x)
	} else {
		this.minimum = append(this.minimum, min(this.minimum[len(this.minimum)-1], x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minimum = this.minimum[:len(this.minimum)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minimum[len(this.minimum)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}