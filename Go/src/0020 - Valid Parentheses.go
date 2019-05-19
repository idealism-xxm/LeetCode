// 链接：https://leetcode.com/problems/valid-parentheses/
// 题意：给定一个由 (, ), {, }, [ 和 ] 组成的字符串，判断括号是否合法

// 输入：()[]{}
// 输出：true

// 输入：(]
// 输出：false

// 思路：模拟栈即可，遍历字符串
//      如果是前括号则入栈
//      如果是后括号，栈顶是对应的前括号，则出栈；否则直接返回 false
//      最后栈为空，则括号合法

func isValid(s string) bool {
    validCharMap := map[int32] int32{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    stack := make([]int32, len(s))
    top := -1
    for _, char := range s {
        if char == '(' || char == '[' || char == '{' { // 前括号直接入栈
            top++
            stack[top] = char
        } else {
            if top < 0 || stack[top] != validCharMap[char] { // 若栈顶括号不是匹配的前括号，则直接返回 false
                return false
            }
            // 栈顶括号是匹配的前括号，出栈
            top--
        }
    }

    return top == -1 // 栈空，则括号合法
}