// 链接：https://leetcode.com/problems/break-a-palindrome/
// 题意：给定一个回文串 palindrome ，替换其中一个字母，使其不再回文。
//      求能得到的字典序最小的非回文串？
//      如果不存在，则返回空串。


// 数据限制：
//  1 <= palindrome.length <= 1000
//  palindrome 仅由英文小写字母组成


// 输入： palindrome = "abccba"
// 输出： "aaccba"
// 解释： 替换其中一个字母，能得到很多非回文串，
//       例如："(z)bccba", "a(a)ccba", and "ab(a)cba" 。
//       其中 "aaccba" 是能得到的字典序最小的非回文串。

// 输入： palindrome = "a"
// 输出： ""
// 解释： 无论替换成什么字母，一个字母必定是回文串，所以返回空串。


// 思路： 贪心
//
//      长度为 1 的字符串必定是回文串，直接返回空串即可。
//
//      长度大于 1 的回文串必定能通过替换一个字母转换为非回文串。
//
//      此时我们需要让这个非回文串最小，首先想到的就是减小其字典序，
//      那么必定要贪心地将第一个不为 'a' 的字母替换为 'a' 。
//
//      【注意】该字母不能是奇数长度回文串正中间的字母，因为其不影响回文性。
//
//      如果回文串中的字母全是 'a' ，那么只能增大其字典序，将最后一个字母替换为 'b' 。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 palindrome 中 O(n) 个字母
//      空间复杂度：O(n)
//          1. 需要维护结果中全部 O(n) 个字母
//             （在 Rust 中直接复用了原有的内存，可以只需要 O(1) 的额外空间）


impl Solution {
    pub fn break_palindrome(palindrome: String) -> String {
        // 长度为 1 的字符串必定是回文串，直接返回空串
        if palindrome.len() == 1 {
            return "".to_string();
        }

        let mut bytes = palindrome.into_bytes();
        // 找到 bytes 中第一个不为 'a' 的字母的下标
        let pos = bytes
            .iter()
            // 由于本身是回文串，所以只需要在前一半中寻找即可。
            // 【注意】该字母不能是奇数长度回文串正中间的字母，因为其不影响回文性
            .take(bytes.len() >> 1)
            // 找到第一个不是 'a' 的字母的下标
            .position(|&ch| ch != b'a');

        if let Some(pos) = pos {
            // 如果存在这样的下标，则直接替换为 'a' 即可
            bytes[pos] = b'a';
        } else {
            // 此时字符串中全是 'a' ，把最后一个字母替换为 'b' 时，得到的字符串的字典序最小
            let pos = bytes.len() - 1;
            bytes[pos] = b'b';
        }

        // 转换成字符串返回即可
        String::from_utf8(bytes).unwrap()
    }
}
