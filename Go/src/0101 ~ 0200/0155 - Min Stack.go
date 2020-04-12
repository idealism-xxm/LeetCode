// 链接：https://leetcode.com/problems/min-stack/
// 题意：实现一个栈，并支持在 O(1) 内进行 push, pop, top 和获取最小值 ？

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> Returns -3.
// minStack.pop();
// minStack.top();      --> Returns 0.
// minStack.getMin();   --> Returns -2.

// 思路1： 普通栈 + 单调栈
//
//		这次完全没想到竟然还可以这样处理
//		维护两个栈，
//		一个是正常的栈，会处理所有的数，
//		另一个是单调递减栈，仅当其为空或者栈顶数字大于等于当前数字时，才入栈
//		（取等于是为了保证 pop 时方便处理，确保栈顶仍是最小数）
//
//  	时间复杂度： O(1)
//		空间复杂度： O(n)

type MinStack struct {
	// 存储所有数的栈
	allStack []int
	// 只有 minStack 为空或者新加入的数字更小/相等时，才放入此栈
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	// 所有数都会放入 allStack
	this.allStack = append(this.allStack, x)
	// 只有 minStack 为空或者新加入的数字更小/相等时，才放入 minStack
	if minTop := len(this.minStack) - 1; minTop == -1 || x <= this.minStack[minTop] {
		this.minStack = append(this.minStack, x)
	}
}


func (this *MinStack) Pop()  {
	allTop := len(this.allStack) - 1
	topNum := this.allStack[allTop]
	// allStack 的栈顶数字出栈
	this.allStack = this.allStack[:allTop]
	// 只有两个栈的栈顶数字相同时， minStack 才出栈
	if minTop := len(this.minStack) - 1; topNum == this.minStack[minTop] {
		this.minStack = this.minStack[:minTop]
	}
}


func (this *MinStack) Top() int {
	return this.allStack[len(this.allStack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// 思路2： 普通栈
//
//		后来发现只用普通栈也可以（其实这个应该是最容易想到的，但是没有考虑到加字段的方式）
//		栈中的每个元素是一个结点，有两个值：栈顶数字和栈中最小数字
//		由于每次只会进行 push 和 pop ，那么
//		新入栈的结点的最小值 = min(栈顶最小值, 新入栈的数字)
//
//  	时间复杂度： O(1)
//		空间复杂度： O(n)

type node struct {
	Val int
	Min int
}

type MinStack struct {
	stack []node
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	if top := len(this.stack) - 1; top == -1 {
		// 栈空时，直接放入即可
		this.stack = []node{node{Val: x, Min: x}}
	} else {
		// 栈非空时，最小值需要取较小者
		this.stack = append(this.stack, node{Val: x, Min: min(x, this.stack[top].Min)})
	}
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1].Val
}


func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack) - 1].Min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
