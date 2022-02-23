// 链接：https://leetcode.com/problems/excel-sheet-column-number/
// 题意：给定一个 Excel 的标题，返回其对应的列号？
//    	A -> 1
//    	B -> 2
//    	C -> 3
//    	...
//    	Z -> 26
//    	AA -> 27
//    	AB -> 28
//    	...


// 数据限制：
//  1 <= columnTitle.length <= 7
//  columnTitle 仅由英文大写字母组成
//  columnTitle 的取值范围是 ["A", "FXSHRXW"]


// 输入： "A"
// 输出： 1

// 输入： "AB"
// 输出： 28

// 输入： "ZY"
// 输出： 701


// 思路： 模拟
//
//      本题其实就是进制转换，但有一点需要特别注意：
//          列号是 26 进制，但每一位都是从 1 开始计数的，
//          也就是取值范围是 [1, 26] ，
//          而 ch - 'A' 的取值范围是 [0, 25] ，
//          所以还需要 +1 转换到对应范围
//
//
//		时间复杂度： O(n)
//          1. 只需要遍历全部 O(n) 个字母一次
//		空间复杂度： O(1)
//          1. 只维护常数个变量，所以空间复杂度为 O(1)


impl Solution {
    pub fn title_to_number(column_title: String) -> i32 {
        // 维护 ans ，表示最终列号
        let mut ans = 0;
        // 遍历 column_title ，计算列号
        for &ch in column_title.as_bytes() {
            // 列号是 26 进制，但每一位都是从 1 开始计数的，
            // 也就是取值范围是 [1, 26] ，
            // 而 ch - 'A' 的取值范围是 [0, 25] ，
            // 所以还需要 +1 转换到对应范围
            ans = ans * 26 + (ch - b'A' + 1) as i32;
        }

        ans
    }
}