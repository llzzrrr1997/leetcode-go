package getting_started

//no20
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		if s[i] == ')' && len(stack) != 0 {
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
			continue
		} else if s[i] == ')' && len(stack) == 0 {
			return false
		}
		if s[i] == '}' && len(stack) != 0 {
			if stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
			continue
		} else if s[i] == '}' && len(stack) == 0 {
			return false
		}
		if s[i] == ']' && len(stack) != 0 {
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
			continue
		} else if s[i] == ']' && len(stack) == 0 {
			return false
		}
	}
	return len(stack) == 0
}

//232
type MyQueue struct {
	InStack  []int
	OutStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.OutStack) == 0 {
		this.move()
	}
	x := this.OutStack[len(this.OutStack)-1]
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return x
}

func (this *MyQueue) move() {
	for i := len(this.InStack) - 1; i >= 0; i-- {
		this.OutStack = append(this.OutStack, this.InStack[i])
		this.InStack = this.InStack[:i]
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.OutStack) == 0 {
		this.move()
	}
	return this.OutStack[len(this.OutStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.InStack) == 0 && len(this.OutStack) == 0
}
