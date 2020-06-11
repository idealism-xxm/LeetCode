// 链接：https://leetcode.com/problems/shortest-palindrome/
// 题意：给定一个字符串，求在其前面添加字符后，能形成的最短回文串？

// 输入： "aacecaaa"
// 输出： "aaacecaaa"

// 输入： "abcd"
// 输出： "dcbabcd"

// 思路1： 枚举
//
//		题目虽然是让求能形成的最短回文串，
//      但仔细一想其实我们只需要求目前给定的串包含第一个字符的最长回文子串即可，
//      假设最终的回文串最短，那么一定是以当前串的某一（两）个字符为中心对称，
//      且第一个字符一定在回文串在中，剩余后边的字符翻转后进行添加在前面
//      所以我们可以从后枚举回文子串的结束位置，找到第一个可以形成回文子串的地方即可，
//      然后把后边的字符翻转后添加在前面就形成了最短回文串
//
//      时间复杂度： O(n ^ 2)
//      空间复杂度： O(1)

impl Solution {
    pub fn shortest_palindrome(s: String) -> String {
        let bytes = s.as_bytes();
        for r in (0..bytes.len()).rev() {
            // 如果是回文子串，则满足题意
            if Solution::is_palindrome(&bytes, r) {
                // 获取不再回文子串的部分
                let mut result = bytes[r+1..].to_vec();
                // 翻转
                result.reverse();
                // 拼接
                result.append(&mut bytes.to_vec());
                // 转成字符串返回
                return String::from_utf8(result).unwrap();
            }
        }
        // 永远不可能走到这
        return s;
    }

    fn is_palindrome(bytes: &[u8], r: usize) -> bool {
        let mut l: usize = 0;
        let mut r = r;
        while l < r {
            // 如果存在一个字符不一样，则不是回文串
            if bytes[l] != bytes[r] {
                return false;
            }
            l += 1;
            r -= 1;
        }
        // 完全对称，返回 true
        true
    }
}

// 思路2： KMP
//
//		KMP 求 next 数组其实就是求最长公共前后缀，
//      我们将 s + "#" + s.reverse() 拼接起来，求一下 next 数组，
//      就能找到此时的最长公共前后缀，而这个最长公共前后缀就是最长的回文子串，
//      因此我们只要将剩余部分的串翻转后拼在前面就可以得到一个最短回文串
//      放入不存在的字符 '#' 是为了防止前后缀跨越了两个串
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

impl Solution {
    pub fn shortest_palindrome(s: String) -> String {
        // bytes 用于组装模式串
        let mut bytes = s.as_bytes().to_vec();
        // 翻转 bytes ，用于拼接为模式串和后续拼接为答案串
        let mut reversed = bytes.clone();
        reversed.reverse();
        // 组装模式串
        bytes.push(b'#');
        bytes.extend(reversed.iter());

        // 获取 next 数组最后一个 next 值（第 bytes.len() + 1 个 next 值）
        // last_next 即为最长公共前后缀的长度
        let last_next = Solution::get_last_next(&bytes);
        // 所以要从 reversed 中取剩余长度的字符串拼在前面
        let mut result = reversed[..reversed.len() - last_next].to_vec();
        result.extend(s.bytes());

        // 转成字符串返回
        String::from_utf8(result).unwrap()
    }

    fn get_last_next(pattern: &Vec<u8>) -> usize {
        // 多计算 1 个，即需要包含最后一个字符的最长公共前后缀
        let mut next = vec![-1; pattern.len() + 1];
        let mut j = 0;
        let mut k = -1;
        while j < pattern.len() {
            // 如果 k 已经移动到最开始，或者当前字符相同，则可以同时移动
            if k == -1 || pattern[j as usize] == pattern[k as usize] {
                // 同时移动到下一个字符
                j += 1;
                k += 1;
                // 不进行优化，直接赋值即可，这样保证所有字符对应的最长公共前后缀都非负
                next[j as usize] = k;
            } else {
                // 当前字符不同， k 需要往前找，直到两个字符相同，或者 k 已到最开始
                k = next[k as usize];
            }
        }

        // 最长公共前后缀
        next[pattern.len()] as usize
    }
}
