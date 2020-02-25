// 链接：https://leetcode.com/problems/generate-parentheses/
// 题意：求出 n 对合法括号的所有可能情况

// 输入：3
// 输出：[
//          "((()))",
//          "(()())",
//          "(())()",
//          "()(())",
//          "()()()"
//       ]

// 思路：暴力递归即可
//      参数是还可以放入的前括号和后括号
//      当还可以放入前括号时，放入前括号递归，获得相应的结果
//      当还可以放入后括号时，放入后括号递归，获得相应的结果

func generateParenthesis(n int) []string {
    if n == 0 {
        return []string{}
    } 

    return doGenerateParenthesis("", n, 0)
}

func doGenerateParenthesis(cur string, front, back int) []string {
    if front == 0 && back == 0 { // 如果前后括号都已全部放完，则本次遍历的字符已经全部完成，直接返回
        return []string {cur}
    }

    var result []string // 本次递归内产生的结果

    if front > 0 { // 当还可以放入前括号时，放入前括号递归，获得相应的结果
        result = append(result, doGenerateParenthesis(cur + "(", front - 1, back + 1)...)
    }

    if back > 0 { // 当还可以放入后括号时，放入后括号递归，获得相应的结果
        result = append(result, doGenerateParenthesis(cur + ")", front, back - 1)...)
    }

    return result
}