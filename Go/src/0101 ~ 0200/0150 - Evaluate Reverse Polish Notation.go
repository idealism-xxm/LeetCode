// 链接：https://leetcode.com/problems/evaluate-reverse-polish-notation/
// 题意：给定一个逆波兰表达式，求解最终结果 ？

// 输入： ["2", "1", "+", "3", "*"]
// 输出： 9
// 解释： ((2 + 1) * 3) = 9

// 输入： ["4", "13", "5", "/", "+"]
// 输出： 6
// 解释： (4 + (13 / 5)) = 6

// 输入： ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// 输出： 22
// 解释：
//	  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//  = ((10 * (6 / (12 * -11))) + 17) + 5
//  = ((10 * (6 / -132)) + 17) + 5
//  = ((10 * 0) + 17) + 5
//  = (0 + 17) + 5
//  = 17 + 5
//  = 22

// 思路： 模拟
//
//		逆波兰表达式求值很简单，
//		每次遇到一个数字时，直接入栈即可
//		每次遇到一个运算符时，将栈顶的两个数拿出来计算，将结果再放回栈顶即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

import "strconv"

func evalRPN(tokens []string) int {
	// tokens 最多有 (len(tokens) + 1) >> 1 数字
	stackCapacity := (len(tokens) + 1) >> 1
	stack := make([]int, stackCapacity)
	top := -1
	for _, token := range tokens {
		switch token {
		// 运算符直接取栈顶两个元素，结果入栈即可
		case "+":
			stack[top - 1] += stack[top]
			top--
		case "-":
			stack[top - 1] -= stack[top]
			top--
		case "*":
			stack[top - 1] *= stack[top]
			top--
		case "/":
			stack[top - 1] /= stack[top]
			top--
		default:
			// 数字直接入栈即可
			top++
			stack[top], _ = strconv.Atoi(token)
		}
	}
	return stack[0]
}
