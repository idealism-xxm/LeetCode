// 链接：https://leetcode.com/problems/palindrome-partitioning/
// 题意：给定一个字符串 s ，将 s 分割成一些回文子串，返回所有可能的分割方案。

// 数据限制：
//  1 <= s.length <= 16
//  s 仅由小写英文字母组成

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]

// 输入：s = "a"
// 输出：[["a"]]

// 思路：DP + 回溯
//
//      我们可以直接使用 dfs(i, cur) 递归枚举所有可能的方案，
//      其中 i 表示当前应该对 s[i..] 继续枚举下一个回文子串，
//          cur 表示 s[..i] 中一种可行方案
//
//      1. i == s.len(): cur 就是一种可行的方案，
//          可以直接放到 ans 中，然后退出递归
//      2. i != s.len(): s 中还有剩余的子串需要分割，
//          此时我们枚举下一个子串 s[i..=j] 。
//
//          由于需要判断 s[i..=j] 是否是回文串，如果直接判断的话，
//          那么存在大量重复的判断，会增加时间复杂度。
//
//          所以我们可以使用 DP 预处理的 is_palindrome 进行判断，
//          这样预处理时的时间复杂度为 O(n ^ 2) ，
//          判断的时间复杂度就是 O(1)
//
//      预处理 is_palindrome 可以使用 DP 完成，
//      我们先枚举子串的结束下标 j (0..s.len())，
//      再枚举子串的起始下标 i (0..=j) 。
//
//      那么如果当前子串 s[i..=j] 首尾字符相同，
//      且（ s[i+1..=j-1] 是空串 或 s[i+1..=j-1] 是回文串）时，
//      则 s[i..=j] 是回文串
//
//
//      时间复杂度：O(n ^ 2 + 2 ^ n) 。 预处理 is_palindrome 是 O(n ^ 2) ， dfs 是 O(2 ^ n)
//      空间复杂度：O(n ^ 2 + 2 ^ n) 。 is_palindrome 是 O(n ^ 2) ， ans 是 O(2 ^ n)


impl Solution {
    pub fn partition(s: String) -> Vec<Vec<String>> {
        // 先转成字节切片
        let s = s.as_bytes();
        // is_palindrome[i][j] 表明子串 s[i..=j] 是否是回文串
        let mut is_palindrome = vec![vec![false; s.len()]; s.len()];
        // 枚举子串结束下标
        for j in 0..s.len() {
            // 枚举子串起始下标
            for i in 0..=j {
                // 如果当前子串 s[i..=j] 首尾字符相同，
                // 且（ s[i+1..=j-1] 是空串 或 s[i+1..=j-1] 是回文串），
                // 则 s[i..=j] 是回文串
                if s[i] == s[j] && (j - i <= 2 || is_palindrome[i + 1][j - 1]) {
                    is_palindrome[i][j] = true;
                }
            }
        }

        // ans 用于收集所有可能的方案
        let mut ans = vec![];
        Self::dfs(s, &is_palindrome, 0, &mut vec![], &mut ans);
        ans
    }

    // dfs 递归搜索答案。
    //  s               表示待拆分的字符串
    //  is_palindrome   用于判断子串 s[i..=j] 是否是回文串
    //  i               表示当前应该对 s[i..] 继续搜索
    //  cur             表示当前 s[..i] 的一种搜索结果
    //  ans             表示已经收集的方案列表
    fn dfs(s: &[u8], is_palindrome: &Vec<Vec<bool>>, i: usize, cur: &mut Vec<String>, ans: &mut Vec<Vec<String>>) {
        // 如果 i == s.len() ，则说明一种方案已经搜索完成，
        // 把当前方案 cur 放入 ans 后返回即可
        if i == s.len() {
            ans.push(cur.clone());
            return
        }
        // 遍历下一个回文子串 s[i..=j]
        (i..s.len())
            // 仅处理回文子串
            .filter(|&j| is_palindrome[i][j])
            .for_each(|j| {
                // 将当前方案的回文子串收集到 cur 中
                cur.push(String::from_utf8(s[i..=j].to_vec()).unwrap());
                // 递归处理
                Self::dfs(s, is_palindrome, j + 1, cur, ans);
                // 移除当前方案的回文子串
                cur.pop();
            });
    }
}
