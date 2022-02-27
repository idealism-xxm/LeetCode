// 链接：https://leetcode.com/problems/compare-version-numbers/
// 题意：给定两个版本号（只含有数字和 '.'），比较这两个版本号的大小 ？


// 数据限制：
//  1 <= version1.length, version2.length <= 500
//  version1 和 version2 只含有数字和 '.'
//  version1 和 version2 都是合法的版本号
//  version1 和 version2 中的修订号都可以存储在 32 位整型中


// 输入： version1 = "1.01", version2 = "1.001"
// 输出： 0
// 解释： 忽略前导零， "01" 和 "001" 都代表数字 1

// 输入： version1 = "1.0", version2 = "1.0.0"
// 输出： 0
// 解释： 第一个版本号的第三个修订号没有，则默认为 0

// 输入： version1 = "0.1", version2 = "1.1"
// 输出： -1
// 解释： 第一个版本号的第一个修订号是 0 ，
//       而第二个版本号的第一个修订号是 1 ，
//       所以 version1 < version2


// 思路： 模拟
//
//		先按照 . 划分出修订号列表 revisions1 和 revisions2 。
//
//		然后对两者中较短的补 "0" ，让它们长度相等。
//
//	    再按顺序开始比较修订号 revision1 和 revision2 ：
//          1. revision1 < revision2: 则说明 version1 < version2 ，
//              直接返回 -1
//          2. revision1 > revision2: 则说明 version1 > version2 ，
//              直接返回 1
//          3. revision1 == revision2: 则说明现在还无法判断版本号大小，
//              继续处理下一个修订号
//
//      如果最终所有的修订号都相等，则说明 version1 == version2 ，返回 0 。
//
//
//		当然也可以优化为空间复杂度是 O(1) 的解法：
//          不进行划分，而是同时扫描两个版本号，
//          并计算每一个修订号，遇到 '.' 或者末尾时停止，
//		    此时按照前面提到的方法，对比两者的修订号即可，
//          注意不够的默认补 0 。
//
//
//      时间复杂度：O(max(n, m))
//          1. 需要遍历 version1 的全部 O(n) 个字符，
//              生成长度为 O(n) 的修订号列表 revisions1
//          2. 需要遍历 version2 的全部 O(m) 个字符，
//              生成长度为 O(m) 的修订号列表 revisions2
//          3. 需要同时遍历 revisions1 和 revisions2 全部 O(max(n, m)) 个修订号
//      空间复杂度：O(max(n, m))
//          1. 需要生成长度为 O(max(n, m)) 的修订号列表 revisions1 和 revisions2

use std::cmp::Ordering;

impl Solution {
    pub fn compare_version(version1: String, version2: String) -> i32 {
        // 获取 version1 和 version2 的所有整型修订号列表
        let mut revisions1 = Self::get_revisions(&version1);
        let mut revisions2 = Self::get_revisions(&version2);

        // 如果 revisions1 的长度更小，则不断在后面补 0 ，
        // 直至两者长度相等
        while revisions1.len() < revisions2.len() {
            revisions1.push(0);
        }
        // 如果 revisions2 的长度更小，则不断在后面补 0 ，
        // 直至两者长度相等
        while revisions2.len() < revisions1.len() {
            revisions2.push(0);
        }
        
        // 此时两者长度相等，所以 cmp 的结果就是 revisions 的对比结果
        match revisions1.cmp(&revisions2) {
            // 如果前者小，则返回 -1
            Ordering::Less => -1,
            // 如果两者相等，则返回 0
            Ordering::Equal => 0,
            // 如果后者小，则返回 1
            Ordering::Greater => 1,
        }
    }

    #[inline]
    pub fn get_revisions(version: &String) -> Vec<i32> {
        version
            // 先按 '.' 分割成不同的 revision
            .split('.')
            // 再将每个 revision 转成整型
            .map(|revision| revision.parse::<i32>().unwrap())
            // 最后收集成一个 Vec<i32> 并返回
            .collect()
    }
}
