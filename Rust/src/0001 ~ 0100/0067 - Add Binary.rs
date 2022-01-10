// 链接：https://leetcode.com/problems/add-binary/
// 题意：给定两个 01 串 a 和 b ，返回它们的和？

// 输入： a = "11", b = "1"
// 输出： "100"

// 输入： a = "1010", b = "1011"
// 输出： "10101"


// 思路： 模拟
//
//      按照通常的加法器模拟即可，从个位开始按位加，注意进位，特别是最高位进位
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

use std::iter;

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        // 进位，初始为 0
        let mut carry = 0;
        // 收集 a 和 b 按位加的结果
        let mut result = a
            // 返回底层的字节切片
            .as_bytes()
            // 转换成迭代器
            .iter()
            // 提前反向迭代（后面会加上无限的 '0' ，所以不能后面同时反向）
            .rev()
            // 在 a 后面串上无限的 '0' ，以免 a 的长度更短时提前终止
            .chain(iter::repeat(&b'0'))
            // 同时反向遍历 a 和 b （对 b 也做前面对 a 做过的处理）
            .zip(b.as_bytes().iter().rev().chain(iter::repeat(&b'0')))
            // 遍历结束条件是 a 和 b 原本的字节都遍历完成，
            // a 和 b 末尾都是无限的 '0' ，所以只取必要的字节
            .take(a.len().max(b.len()))
            // 从个位开始进行计算每一位的结果
            .map(|(ach, bch)| {
                // 计算当前位的结果，共有 4 中情况： 0, 1, 2, 3
                let sum = (*ach - b'0') + (*bch - b'0') + carry;
                // 结果是 2 和 3 时有进位，所以直接取次低位
                carry = sum >> 1;
                
                // 最后根据结果返回当前位的计算结果
                match sum & 1 {
                    // 0 就直接返回 0
                    0 => '0',
                    // 剩余只有 1 ，直接返回 1
                    _ => '1',
                }
            })
            // 收集结果成 Vec<char>
            .collect::<Vec<_>>();
        
        // 如果最后还有进位，则还要加上 '1'
        if carry == 1{
            result.push('1')
        }
        
        // 转成迭代器后反向收集成 String
        result.iter().rev().collect()
    }
}