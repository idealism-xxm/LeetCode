// 链接：https://leetcode.com/problems/longest-uploaded-prefix/
// 题意：给定 1 ~ n 共 n 个数，题目按指定顺序给出这 n 个数。
//      实现一个数据结构，需要支持以下操作：
//          1. LUPrefix(int n): 初始化该数据结构
//          2. void upload(int video): 接收数字 video
//          3. int longest(): 在当前已接收的数字中，
//              返回最长前缀 [1, ..., i] 的长度


// 数据限制：
//  1 <= n <= 10 ^ 5
//  1 <= video <= n
//  所有的 video 都各不相同
//  总共最多会调用 2 * 10 ^ 5 次 upload 和 longest
//  至少调用 longest 一次


// 输入： ["LUPrefix", "upload", "longest", "upload", "longest", "upload", "longest"]
//       [[4], [3], [], [1], [], [2], []]
// 输出： [null, null, 0, null, 1, null, 3]
// 解释： LUPrefix server = new LUPrefix(4);   // 初始化 4 个数的数据结构
//       server.upload(3);                    // 接收数字 3
//       server.longest();                    // 因为数字 1 还未接收，不存在前缀，返回 0
//       server.upload(1);                    // 接收数字 1
//       server.longest();                    // 前缀 [1] 是最长的，返回 1
//       server.upload(2);                    // 接收数字 2
//       server.longest();                    // 前缀 [1,2,3] 是最长的，返回 1


// 思路： Set/Map
//
//      我们可以用集合 nums 维护已接收的数字，方便后续处理。
//
//      并用 prefix 维护已接收的数字中最长前缀的长度，初始化为 0 。
//
//      那么当调用 upload 时，先将数字放入 nums 中。
//
//      再不断更新 prefix 为 prefix + 1 ，直至 prefix + 1 不在 nums 中。
//      
//
//      时间复杂度：upload - O(1) | longest - O(1)
//          1. upload: 平均时间复杂度为 O(1) 。因为最多会调用 O(n) 次 upload ，
//              最多会更新 O(n) 次 prefix 。
//          2. longest: 直接返回维护好的 prefix 即可
//      空间复杂度：O(n)
//          1. 需要 nums 中全部 O(n) 个已接收的数字


use std::collections::HashSet;


struct LUPrefix {
    // prefix 维护已接收的数字中最长前缀的长度
    prefix: i32,
    // nums 维护已接收的数字
    nums: HashSet<i32>,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LUPrefix {

    fn new(n: i32) -> Self {
        LUPrefix {
            // 初始没有接收任何数字，最长前缀的长度为 0
            prefix: 0,
            nums: HashSet::with_capacity(n as usize),
        }
    }
    
    fn upload(&mut self, video: i32) {
        // 接收数字 video
        self.nums.insert(video);
        // 如果 prefix + 1 已接收，则最长前缀就是 prefix + 1
        while self.nums.contains(&(self.prefix + 1)) {
            self.prefix += 1;
        }
    }
    
    fn longest(&self) -> i32 {
        self.prefix
    }
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * let obj = LUPrefix::new(n);
 * obj.upload(video);
 * let ret_2: i32 = obj.longest();
 */
