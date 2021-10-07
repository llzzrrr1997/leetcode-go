package foundation

import (
	"fmt"
	"strings"
)

//no155
type MinStack struct {
	Data []int
	Min  []int
}

func MinStackConstructor() MinStack {
	stack := MinStack{}
	stack.Data = make([]int, 0)
	stack.Min = make([]int, 0)
	return stack
}

func (this *MinStack) Push(val int) {
	this.Data = append(this.Data, val)
	n := len(this.Min)
	if n == 0 {
		this.Min = append(this.Min, val)
		return
	}
	if val < this.Min[n-1] {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, this.Min[n-1])
	}
}

func (this *MinStack) Pop() {
	this.Data = this.Data[:len(this.Data)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

//no1249
func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0)
	for i, val := range s {
		if val == '(' || val == ')' {
			if val == ')' {
				n := len(stack)
				if n == 0 || s[stack[n-1]] == ')' {
					stack = append(stack, i)
				} else if s[stack[n-1]] == '(' {
					stack = stack[:n-1]
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	fmt.Println(stack, int32('('))
	sb := strings.Builder{}
	for i, val := range s {
		if len(stack) != 0 && i == stack[0] {
			stack = stack[1:]
		} else {
			sb.WriteRune(val)
		}
	}
	return sb.String()
}

//no1823
func findTheWinner(n int, k int) int {
	data := make([]int, n)
	for i := 1; i <= n; i++ {
		data[i-1] = i
	}
	start := 0
	for len(data) > 1 {
		start = (start+k)%len(data) - 1
		if start < 0 {
			start += len(data)
		}
		data = append(data[:start], data[start+1:]...)
	}
	return data[0]
}
